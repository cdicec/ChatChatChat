package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://127.0.0.1:6543/write"
	fmt.Println("URL:>", url)

	req, _ := http.NewRequest("POST", url, bytes.NewBufferString("Hi!"))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "text/html")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
