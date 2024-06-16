package main

import "fmt"

func main() {
	fmt.Println("GO Program for understanding maps")

	programmingLangs := make(map[string]string) //[key]value

	programmingLangs["C"] = "C language"
	programmingLangs["PY"] = "Python"
	programmingLangs["Js"] = "Javascript"
	programmingLangs["JV"] = "Java"

	fmt.Println("The map is", programmingLangs)
	fmt.Println("JS stands for: ", programmingLangs["JS"])

	//Deleting the values
	delete(programmingLangs, "C")
	fmt.Println("The map is", programmingLangs)

	//Loops in Golang
	for key, value := range programmingLangs {
		fmt.Printf("For key %v, the value is %v \n", key, value)
	}
}
