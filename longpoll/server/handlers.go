package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func handlerRead(w http.ResponseWriter, req *http.Request) {
	listener := make(chan string)
	listeners = append(listeners, listener)
	messageText := <-listener
	// messageText := <-messages
	fmt.Fprintf(w, "%s", messageText)
}

func handlerWrite(w http.ResponseWriter, req *http.Request) {
	// accepts raw body data. Cookie validation required
	buf := new(bytes.Buffer)
	buf.ReadFrom(req.Body)
	messageText := buf.String()
	// messages <- messageText
	for _, listener := range listeners {
		listener <- messageText
	}
	fmt.Fprintf(w, "Your message \"%s\" is on server now.", messageText)
}

func handlerRegistration(w http.ResponseWriter, req *http.Request) {
	// accepts json. See ../client/client_reg
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

func handlerDefault(w http.ResponseWriter, req *http.Request) {
	result := "Answer from default handler"
	fmt.Fprintf(w, "%s", result)
}
