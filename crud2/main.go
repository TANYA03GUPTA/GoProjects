package main

import (
	"crud2/database"
	"crud2/routes"

	_ "github.com/lib/pq"
	//"golang.org/x/text/cases"
)
func main(){
	db := database.ConnectDB()
	defer db.Close()
	routes.SetupRoutes(db)
}

//const (
//	host     = "localhost"
//	port     = 5432
//	user     = "postgres"
//	password = "2003"
//	dbname   = "second_db"
//)
//var name string 
//var subj string
//var rollno int
//
//func main(){
//	psqlconn := fmt.Sprintf("host = %s  port = %d user = %s password = %s dbname = %s sslmode=disable", host,port,user, password, dbname)
//	db,err := sql.Open("postgres", psqlconn)
//	CheckErr(err)
//	defer db.Close()
//
//	var choice int 
//	for {
//		fmt.Println("choose the operation you want to perform : \n 1. read Data \n 2.Update Data \n 3.Delete data \n 4.Insert Data")
//		fmt.Scan(&choice)
//		switch choice {
//		case 1: Read(db)
//		case 2 : Update(db)
//		case 3: Delete(db)
//		case 4 : Insert(db)
//		case 5: os.Exit(0)
//		}
//	}
//    
//}
//func Read(db *sql.DB){
//	readStm1t := `Select * FROM "Student"`
//	rows,err := db.Query(readStm1t)
//	if err != nil {
//		CheckErr(err)
//	}else{
//		for rows.Next() {
//			rows.Scan(&rollno, &name, &subj)
//			fmt.Printf("%d - %s - %s \n", rollno, name,subj)
//		   }
//		 
//	}
//}
//func Insert(db *sql.DB){
//	rollno =21
//	name = "Shark"
//	subj ="french"
//	fmt.Println("Insert called")
//	Inserttodb(db, rollno ,name, subj)
//}
//func Inserttodb(db *sql.DB, rollno int,name,subj string){
//	insrtst := `insert into "Student"(rollno,name,subj) values($1 ,$2, $3)`
//	_,er := db.Exec(insrtst,rollno,name,subj)
//	CheckErr(er)
//}
//func Update(db *sql.DB){
//	updtstmt := `UPDATE "Student" SET "subj" = $1 WHERE "rollno"=$2`
//    _,er3 := db.Exec(updtstmt,"German",21)
//    CheckErr(er3)
//	fmt.Println("Updates made !")
//
//}
//func Delete(db *sql.DB){
//	fmt.Println("delete !!!")
//	deleteStmt := `DELETE FROM "Student" WHERE "rollno" = $1`
//	_, err := db.Exec(deleteStmt, 16)
//	CheckErr(err)
//	fmt.Println("Deleted employee with EmpId 89")
//}
//func CheckErr(err error){
//	if err != nil{
//		panic(err)
//	}
//}
//
