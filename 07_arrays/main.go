package main

import "fmt"

//In golang arrays are very less used DS
func main() {
	fmt.Println("GO Program to understand the arrays")

	//Declaring the arrays
	var sportsList [4]string

	sportsList[0] = "Carrom"
	sportsList[1] = "Football"
	sportsList[3] = "Cricket"

	fmt.Println("The sports list is: ", sportsList)
	fmt.Println("The len of sports list is: ", len(sportsList))

	var vegList = [4]string{"potato", "brinjal", "onion", "peas"}
	fmt.Println("The veg list is :", vegList)
}
