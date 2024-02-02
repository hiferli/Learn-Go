package main

import "fmt"

type Student struct {
	Name       string
	ID         uint32
	Department string
	Present    bool
}

func (student Student) markPresent() {
	student.Present = true
}

type Teacher interface {
	markSelfPresent()
	markStudentPresent()
}

func main() {
	var A Student = Student{
		Name: "A",
		ID:   1,
	}

	var B Student = Student{"B", 2, "CS", true}

	fmt.Println(A) // {A 1  false}
	fmt.Println(B) // {B 2 CS true}

	var C = struct {
		Name       string
		Age        uint32
		Occupation string
	}{"C", 30, "Teacher"}

	fmt.Println(C) // {C 30 Teacher}
}
