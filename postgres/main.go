package main

import (
	"database/sql"
	"fmt"

	"os"

	_ "github.com/lib/pq"
)

type student struct {
	Id   int
	Name string
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

var db *sql.DB

func insert() {
	var (
		id   int
		name string
	)
	fmt.Println("Enter student ID")
	fmt.Scanln(&id)
	fmt.Println("Enter name of the student")
	fmt.Scanln(&name)
	_, err := db.Exec("INSERT INTO studentrecord(id,name)values($1,$2);", id, name)
	if err != nil {
		panic(err)
	}
}
func update() {
	var (
		id     int
		name   string
		choice int
	)
	fmt.Println("Enter 1 if you want to search by ID or enter 2 if you want to Search by Name")
	fmt.Scanln(&choice)
	if choice == 1 {
		fmt.Println("Enter the student Id")
		fmt.Scanln(&id)
		fmt.Println("Enter the updated name")
		fmt.Scanln(&name)
		_, err := db.Exec("UPDATE studentrecord SET name=$2 WHERE ID=$1", id, name)
		if err != nil {
			panic(err)
		}
	} else if choice == 2 {
		fmt.Println("Enter the student name")
		fmt.Scanln(&id)
		fmt.Println("Enter the updated id")
		fmt.Scanln(&name)
		_, err := db.Exec("UPDATE studentrecord SET id=$1 WHERE name=$2", id, name)
		if err != nil {
			panic(err)
		}

	}

}
func delete() {
	var id int
	fmt.Println("Enter the ID of student which you want to delete the record")
	fmt.Scanln(&id)
	_, err := db.Exec("DELETE FROM studentrecord record where ID=$1", id)
	if err != nil {
		panic(err)
	}
}
func show() {
	row, err := db.Query("SELECT * FROM studentrecord;")
	if err != nil {
		panic(err)
	}

	std := make([]student, 0)

	for row.Next() {
		st := student{}
		err := row.Scan(&st.Id, &st.Name)
		if err != nil {
			panic(err)
		}
		std = append(std, st)
	}
	if row.Err(); err != nil {
		panic(err)
	}
	for _, st := range std {
		fmt.Println(st.Id, st.Name)
	}

}

func main() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:Shantanu17@@localhost/abcschool?sslmode=disable")

	CheckError(err)
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("connected to Database")
	for {
		fmt.Println(" \n ABC school Mangement system")
		fmt.Println("1.Add student \n 2.Delete student \n 3.Update the details of Student\n 4.Show details of student \n 5.For exit program \n ")
		var choice int
		fmt.Printf("Choose Option(in Digit)")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			insert()
		case 2:
			delete()
		case 3:
			update()
		case 4:
			show()
		case 5:
			os.Exit(0)
		default:
			fmt.Println("option does not exist")

		}

	}

}
