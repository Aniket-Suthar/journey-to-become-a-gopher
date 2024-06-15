package main

import "fmt"

//Declaring constants
const LoginToken string = "hellologin" //NOTE: first letter caps in go means it is publicly accessed

func main() {
	//Anything declared in go should be used else gives an error
	var username string = "Aniket"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedin bool = true
	fmt.Println(isLoggedin)
	fmt.Printf("Variable is of type: %T \n", isLoggedin)

	var val int = 256
	fmt.Println(val)
	fmt.Printf("Variable is of type: %T \n", val)

	var smallVal uint = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.24356435767567 //more precision than float32
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//Default values and aliases
	var anotherValue int
	fmt.Println(anotherValue)

	var anotherString string
	fmt.Println("Default of value of string is:", anotherString)

	//Implicit type - decided by lexer
	var newVar = "helloaniket.in"
	fmt.Printf("Implicit type is %T :", newVar)

	//Walrus operator - can only be used inside the method, but for global declaration var should be used
	numOfUser := 200
	fmt.Println(numOfUser)

	//Calling the const
	fmt.Println("Login token is: ", LoginToken)

}
