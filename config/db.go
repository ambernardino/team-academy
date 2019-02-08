package config

import (
	"team-academy/grade"
	"team-academy/professor"
	"team-academy/professor_subject"
	"team-academy/student"
	"team-academy/student_subject"
	"team-academy/subject"
	"time"

	"github.com/jinzhu/gorm"
)

func PopulateDatabase(db *gorm.DB) (err error) {
	existsProfessorTable, err := professor.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	existsStudentTable, err := student.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	existsSubjectTable, err := subject.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	existsStudentSubjectTable, err := student_subject.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	existsProfessorSubjectTable, err := professor_subject.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	existsGradeTable, err := grade.CreateTableIfNotExists(db)
	if err != nil {
		return
	}

	if !existsSubjectTable {
		newSubject := subject.Subject{ID: 1, Name: "Análise Matemática I", Description: "Cálculo Diferencial e Integral em R"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{ID: 2, Name: "Álgebra Linear e Geometria Analítica", Description: "Matrizes, sistemas, determinantes, etc."}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{ID: 3, Name: "Desenho Assistido por Computador", Description: "AutoCAD para plantas elétricas"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{ID: 4, Name: "Programação de Microprocessadores", Description: "Programação em C"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{ID: 5, Name: "Sistemas Lógicos I", Description: "Lógica Booleana, Portas Lógicas, Biestáveis, etc."}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{ID: 6, Name: "Competências Transversais às Ciências e Tecnologia", Description: "Curriculum Vitae, entrevistas de emprego, bibliografias, etc."}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{ID: 7, Name: "Algoritmos e Estruturas de Dados", Description: "TADs"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{ID: 8, Name: "Análise Matemática II B", Description: "Cálculo Diferencial e Integral em Rn"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{ID: 9, Name: "Física I", Description: "Mecânica Clássica"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{ID: 10, Name: "Sistemas Lógicos II", Description: "Microprocessador de 16-bits"}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}

		newSubject = subject.Subject{ID: 11, Name: "Teoria de Circuitos Elétricos", Description: "Lei de Ohm, nós, malhas, circuitos reativos, etc."}
		err = subject.CreateSubject(db, newSubject)
		if err != nil {
			return
		}
	}

	if !existsProfessorTable {
		newProfessor := professor.Professor{ID: 1, FirstName: "Ana", LastName: "Sá", CursoID: 1, StartDate: time.Now().UTC().Unix()}
		err = professor.CreateProfessor(db, newProfessor)
		if err != nil {
			return
		}

		newProfessor = professor.Professor{ID: 2, FirstName: "António", LastName: "Paiva", CursoID: 2, StartDate: time.Now().UTC().Unix()}
		err = professor.CreateProfessor(db, newProfessor)
		if err != nil {
			return
		}

		newProfessor = professor.Professor{ID: 3, FirstName: "Helena", LastName: "Fino", CursoID: 3, StartDate: time.Now().UTC().Unix()}
		err = professor.CreateProfessor(db, newProfessor)
		if err != nil {
			return
		}

		newProfessor = professor.Professor{ID: 4, FirstName: "Fernanda", LastName: "Barbosa", CursoID: 4, StartDate: time.Now().UTC().Unix()}
		err = professor.CreateProfessor(db, newProfessor)
		if err != nil {
			return
		}

		newProfessor = professor.Professor{ID: 5, FirstName: "Ruy", LastName: "Costa", CursoID: 5, StartDate: time.Now().UTC().Unix()}
		err = professor.CreateProfessor(db, newProfessor)
		if err != nil {
			return
		}
	}

	if !existsStudentTable {
		newStudent := student.Student{ID: 1, FirstName: "Francisco", LastName: "Peres", CursoID: 3, StartDate: time.Now().UTC().Unix()}
		err = student.CreateStudent(db, newStudent)
		if err != nil {
			return
		}

		newStudent = student.Student{ID: 2, FirstName: "Pedro", LastName: "Oliveira", CursoID: 3, StartDate: time.Now().UTC().Unix()}
		err = student.CreateStudent(db, newStudent)
		if err != nil {
			return
		}

		newStudent = student.Student{ID: 3, FirstName: "Shazia", LastName: "Sulemane", CursoID: 3, StartDate: time.Now().UTC().Unix()}
		err = student.CreateStudent(db, newStudent)
		if err != nil {
			return
		}

		newStudent = student.Student{ID: 4, FirstName: "Tiago", LastName: "Marques", CursoID: 3, StartDate: time.Now().UTC().Unix()}
		err = student.CreateStudent(db, newStudent)
		if err != nil {
			return
		}

		newStudent = student.Student{ID: 5, FirstName: "António", LastName: "Bernardino", CursoID: 4, StartDate: time.Now().UTC().Unix()}
		err = student.CreateStudent(db, newStudent)
		if err != nil {
			return
		}

		newStudent = student.Student{ID: 6, FirstName: "Pedro", LastName: "Grilo", CursoID: 4, StartDate: time.Now().UTC().Unix()}
		err = student.CreateStudent(db, newStudent)
		if err != nil {
			return
		}

		newStudent = student.Student{ID: 7, FirstName: "Daniela", LastName: "Gonçalves", CursoID: 5, StartDate: time.Now().UTC().Unix()}
		err = student.CreateStudent(db, newStudent)
		if err != nil {
			return
		}
	}

	if !existsStudentSubjectTable {
		for i := 1; i <= 4; i++ {
			for j := 1; j <= 11; j++ {
				err = student_subject.AddStudentToSubject(db, i, j)
				if err != nil {
					return
				}
			}
		}

		err = student_subject.AddStudentToSubject(db, 5, 4)
		if err != nil {
			return
		}

		err = student_subject.AddStudentToSubject(db, 5, 7)
		if err != nil {
			return
		}

		err = student_subject.AddStudentToSubject(db, 6, 5)
		if err != nil {
			return
		}

		err = student_subject.AddStudentToSubject(db, 6, 10)
		if err != nil {
			return
		}

		err = student_subject.AddStudentToSubject(db, 7, 6)
		if err != nil {
			return
		}
	}

	if !existsProfessorSubjectTable {
		err = professor_subject.AddProfessorToSubject(db, 1, 1)
		if err != nil {
			return
		}

		err = professor_subject.AddProfessorToSubject(db, 1, 2)
		if err != nil {
			return
		}

		err = professor_subject.AddProfessorToSubject(db, 1, 8)
		if err != nil {
			return
		}

		err = professor_subject.AddProfessorToSubject(db, 2, 9)
		if err != nil {
			return
		}

		err = professor_subject.AddProfessorToSubject(db, 3, 3)
		if err != nil {
			return
		}

		err = professor_subject.AddProfessorToSubject(db, 3, 5)
		if err != nil {
			return
		}

		err = professor_subject.AddProfessorToSubject(db, 3, 10)
		if err != nil {
			return
		}

		err = professor_subject.AddProfessorToSubject(db, 3, 11)
		if err != nil {
			return
		}

		err = professor_subject.AddProfessorToSubject(db, 4, 4)
		if err != nil {
			return
		}

		err = professor_subject.AddProfessorToSubject(db, 4, 7)
		if err != nil {
			return
		}

		err = professor_subject.AddProfessorToSubject(db, 5, 6)
		if err != nil {
			return
		}
	}

	if !existsGradeTable {
		for i := 1; i <= 11; i++ {
			for j := 1; j <= 7; j++ {
				newGrade := grade.Grade{SubjectID: i, StudentID: j, Rank: "9,5"}
				err = grade.GiveGrade(db, newGrade)
				if err != nil {
					return
				}
			}
		}
	}

	return
}
