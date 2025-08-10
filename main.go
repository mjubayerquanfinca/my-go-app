package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mjubayerquanfinca/my-go-app/handler"
)

func main() {
	http.HandleFunc("/", handler.HelloHandler)

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
