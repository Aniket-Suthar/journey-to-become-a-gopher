package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "GO Program for type conversions"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your message for type conversion :")
	input, err := reader.ReadString('\n')

	//type conversion
	increasedVal, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Converted val :", increasedVal+1) // will give error as conversion will include "\n" as well hence trimming should be done
	}

}
