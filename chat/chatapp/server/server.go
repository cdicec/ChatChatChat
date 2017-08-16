package main

import(
	"net/http"
	"fmt"
)

func redirectToHttps(w http.ResponseWriter, req *http.Request) {
    // http.Redirect(w, req, "https://"+ host + ":" + port + req.RequestURI, http.StatusMovedPermanently)
    http.Redirect(w, req, "https://127.0.0.1:3456" + req.RequestURI, http.StatusMovedPermanently)
}

func main() {
	/*
	старт сервера - слушает + принимает соединения
	при  accept  уходит в join_chat
	*/
	http.HandleFunc("/", handler)
	go http.ListenAndServeTLS(":3456", "cert.pem", "key.pem", nil)
	http.ListenAndServe(":6543", http.HandlerFunc(redirectToHttps))
}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
    fmt.Fprintf(w, "Hi there!")
}

// func join_chat(conn *net.Conn) {
	
// 		авторизация - запрос никнейма -> создание структуры Client
// 		присоединение к главной комнате
// }





