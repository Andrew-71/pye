package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("=== Working on port 7102 ===")

	router := http.NewServeMux()

	router.HandleFunc("GET /pem", publicKey)

	router.HandleFunc("POST /register", Register)
	router.HandleFunc("POST /login", Login)

	router.HandleFunc("GET /login", Login) // TODO: temp

	http.ListenAndServe(":7102", router)
}
