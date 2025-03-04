package Routes

import (
	"crud2/controllers"
	"databse/sql"
	"fmt"
	"os"
)

func SetupRoutes(db *sql.DB){
	for{
		fmt.Println("Choose the Operation you want to perform: \n 1. Read Data \n 2. Update Data \n 3. Delete Data \n 4. Insert Data \n 5. Exit")
	    var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			controllers.ReadStudents(db)
		case 2:
			controllers.UpdateStudent(db)
		case 3:
			controllers.DeleteStudent(db)
		case 4:
			controllers.InsertStudent(db)
		case 5:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}