package controllers
import (
    "database/sql"
	"fmt"
	"crud2/models"
	"crud2/services"
)

func ReadStudents(db *sql.DB) {
	students, err := services.GetAllStudents(db)
	if err != nil {
		fmt.Println("Error fetching students:", err)
		return
	}
	for _, student := range students {
		fmt.Printf("%d - %s - %s\n", student.RollNo, student.Name, student.Subj)
	}
}

func InsertStudent(db *sql.DB) {
	student := models.Student{
		RollNo: 21,
		Name:   "Shark",
		Subj:   "French",
	}
	err := services.InsertStudent(db, student)
	if err != nil {
		fmt.Println("Error inserting student:", err)
	} else {
		fmt.Println("Student inserted successfully!")
	}
}

func UpdateStudent(db *sql.DB) {
	err := services.UpdateStudent(db, 21, "German")
	if err != nil {
		fmt.Println("Error updating student:", err)
	} else {
		fmt.Println("Student updated successfully!")
	}
}

func DeleteStudent(db *sql.DB) {
	err := services.DeleteStudent(db, 16)
	if err != nil {
		fmt.Println("Error deleting student:", err)
	} else {
		fmt.Println("Student deleted successfully!")
	}
}