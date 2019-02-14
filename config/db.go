package config

import (
	"math/rand"
	"strconv"
	"strings"
	"team-academy/classroom"
	"team-academy/department"
	"team-academy/grade"
	"team-academy/professor"
	"team-academy/professor_subject"
	"team-academy/schedule"
	"team-academy/shift"
	"team-academy/student"
	"team-academy/student_shift"
	"team-academy/student_subject"
	"team-academy/subject"

	"github.com/jinzhu/gorm"
)

func PopulateDatabase(db *gorm.DB) (err error) {
	existsSubjectTable, err := subject.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	existsProfessorTable, err := professor.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	existsStudentTable, err := student.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	existsProfessorSubjectTable, err := professor_subject.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	existsStudentSubjectTable, err := student_subject.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	existsGradeTable, err := grade.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	existsDepartmentTable, err := department.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	existsClassroomTable, err := classroom.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	existsShiftTable, err := shift.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	existsStudentShiftTable, err := student_shift.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	existsScheduleTable, err := schedule.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	rand.Seed(1)

	if !existsSubjectTable {
		newSubject := subject.Subject{Name: "Análise Matemática I", Description: "Cálculo Diferencial e Integral em R"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{Name: "Álgebra Linear e Geometria Analítica", Description: "Matrizes, sistemas, determinantes, etc."}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{Name: "Desenho Assistido por Computador", Description: "AutoCAD para plantas elétricas"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{Name: "Programação de Microprocessadores", Description: "Programação em C"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{Name: "Sistemas Lógicos I", Description: "Lógica Booleana, Portas Lógicas, Biestáveis, etc."}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{Name: "Competências Transversais às Ciências e Tecnologia", Description: "Curriculum Vitae, entrevistas de emprego, bibliografias, etc."}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{Name: "Algoritmos e Estruturas de Dados", Description: "TADs"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{Name: "Análise Matemática II B", Description: "Cálculo Diferencial e Integral em Rn"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{Name: "Física I", Description: "Mecânica Clássica"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{Name: "Sistemas Lógicos II", Description: "Microprocessador de 16-bits"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{Name: "Teoria de Circuitos Elétricos", Description: "Lei de Ohm, nós, malhas, circuitos reativos, etc."}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{Name: "Análise Matemática III B", Description: "Séries e números complexos"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{Name: "Cálculo Numérico", Description: "Métodos numéricos para resolução de problemas de cálculo"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{Name: "Física III", Description: "Eletromagnetismo com o Meme Paiva"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{Name: "Introdução às Telecomunicações", Description: "Transformada de Fourier, Modulações Analógicas e Digitais, etc."}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{Name: "Microprocessadores", Description: "Memória RAM, Protocolos de Transferência de Dados, Assembly"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{Name: "Ciências, Tecnologia e Sociedade", Description: "Aulas de humanidades numa faculdade de ciências"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{Name: "Análise Matemática IV B", Description: "Equações diferenciais"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{Name: "Eletrónica I", Description: "Díodos, Transístores, Ampops, etc."}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{Name: "Probabilidade e Estatística C", Description: "Probabilidade e Estatística"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{Name: "Sistemas de Telecomunicações", Description: "Camadas do Modelo OSI, Protocolos, Java"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{Name: "Teoria de Sinais", Description: "D'onde"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{Name: "Eletrotecnia Teórica", Description: "TEDx - Introdução à Meditação com Mário Neves"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{Name: "Eletrónica II", Description: "Ampops, Realimentação, Estabilidade, Osciladores, Filtros"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{Name: "Física II", Description: "Cinética e Termodinâmica"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{Name: "Sistemas de Tempo Real", Description: "Tutorial de Solitário com Luís Matos"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{Name: "Teoria de Controlo", Description: "Diagramas de blocos, etc."}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{Name: "Programa de Introdução à Investigação Científica", Description: "Estágio"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{Name: "Controlo por Computador", Description: "TC Digital"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{Name: "Conversão Eletromecânica de Energia", Description: "ET 2"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{Name: "Instrumentação e Medidas Elétricas", Description: "Carrola FTW"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{Name: "Modelação de Dados em Engenharia", Description: "Databases are easy"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{Name: "Propagação e Radiação", Description: "Female Ventim"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}
	}

	if !existsProfessorTable {
		for i := 1; i <= 10; i++ {
			firstName := randomFirstName()
			lastName := randomLastName()
			email := generateProfessorEmail(i, firstName, lastName)
			newProfessor := professor.Professor{FirstName: firstName, LastName: lastName, CursoID: randomInt(1, 5), StartDate: int64(randomInt(220924800, 1419984000)), Email: email}
			err = professor.CreateProfessor(db, newProfessor)
			if err != nil {
				return
			}
		}
	}

	if !existsStudentTable {
		for i := 1; i <= 10; i++ {
			firstName := randomFirstName()
			lastName := randomLastName()
			email := generateStudentEmail(i, firstName, lastName)
			newStudent := student.Student{FirstName: firstName, LastName: lastName, CursoID: randomInt(1, 5), StartDate: int64(randomInt(1420070400, 1546300800)), Email: email}
			err = student.CreateStudent(db, newStudent)
			if err != nil {
				return
			}
		}
	}

	if !existsProfessorSubjectTable {
		for i := 1; i <= 10; i++ {
			for j := 1; j <= 33; j++ {
				add := randomInt(0, 9)
				if add >= 7 {
					err = professor_subject.AddProfessorToSubject(db, i, j)
					if err != nil {
						return
					}
				}
			}
		}
	}

	if !existsStudentSubjectTable {
		for i := 1; i <= 10; i++ {
			for j := 1; j <= 33; j++ {
				add := randomInt(0, 9)
				if add >= 7 {
					err = student_subject.AddStudentToSubject(db, i, j)
					if err != nil {
						return
					}
				}
			}
		}
	}

	if !existsGradeTable {
		for i := 1; i <= 10; i++ {
			for j := 1; j <= 33; j++ {
				newGrade := grade.Grade{StudentID: i, SubjectID: j, Rank: randomGrade(0.0, 20.0)}
				grade.GiveGrade(db, newGrade)
			}
		}
	}

	if !existsDepartmentTable {
		newDepartment := department.Department{Name: "Departamento de Engenharia Eletrotécnica"}
		err = department.CreateDepartment(db, newDepartment)
		if err != nil {
			return
		}

		newDepartment = department.Department{Name: "Departamento de Matemática"}
		err = department.CreateDepartment(db, newDepartment)
		if err != nil {
			return
		}
	}

	if !existsClassroomTable {
		for i := 1; i <= 2; i++ {
			newClassroom := classroom.Classroom{Name: strconv.Itoa(i), DepartmentID: i}
			err = classroom.CreateClassroom(db, newClassroom)
			if err != nil {
				return
			}
		}
	}

	if !existsShiftTable {
		for i := 1; i <= 11; i++ {
			newShift := shift.Shift{SubjectID: randomInt(1, 11), ClassroomID: randomInt(1, 24), Type: randomShiftType(), ShiftNum: randomInt(1, 5)}
			err = shift.CreateShift(db, newShift)
			if err != nil {
				return
			}
		}
	}

	if !existsStudentShiftTable {
		for i := 1; i <= 10; i++ {
			err = student_shift.AddStudentToShift(db, randomInt(1, 11), randomInt(1, 11))
			if err != nil {
				return
			}
		}
	}

	if !existsScheduleTable {
		for i := 1; i <= 14; i++ {
			newSchedule := schedule.Schedule{SubjectID: randomInt(1, 11), ShiftID: randomInt(1, 14), Weekday: randomInt(0, 6), StartTime: 28800, EndTime: 36600}
			err = schedule.CreateSchedule(db, newSchedule)
			if err != nil {
				return
			}
		}
	}

	return
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randomFirstName() string {
	nameList := []string{"José", "António", "João", "Manuel", "Carlos",
		"Paulo", "Fernando", "Luís", "Joaquim", "Jorge",
		"Maria", "Ana", "Isabel", "Rosa", "Paula",
		"Anabela", "Fernanda", "Teresa", "Cristina", "Helena",
		"João", "Diogo", "Pedro", "Tiago", "Gonçalo",
		"Rodrigo", "Miguel", "Francisco", "José", "André",
		"Ana", "Maria", "Beatriz", "Mariana", "Inês",
		"Joana", "Carolina", "Catarina", "Sara", "Daniela"}

	return nameList[randomInt(0, len(nameList)-1)]
}

func randomLastName() string {
	nameList := []string{"Silva", "Santos", "Pereira", "Ferreira", "Costa",
		"Oliveira", "Rodrigues", "Martins", "Fernandes", "Sousa",
		"Gonçalves", "Gomes", "Lopes", "Carvalho", "Ribeiro",
		"Pinto", "Marques", "Almeida", "Alves", "Teixeira",
		"Dias", "Monteiro", "Correia", "Moreira", "Mendes",
		"Vieira", "Cardoso", "Soares", "Nunes", "Rocha",
		"Coelho", "Duarte", "Cunha", "Tavares", "Ramos",
		"Cruz", "Neves", "Reis", "Freitas", "Araújo"}

	return nameList[randomInt(0, len(nameList)-1)]
}

func generateProfessorEmail(id int, firstName, lastName string) string {
	return strings.ToLower(string(firstName[0])) + "." + strings.ToLower(lastName) + "_" + strconv.Itoa(id) + "@fct.unl.pt"
}

func generateStudentEmail(id int, firstName, lastName string) string {
	return strings.ToLower(string(firstName[0])) + "." + strings.ToLower(lastName) + "_" + strconv.Itoa(id) + "@campus.fct.unl.pt"
}

func randomGrade(min, max float64) string {
	r := min + rand.Float64()*(max-min)
	return strconv.FormatFloat(r, 'f', 2, 64)
}

func randomShiftType() string {
	types := [3]string{"T", "TP", "P"}
	return types[randomInt(0, len(types))]
}
