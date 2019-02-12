package component

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"reflect"
	"runtime"
	"strings"
)

// The maximum number of stackframes on any error.
var MaxStackDepth = 50

// Error is an error with an attached stacktrace. It can be used
// wherever the builtin error interface is expected.
type Error struct {
	Err    error
	stack  []uintptr
	frames []StackFrame
	prefix string
}

// A StackFrame contains all necessary information about to generate a line
// in a callstack.
type StackFrame struct {
	// The path to the file containing this ProgramCounter
	File string
	// The LineNumber in that file
	LineNumber int
	// The Name of the function that contains this ProgramCounter
	Name string
	// The Package that contains this function
	Package string
	// The underlying ProgramCounter
	ProgramCounter uintptr
}

// New makes an Error from the given value. If that value is already an
// error then it will be used directly, if not, it will be passed to
// fmt.Errorf("%v"). The stacktrace will point to the line of code that
// called New.
func New(e interface{}) *Error {
	var err error
	switch e := e.(type) {
	case error:
		err = e
	default:
		err = fmt.Errorf("%v", e)
	}
	stack := make([]uintptr, MaxStackDepth)
	length := runtime.Callers(2, stack[:])
	return &Error{
		Err:   err,
		stack: stack[:length],
	}
}

// Wrap makes an Error from the given value. If that value is already an
// error then it will be used directly, if not, it will be passed to
// fmt.Errorf("%v"). The skip parameter indicates how far up the stack
// to start the stacktrace. 0 is from the current call, 1 from its caller, etc.
func Wrap(e interface{}, skip int) *Error {
	var err error

	switch e := e.(type) {
	case *Error:
		return e
	case error:
		err = e
	default:
		err = fmt.Errorf("%v", e)
	}

	stack := make([]uintptr, MaxStackDepth)
	length := runtime.Callers(2+skip, stack[:])

	return &Error{
		Err:   err,
		stack: stack[:length],
	}
}

// WrapPrefix makes an Error from the given value. If that value is already an
// error then it will be used directly, if not, it will be passed to
// fmt.Errorf("%v"). The prefix parameter is used to add a prefix to the
// error message when calling Error(). The skip parameter indicates how far
// up the stack to start the stacktrace. 0 is from the current call,
// 1 from its caller, etc.
func WrapPrefix(e interface{}, prefix string, skip int) *Error {
	err := Wrap(e, 1+skip)
	if err.prefix != "" {
		prefix = fmt.Sprintf("%s: %s", prefix, err.prefix)
	}

	return &Error{
		Err:    err.Err,
		stack:  err.stack,
		prefix: prefix,
	}
}

// IsEqual detects whether the error is equal to a given error. Errors
// are considered equal by this function if they are the same object,
// or if they both contain the same error inside an errors.Error.
func IsEqual(e error, original error) bool {
	if e == original {
		return true
	}

	if e, ok := e.(*Error); ok {
		return IsEqual(e.Err, original)
	}

	if original, ok := original.(*Error); ok {
		return IsEqual(e, original.Err)
	}

	return false
}

// Errorf creates a new error with the given message. You can use it
// as a drop-in replacement for fmt.Errorf() to provide descriptive
// errors in return values.
func Errorf(format string, a ...interface{}) *Error {
	return Wrap(fmt.Errorf(format, a...), 1)
}

// Error returns the underlying error's message.
func (err *Error) Error() string {
	msg := err.Err.Error()
	if err.prefix != "" {
		msg = fmt.Sprintf("%s: %s", err.prefix, msg)
	}

	return msg
}

// Stack returns the callstack formatted the same way that go does
// in runtime/debug.Stack()
func (err *Error) Stack() []byte {
	buf := bytes.Buffer{}
	for _, frame := range err.StackFrames() {
		buf.WriteString(frame.String())
	}

	return buf.Bytes()
}

// Callers satisfies the bugsnag ErrorWithCallerS() interface
// so that the stack can be read out.
func (err *Error) Callers() []uintptr {
	return err.stack
}

// ErrorStack returns a string that contains both the
// error message and the callstack.
func (err *Error) ErrorStack() string {
	return err.TypeName() + " " + err.Error() + "\n" + string(err.Stack())
}

// StackFrames returns an array of frames containing information about the
// stack.
func (err *Error) StackFrames() []StackFrame {
	if err.frames == nil {
		for _, pc := range err.stack {
			st := NewStackFrame(pc)
			if strings.HasPrefix(st.Package, "plicca-go/cmd/") {
				st.ProgramCounter = 0
				err.frames = append(err.frames, st) //err.frames[i] = st
			}
		}
	}
	return err.frames
}

// TypeName returns the type this error. e.g. *errors.stringError.
func (err *Error) TypeName() string {
	return reflect.TypeOf(err.Err).String()
}

// NewStackFrame popoulates a stack frame object from the program counter.
func NewStackFrame(pc uintptr) (frame StackFrame) {
	frame = StackFrame{ProgramCounter: pc}
	if frame.Func() == nil {
		return
	}

	frame.Package, frame.Name = packageAndName(frame.Func())

	// pc -1 because the program counters we use are usually return addresses,
	// and we want to show the line that corresponds to the function call
	frame.File, frame.LineNumber = frame.Func().FileLine(pc - 1)
	return
}

// Func returns the function that contained this frame.
func (frame *StackFrame) Func() *runtime.Func {
	if frame.ProgramCounter == 0 {
		return nil
	}

	return runtime.FuncForPC(frame.ProgramCounter)
}

// String returns the stackframe formatted in the same way as go does
// in runtime/debug.Stack()
func (frame *StackFrame) String() string {
	str := fmt.Sprintf("%s:%d\n", frame.File, frame.LineNumber) //, frame.ProgramCounter)
	source, err := frame.SourceLine()
	if err != nil {
		return str
	}

	return str + fmt.Sprintf("\t%s: %s\n", frame.Name, source)
}

// SourceLine gets the line of code (from File and Line) of the original source if possible.
func (frame *StackFrame) SourceLine() (string, error) {
	data, err := ioutil.ReadFile(frame.File)
	if err != nil {
		return "", New(err)
	}

	lines := bytes.Split(data, []byte{'\n'})
	if frame.LineNumber <= 0 || frame.LineNumber >= len(lines) {
		return "???", nil
	}

	// -1 because line-numbers are 1 based, but our array is 0 based
	return string(bytes.Trim(lines[frame.LineNumber-1], " \t")), nil

}

func packageAndName(fn *runtime.Func) (string, string) {
	name := fn.Name()
	pkg := ""

	// The name includes the path name to the package, which is unnecessary
	// since the file name is already included.  Plus, it has center dots.
	// That is, we see
	//  runtime/debug.*T·ptrmethod
	// and want
	//  *T.ptrmethod
	// Since the package path might contains dots (e.g. code.google.com/...),
	// we first remove the path prefix if there is one.
	if lastslash := strings.LastIndex(name, "/"); lastslash >= 0 {
		pkg += name[:lastslash] + "/"
		name = name[lastslash+1:]
	}

	if period := strings.Index(name, "."); period >= 0 {
		pkg += name[:period]
		name = name[period+1:]
	}

	name = strings.Replace(name, "·", ".", -1)
	return pkg, name
}
