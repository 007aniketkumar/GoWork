package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("hello world")

	testingIfElseFun()

	fmt.Println("The exception propagated is:", testingErrorHandling())

	//testing loops
	printNumbersInLoop()

	//testing maps
	funWithMaps()

}

func testingIfElseFun() {

	//you can contrain the variables only for the if else block , something like this

	if marksinMaths, marksinEnglish := 80, 70; marksinMaths > marksinEnglish {
		fmt.Println("Good in Maths")
	} else {
		fmt.Println("Good in English")
	}
}

//error handling
func testingErrorHandling() error {

	_, err := os.Open("/usr/local/test.txt") //more like try catch

	if name, address := "Aniket", "Bangalore"; name == address {
		fmt.Println("names are same")
	} else if err == nil {
		fmt.Println("file found")
	} else {
		fmt.Println("file not found")
	}

	return err
}

//Looping
//slice is pass by reference-
// Slices are internally build on top of array , and slice are pointers to the underlying array. So any change in slice
//results in change overall
func printNumbersInLoop() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//declare a slce -which is like an unordered list
	//they can also be created using the make command
	names := []string{"hello", "world", "endofWorld"}

	//same as above, make is special command
	//names := make([]string, 3, 3)

	//slice of the above slice
	subNames := names[0:2] // include 0,1 and exclude 2

	//testing range : it returns 2 variables index and value
	for indx, name := range names {
		fmt.Println("item ", name, "is present at", indx)
	}

	//printing the sub slice
	for _, sbName := range subNames {
		fmt.Println("sub slice content:", sbName)
	}

	//appending slices: as long as they are of the same type
	newSlice := []string{"moon", "earth"}

	//let's append the names to newSlice

	names = append(names, newSlice...)
	fmt.Println("post appending names: ", names)

}

//let's checkout maps
func funWithMaps() {
	myMap := map[string]int{"aniket": 1, "aman": 2, "gyan": 3}

	//iterating over the map using the for range loop
	for mapKeys, mapvalues := range myMap {
		fmt.Println("Printing the mapvalues", mapvalues, "at keys", mapKeys)
	}

	//to update the existing map
	myMap["Sohan"] = 4

	//to delete an entry in a map

	delete(myMap, "aman")

	fmt.Println("The updated map is", myMap)

}
