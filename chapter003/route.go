package main

import "net/http"

func routes(){
	http.HandleFunc("/", handleIndex)
}
