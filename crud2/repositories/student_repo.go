package repositories


import (
	"database/sql"
	"crud2/models"
)

func GetAllStudents(db *sql.DB) ([]models.Student, error) {
	rows, err := db.Query(`SELECT rollno, name, subj FROM "Student"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var student models.Student
		err := rows.Scan(&student.RollNo, &student.Name, &student.Subj)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}

func InsertStudent(db *sql.DB, student models.Student) error {
	query := `INSERT INTO "Student" (rollno, name, subj) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, student.RollNo, student.Name, student.Subj)
	return err
}

func UpdateStudent(db *sql.DB, rollNo int, newSubj string) error {
	query := `UPDATE "Student" SET subj = $1 WHERE rollno = $2`
	_, err := db.Exec(query, newSubj, rollNo)
	return err
}

func DeleteStudent(db *sql.DB, rollNo int) error {
	query := `DELETE FROM "Student" WHERE rollno = $1`
	_, err := db.Exec(query, rollNo)
	return err
}
