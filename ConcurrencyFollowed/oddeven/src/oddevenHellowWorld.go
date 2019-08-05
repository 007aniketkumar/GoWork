package main

import (
	"fmt"
	"time"
)

func main() {

	//let us introduce sleep so that other functions get time to execute

	//declaring anonymous functions
	//one prints Hello
	//making it concurrent

	go func() { //adding go will create co routines

		for i := 0; i < 100; i = i + 2 {
			fmt.Println("hello")
			fmt.Println(i)
			time.Sleep(1 * time.Millisecond)

		}
	}()

	go func() {
		for i := 1; i < 100; i = i + 2 {
			fmt.Println("world")
			fmt.Println(i)
			time.Sleep(1 * time.Millisecond)

		}
	}()
	time.Sleep(2 * time.Second)

}
