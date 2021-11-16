package main

import (
	"fmt"
	"os"
)

type student struct {
	ID      int
	Name    string
	FName   string
	Admyear int
}

func Newstd(id int, name string, fname string, admyear int) *student {
	return &student{id, name, fname, admyear}
}

type studentdata struct {
	Sdata map[int]*student
}

func (s studentdata) showdata() {
	for _, stu := range s.Sdata {

		fmt.Printf("StudentId:%d\n Student Name:%s\n Student's Father Name:%s\n Student admission Year:%d", stu.ID, stu.Name, stu.FName, stu.Admyear)
	}

}
func (s studentdata) addstd() {
	var (
		id      int
		name    string
		fname   string
		admyear int
	)
	fmt.Println("Enter the Student ID")
	fmt.Scanln(&id)
	fmt.Println("Enter the Name of the student")
	fmt.Scanln(&name)
	fmt.Println("Enter the fathers name of student")
	fmt.Scanln(&fname)
	fmt.Println("Enter the year of student admission")
	fmt.Scanln(&admyear)

	if _, ok := s.Sdata[id]; !ok {
		newstd := Newstd(id, name, fname, admyear)
		s.Sdata[id] = newstd
		fmt.Printf("New Student added with ID:%d \n Name:%s \n Fathers Name:%s \n Admission Year:%d", id, name, fname, admyear)
	} else {
		fmt.Println("Record already Exist")
	}

}

func (s studentdata) deletestu() {
	var id int
	fmt.Println("Enter the ID of student which you want or erase data")
	fmt.Scanln(&id)

	if _, ok := s.Sdata[id]; !ok {
		fmt.Println("No student exist of this record")
	} else {
		delete(s.Sdata, id)
		fmt.Printf("Successfully deleted the record of the student ID %d", id)

	}
}
func (s studentdata) upadatestd() {
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
		if _, ok := s.Sdata[id]; !ok {
			fmt.Println("No Record found of this Particular id")
		} else {
			var (
				id      int
				name    string
				fname   string
				admyear int
			)
			fmt.Println("Enter the Student ID")
			fmt.Scanln(&id)
			fmt.Println("Enter the Name of the student")
			fmt.Scanln(&name)
			fmt.Println("Enter the fathers name of student")
			fmt.Scanln(&fname)
			fmt.Println("Enter the year of student admission")
			fmt.Scanln(&admyear)
			newstd := Newstd(id, name, fname, admyear)
			s.Sdata[id] = newstd
			fmt.Printf("\nStudent Updated with ID:%d \n Name:%s \n Fathers Name:%s \n Admission Year:%d", id, name, fname, admyear)
		}

	} else if choice == 2 {
		fmt.Println("Enter the name of the student")
		fmt.Scanln(&name)
		for _, stu := range s.Sdata {
			if stu.Name == name {
				var (
					id      int
					name    string
					fname   string
					admyear int
				)
				fmt.Println("Enter the Student ID")
				fmt.Scanln(&id)
				fmt.Println("Enter the Name of the student")
				fmt.Scanln(&name)
				fmt.Println("Enter the fathers name of student")
				fmt.Scanln(&fname)
				fmt.Println("Enter the year of student admission")
				fmt.Scanln(&admyear)
				newstd := Newstd(id, name, fname, admyear)
				s.Sdata[id] = newstd
				fmt.Printf("Student Updated with ID:%d \n Name:%s \n Fathers Name:%s \n Admission Year:%d", id, name, fname, admyear)

			} else {
				fmt.Println("No such data found")
			}
		}

	}

}

func main() {
	Sdata := studentdata{
		make(map[int]*student, 200)}

	for {
		fmt.Println(" \n ABC school Mangement system")
		fmt.Println("1.Add student \n 2.Delete student \n 3.Update the details of Student\n 4.Show details of student \n 5.For exit program \n ")
		var choice int
		fmt.Printf("Choose Option(in Digit)")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			Sdata.addstd()
		case 2:
			Sdata.deletestu()
		case 3:
			Sdata.upadatestd()
		case 4:
			Sdata.showdata()
		case 5:
			os.Exit(0)
		default:
			fmt.Println("option does not exist")

		}

	}
}
