package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

func urlSolver(room *Room, w http.ResponseWriter, req *http.Request) {
	fmt.Println("Requested url is: ", req.URL.Path)
	switch req.URL.Path {
	case "/":
		http.ServeFile(w, req, "home.html")
	case "/reg":
		handlerReg(w, req)
	case "/auth":
		handlerAuth(w, req)
	case "/ws":
		serveWs(room, w, req)
	default:
		handlerDefault(w, req)
	}
}

func main() {
	flag.Parse()
	room := newRoom()
	go room.run()
	err := http.ListenAndServe(*addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urlSolver(room, w, r)
	}))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
