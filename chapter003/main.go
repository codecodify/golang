package main

import (
	"log"
	"net/http"
)

func main() {
	routes()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
