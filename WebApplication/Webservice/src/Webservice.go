package main

//tutorial reference : https://tutorialedge.net/golang/creating-restful-api-with-golang/
//https://tutorialedge.net/golang/creating-simple-web-server-with-golang/

import (
	"fmt"
	"log"
	"net/http"
)

func homepage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome to home page")
	fmt.Print("Endpoint hit:homepage")
}

func handleRequests() {

	http.HandleFunc("/", homepage)
}

func handleRequestsWithParams() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "hello")
		fmt.Print("Endpoint hit:hello")

	})
}

func main() {
	handleRequests()
	handleRequestsWithParams()
	log.Fatal(http.ListenAndServe(":80", nil))

}
