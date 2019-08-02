package main

import (
	"fmt"
	"os"      // to access environment variables
	"reflect" //check the type of variable , looks similar to reflection in java
	"runtime" //to get the run time details
)

//declaring variables outside a program
var (
	name   string //initialised to null
	rollNo int    // initialised to 0 by default

	//declare 2 variables of same type in the same line

	name2, address string

	//dont declare the type itself, go checks the value and declares it for you
	name3 = "Aniket"
)

//defining constants using the const keyword
const name4 = "Do not change me"

func main() {

	//local variable declaration :=
	localName := "kumar"
	//using the local variable
	fmt.Println("the local variable is", localName)

	fmt.Println("hello world", runtime.GOOS) //prints the name of the OS

	fmt.Println(name, "The initalised name is", name3)

	//type of variable where the type is not explicitly defined
	fmt.Println("The type of name3 is", reflect.TypeOf(name3))

	/**

	Using pointers

	same logic as C

	& - will print the address
	* - will print the address
	* The below segment will print

	The address of the name3 variable is 0x116b3c0 The corresponding value is Aniket


	*/

	addressofName := &name3

	fmt.Println("The address of the name3 variable is", addressofName,
		"The corresponding value is", *addressofName)

	testPassbyValue(name3)
	fmt.Println("The value of name3  after call by value is", name3)

	testPassbyReference(&name3)
	fmt.Println("The value of name3  after call by reference is", name3)

	//play with environment varibles

	playWihtEnviromentVar()
}

// Test the pass by value piece
//note the variables names can be different , and need not be same as the calling class
func testPassbyValue(name3 string) string {

	name3 = "Name changed to Aman" // no : here , as I am using an existing variable

	fmt.Println(name3)
	return name3

}

//Test pass by reference
func testPassbyReference(pointerofName3 *string) string {
	*pointerofName3 = "name changed to Gyan"
	fmt.Println(*pointerofName3)
	return *pointerofName3
}

//Playing with Environment variables
// import "os"
func playWihtEnviromentVar() {

	fmt.Println(os.Environ()) //prints a list of environment variables
	fmt.Println("Get the current logged in ", os.Geteuid())

}

//varargs

func varagrs(s ...int) {

}
