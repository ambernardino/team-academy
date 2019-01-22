package grade

import (
	"fmt"
	"team-academy/student_subject"
	"github.com/jinzhu/gorm"
)

type Grade struct {
	SubjectID int
	StudentID int
	Rank      string
}

type StudentGrade struct {
    StudentID        int
    StudentFirstName string `gorm:"column:first_name"`
    StudentLastName  string `gorm:"column:last_name"`
    SubjectName      string `gorm:"column:name"`
    SubjectID        int
}

func CreateTableIfNotExists(db *gorm.DB) (err error) {
	if !db.HasTable(Grade{}) {
		return db.CreateTable(Grade{}).Error
	}
	return
}

func GiveGrade(db *gorm.DB, grade Grade) (err error) {

	_, err = db.Table("student").Select("student.id, student_subject.subject_id").Joins("JOIN student_subject ON student_subject.student_id = student.id").Where(&student_subject.StudentSubject{SubjectID : grade.SubjectID}).Rows()
	if err != nil {
		fmt.Println(err)
		return
	}
	return db.Save(&grade).Error
}

func GetGradeByStudentIDAndSubjectID(db *gorm.DB, studentID int, subjectID int) (grade Grade, err error) {
	err = db.First(&grade, &Grade{StudentID: studentID, SubjectID: subjectID}).Error
	return
}

func GetStudentsGrades(db *gorm.DB) (grades []StudentGrade, err error) {
    err = db.Table("student").Select("student_subject.student_id, student.first_name, student.last_name, subject.name, student_subject.subject_id").Joins("JOIN student_subject ON student.id = student_subject.student_id").Joins("JOIN subject on subject.id = student_subject.subject_id").Scan(&grades).Error
    for _, v := range grades {
        fmt.Println(v)
		//fmt.Printf("Nr aluno: %v, Nome aluno: %s %s, Nome cadeira: %s, Nr cadeira: %v\n"
		//, v.StudentID, v.StudentFirstName, v.StudentLastName, v.SubjectName, v.SubjectID)
    }
    return
}


