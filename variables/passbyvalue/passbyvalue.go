package main

import (
	"fmt"
)

func main() {
	course := "Getting started with GO"

	fmt.Println("\n Course name before update is ", course)
	updateCourseName(&course)
	fmt.Println("\n Course name after calling update function is ", course)
}

func updateCourseName(course *string) *string {
	fmt.Println("\n input to update function is ", *course)
	*course = "Docker and Kubernetes"
	fmt.Println("\n course name is now ", *course)
	return course
}
