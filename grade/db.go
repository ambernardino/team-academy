package grade

import (
	"fmt"
	"team-academy/component"
	"team-academy/student"
	"team-academy/student_subject"

	"github.com/jinzhu/gorm"
)

type Grade struct {
	ID        int    `gorm:"AUTO_INCREMENT"`
	SubjectID int    `json:"subject_id,omitempty"`
	StudentID int    `json:"student_id,omitempty"`
	Rank      string `json:"rank,omitempty"`
}

type StudentGrade struct {
	StudentID        int
	StudentFirstName string `gorm:"column:first_name"`
	StudentLastName  string `gorm:"column:last_name"`
	SubjectName      string `gorm:"column:name"`
	SubjectID        int
	Rank             string
}

func CreateTableIfNotExists(db *gorm.DB) (exists bool, err error) {
	if !db.HasTable(Grade{}) {
		return false, db.CreateTable(Grade{}).Error
	}
	return true, nil
}

func GiveGrade(db *gorm.DB, grade Grade) (err error) {
	_, err = student_subject.GetStudentSubject(db, grade.StudentID, grade.SubjectID)
	if err != nil {
		err = component.ErrStudentNotInSubject
		return
	}

	rows, err := db.Table("grade").Select("grade.subject_id, grade.student_id").Joins("JOIN student ON grade.student_id = student.id").Joins("JOIN subject ON grade.subject_id = subject.id").Joins("JOIN student_subject ON grade.student_id = student_subject.student_id AND grade.subject_id = student_subject.subject_id").Where(&Grade{StudentID: grade.StudentID, SubjectID: grade.SubjectID}).Rows()
	if err != nil {
		return
	}

	defer rows.Close()
	if rows.Next() {
		err = component.ErrGradeAlreadyGiven
		return
	}

	return db.Save(&grade).Error
}

func GetGradeByStudentIDAndSubjectID(db *gorm.DB, studentID int, subjectID int) (grade Grade, err error) {
	err = db.First(&grade, &Grade{StudentID: studentID, SubjectID: subjectID}).Error
	return
}

func GetStudentsGrades(db *gorm.DB, id int) (grades []StudentGrade, err error) {
	err = db.Table("student").
		Select("student_subject.student_id, student.first_name, student.last_name, subject.name, student_subject.subject_id, grade.rank").
		Joins("JOIN student_subject ON student.id = student_subject.student_id").
		Joins("JOIN subject ON subject.id = student_subject.subject_id").
		Joins("JOIN grade ON student.id = grade.student_id AND subject.id = grade.subject_id").
		Where(&student.Student{ID: id}).Scan(&grades).Error

	for _, v := range grades {
		fmt.Println(v)
	}
	return
}

func GetGradeByStudentID(db *gorm.DB, ID int) (grades []Grade, err error) {
	err = db.Model(&Grade{}).Where(&Grade{StudentID: ID}).Find(&grades).Error
	return
}

func GetGradeBySubjectID(db *gorm.DB, ID int) (grades []Grade, err error) {
	err = db.Model(&Grade{}).Where(&Grade{SubjectID: ID}).Find(&grades).Error
	return
}

func UpdateGrade(db *gorm.DB, grade Grade) (err error) {
	return db.Model(&Grade{}).Where(&Grade{StudentID: grade.StudentID, SubjectID: grade.SubjectID}).Update(grade).Error
}

func DeleteGrade(db *gorm.DB, id int) (err error) {
	return db.Delete(&Grade{ID: id}).Error
}
