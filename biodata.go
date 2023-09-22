package main

import (
	"assignment-1/model"
	"assignment-1/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run biodata.go [student_code]")
	}

	studentCode := os.Args[1]

	studentList, err := readFile()
	utils.CheckIfError(err, "Error Read File")
	
	student, err := checkStudent(studentCode, &studentList)
	utils.CheckIfError(err, "student not found")
	
	showStudentInformation(student)
	
}

func checkStudent(code string, studentList *model.StudentList) (model.Student, error){
	for _, student := range studentList.StudentList {
		if code == student.Code{
			return student, nil
		}
	}

	return model.Student{}, errors.New("Student not found")
}

func readFile() (model.StudentList, error){
	studentsJson, err := ioutil.ReadFile("data/students.json")
	utils.CheckIfError(err, "Error Read File students.json")

	var studentList model.StudentList
	err = json.Unmarshal(studentsJson, &studentList)
	utils.CheckIfError(err, "Error Unmarshal json")

	if len(studentList.StudentList) < 1 {
		return model.StudentList{}, errors.New("Student empty")
	}

	return studentList, nil
}

func showStudentInformation(student model.Student) {
	fmt.Printf("ID             : %s \n", student.ID)
	fmt.Printf("Name           : %s \n", student.Name)
	fmt.Printf("Address        : %s \n", student.Address)
	fmt.Printf("Job            : %s \n", student.Occupation)
	fmt.Printf("Joining Reason : %s \n", student.JoiningReason)
}