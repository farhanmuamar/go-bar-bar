package main

import (
	"net/http"
	"appsql/Personal"
)

func main() {

	http.HandleFunc("/", personal.ReadPersonal)
	http.ListenAndServe(":8080", nil)
	
}