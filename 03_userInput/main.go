package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "GO Program to take user input"
	fmt.Println(welcome)

	//Creating the reader from bufio package
	reader := bufio.NewReader(os.Stdin)

	//Comma ok | error ok syntax - putting the reader data
	input, _ := reader.ReadString('\n') // if not handling the errors then "_" is used
	fmt.Println("Your message is :", input)
	fmt.Printf("Type of input msg is %T", input) //type is string
}
