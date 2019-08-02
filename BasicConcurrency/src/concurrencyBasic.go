package main

import (
	"fmt"
	"sync"
	"time"
)

//the program is made concurrent by adding go coroutines
//this is achieved by adding the keyword go
//what happens after that , is main function completes its executions and exists, there by exiting all the other go co routines
//ideally it should be the last to exit
//what we do now is :
// we add a type of thread safe counter using sync.WaitGroup
//so keep waiting , until all coroutines are complete.
//Some concept  : the coroutines run on top of OS threads (unlike java). So , context switch between coroutines
//is not expensive

func main() {

	var waitFuncs sync.WaitGroup

	waitFuncs.Add(2)
	fmt.Println("in the main function")

	//declaring an anonymous function to print hello

	go func() {
		//signal that this is complete
		defer waitFuncs.Done()

		time.Sleep(5 * time.Second)
		fmt.Println("Hello")

	}()

	//declaring anonymous function to print World

	go func() {

		defer waitFuncs.Done()

		fmt.Println("World")
	}()

	//wait till the coroutines are complete.
	waitFuncs.Wait()
}
