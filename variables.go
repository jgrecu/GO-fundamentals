package main

import (
	"fmt"
	"reflect"
)

//var (
//	name	= "Nigel"
//	course	= "Docker Deep Dive"
//	module	= 3.2
//)

func main() {
	name	:= "Nigel"
	//course	:= "Docker Deep Dive"
	module	:= 3.2
	fmt.Println("Name is ", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Module is ", module, "and is of type", reflect.TypeOf(module))
}
