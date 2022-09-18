package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	name := "Anirudh Krishna"
	course := "Getting started with GO"
	module := "4"
	clip := 2

	fmt.Println("Name and course are set to", name, "and", course, ".")
	fmt.Println("Modules and clip are set to", module, "and", clip, ".")
	fmt.Println("Name is of type", reflect.TypeOf(name))
	fmt.Println("Module is of type", reflect.TypeOf(module))

	iModule, err := strconv.Atoi(module)
	if err == nil {
		total := iModule + clip
		fmt.Println("Total is ", total)
	}

	fmt.Println("memory address of course variable", &course)

	var ptr *string = &course

	fmt.Println("Pointing course variable at address=", ptr, "which holds this value=", *ptr)
	// ptr = pointer variable holding the memory address of the course variable
	//Adding an *before a pointer variable returns the value storedin the variable bein pointed to
}
