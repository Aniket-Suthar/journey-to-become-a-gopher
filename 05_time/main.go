package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("GO Program to understand time package \n")

	currentTime := time.Now()

	fmt.Println("Formatted time is :", currentTime.Format("01-02-2006 15:04:05 Monday "))
	fmt.Println("Current time is: ", currentTime)

	//Creating the date/time
	createDate := time.Date(2002, time.September, 24, 12, 23, 33, 343, time.UTC)
	fmt.Println("Created date is :", createDate)

	numOfCpu := runtime.NumCPU()
	readTrace := runtime.ReadTrace()
	fmt.Println("The no. cpu usage is:", numOfCpu)
	fmt.Println("The trace is:", readTrace)

}
