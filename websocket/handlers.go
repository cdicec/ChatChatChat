package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func handlerDefault(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	http.Error(w, "Not found", 404)
	return
}

func handlerAuth(w http.ResponseWriter, req *http.Request) {
	// accepts json
	buf := new(bytes.Buffer)
	buf.ReadFrom(req.Body)

	var dat map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &dat); err != nil {
		panic(err)
	}
	password := dat["password"].(string)
	fmt.Println("User password is:")
	fmt.Println(password)
	fmt.Fprintf(w, "Your password is: %s", password)
}

func handlerReg(w http.ResponseWriter, req *http.Request) {
	// accepts json
	buf := new(bytes.Buffer)
	buf.ReadFrom(req.Body)

	var dat map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &dat); err != nil {
		panic(err)
	}
	password := dat["password"].(string)
	fmt.Println("User password is:")
	fmt.Println(password)
	fmt.Fprintf(w, "Your password is: %s", password)
}
