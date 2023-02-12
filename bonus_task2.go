package main

import (
	"fmt"
	"strconv"
)

var Professions []Product
var Students []User

type Student struct {
	ID       int
	FullName string
}

type Profession struct {
	ID     int
	Faculty string
	ProfName int
}

func AddUser(name string) {
	student := Student{len(Students), name}
	Students = append(Students, student)
}

func AddProduct(faculty string, ProfName int) {
	profession := Profession{len(Professions), faculty, ProfName}
	Professions = append(Professions, profession)
}

func (user Student) Print() {
	fmt.Println("ID: " + strconv.Itoa(student.ID) + ", Full Name: " + student.FullName)
}

func (product Profession) Print() {
	fmt.Println("ID: " + strconv.Itoa(profession.ID) + ", Faculty: " + profession.Faculty)
}
