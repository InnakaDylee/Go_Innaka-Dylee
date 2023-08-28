package main

import (
	"fmt"
)

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 {
	var average float64
	for _, value := range s.score{
		average += float64(value)
	}
	average = average/float64(len(s.score))

	return average
}

func (s Student) Min() (min int, name string) {
	studentName := ""
	minScore := s.score[0]

	for index, value := range s.score{
		if value < minScore {
			minScore = value
			studentName = s.name[index]
		}
	}
	return minScore, studentName
}

func (s Student) Max() (max int, name string) {
	studentName := ""
	maxScore := s.score[0]

	for index, value := range s.score{
		if value > maxScore {
			maxScore = value
			studentName = s.name[index]
		}
	}
	return maxScore, studentName
}

func main() {
	var a = Student{}

	for i := 1; i < 6; i++ {
		var name string
		fmt.Print("Input  " , i , " Student’s Name :  ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("Input  " + name + " Score :  ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\n\nAverage Score Students is ", a.Average())

	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is :  "+nameMax+" (", scoreMax, ")")

	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is :  "+nameMin+" (", scoreMin, ")")
}