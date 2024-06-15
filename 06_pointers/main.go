package main

import "fmt"

func main() {
	fmt.Println("GO Program to understand pointers")

	var ptr *int
	fmt.Println("Value of default pointer is: ", ptr)

	//Creating pointers which points to something
	newNum := 23423

	var newPtr = &newNum                                         //creating and referencing the pointer to var "newNum"
	fmt.Println("Mem adddress ref. by the pointer is: ", newPtr) //returns the mem address
	fmt.Println("Value of the pointer is: ", *newPtr)            //returns the actual data

	*newPtr = *newPtr * 3
	fmt.Println("Value of the pointer is: ", *newPtr) //returns the actual data

}
