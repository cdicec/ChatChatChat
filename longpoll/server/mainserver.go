package main

import (
	"fmt"
	"net/http"
)

// var messages chan string
var listeners []chan string

func urlSolver(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Requested url is: ", req.URL.Path)
	switch req.URL.Path {
	case "/reg":
		handlerRegistration(w, req)
	case "/read":
		handlerRead(w, req)
	case "/write":
		handlerWrite(w, req)
	default:
		fmt.Println("Default page")
		handlerDefault(w, req)
	}
}

func main() {
	// messages = make(chan string, 10)
	http.ListenAndServe(":6543", http.HandlerFunc(urlSolver))
}
