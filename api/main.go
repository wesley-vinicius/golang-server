package main

import (
	"github.com/software-engineer-challenge/api/handle"
	"net/http"
)

func main() {
	routes()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func routes() {
	http.HandleFunc("/search", handle.SearchUserHandle)
}
