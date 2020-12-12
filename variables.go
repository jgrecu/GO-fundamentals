package main

import (
	"fmt"
)

//var (
//	name	= "Nigel"
//	course	= "Docker Deep Dive"
//	module	= 3.2
//)

func main() {
	name	:= "Nigel"
	course	:= "Docker Deep Dive"

	fmt.Println("\nHi", name, "you;re currently watching",
		course)

	changeCourse(course)
	fmt.Println("\nYou are watching course", course)
}

func changeCourse(course string) string {
	course = "First Look: Native Docker Clustering"

	fmt.Println("\nTrying to change your course to", course)

	return course
}
