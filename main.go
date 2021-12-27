package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/file", FileRequest)

	log.Println("listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
