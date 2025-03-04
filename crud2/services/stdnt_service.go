package services

import (
	"crud2/models"
	"crud2/repositories"
	"database/sql"
)

func GetAllStudents(db *sql.DB)([]models.Student, error){
	return repositories.GetAllStudents(db)
}
func InsertStudent(db *sql.DB, student models.Student)error{
	return repositories.InsertStudent(db, student)
}
func UpdateStudent(db *sql.DB,rollno int, newSubj string)error{
	return repositories.UpdateStudent(db, rollno,newSubj)
}
func DeleteStudent(db *sql.DB, rollno int)error{
	return repositories.DeleteStudent(db, rollno)
}