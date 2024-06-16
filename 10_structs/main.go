package main

import "fmt"

func main() {
	fmt.Println("GO Program to understand struct")

	//no inheritance and super in golang

	//creating struct
	aniket := User{"aniket", 21, "aniket@go.dev.com", 3459352243}
	fmt.Println(aniket)
	fmt.Printf("The details are %+v", aniket)
	fmt.Printf("The Name is %v and email is %v.", aniket.Name, aniket.Email)

}

type User struct {
	Name   string
	Age    int
	Email  string
	Mobile int
}
