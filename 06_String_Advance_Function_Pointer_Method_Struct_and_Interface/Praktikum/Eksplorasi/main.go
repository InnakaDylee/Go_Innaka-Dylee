package main

import (
	"fmt"
	"strings"
)

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""
	// your code here
	input := strings.ToLower(string(s.name))
	for _, char := range input {
		if char >= 'a' && char <= 'z' {
			newChar := 'z' - (char - 'a')
			nameEncode += string(newChar)
		} else {
			nameEncode += string(char)
		}
	}

	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""
	// your code here
	input := strings.ToLower(string(s.name))
	for _, char := range input {
		if char >= 'a' && char <= 'z' {
			newChar := 'z' - (char - 'a')
			nameDecode += string(newChar)
		} else {
			nameDecode += string(char)
		}
	}

	return nameDecode
}

func main() {
	var menu int
	var a student = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student's name " + a.name + "is : " + c.Encode())

	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nDecode of student's name " + a.name + "is : " + c.Decode())

	}

}
