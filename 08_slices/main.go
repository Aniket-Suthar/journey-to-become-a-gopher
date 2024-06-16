package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("GO Program to understand slices")

	//Creating the slice
	var sportList = []string{} //No memory limitation is present
	fmt.Printf("The type of DS is %T \n", sportList)

	//Adding data using append
	sportList = append(sportList, "Carrom", "TT", "Football", "Hockey")
	fmt.Println("The sports list is:", sportList)

	//Adding the sliced slice to existing slice
	sportList = append(sportList[1:4])
	fmt.Println("The sports list is:", sportList)

	//Initializing slices using mem mgmt operators
	highScores := make([]int, 5)
	highScores[0] = 34
	highScores[1] = 23
	highScores[2] = 2342
	highScores[4] = 3344
	// highScores[5] = 324 //this will give error of index out of bound

	fmt.Println("The high scores ", highScores)

	highScores = append(highScores, 345, 23423, 545) //no errors as append re-initializes the memory
	fmt.Println("The new high scores ", highScores)

	sort.Ints(highScores)
	fmt.Println("The sorted high scores ", highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//Removing the element from the index in slices
	var iplTeamsInFinal = []string{"CSK", "MI", "SRH", "RR", "KKR", "RCB"}
	fmt.Println("The current slice is", iplTeamsInFinal)

	//removing the value at index 2 i.e. "SRH"
	index := 2
	iplTeamsInFinal = append(iplTeamsInFinal[:index], iplTeamsInFinal[index+1:]...)
	fmt.Println("The updated slice is", iplTeamsInFinal)

}
