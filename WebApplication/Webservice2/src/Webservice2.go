package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var counter int

//This is just a much cleaner way of calling a service and creating multiple end points
func main() {

	//one function for the homepage
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "HOMEPAGE1")
	})

	//will need to make this work
	http.Handle("/static", http.FileServer(http.Dir("/index.html")))

	//calling the other end points
	http.HandleFunc("/hello", echoString)
	http.HandleFunc("/users", incrementCounter)

	//Finally listen and serve on a port

	log.Fatal(http.ListenAndServe(":80", nil))
}

func echoString(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(writer, "hello")
}

//this is not thread safe and should be made thread safe.
//as of now this is only for testing
func incrementCounter(writer http.ResponseWriter, request *http.Request) {
	counter = counter + 1
	fmt.Fprintf(writer, strconv.Itoa(counter))
}
