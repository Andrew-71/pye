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

	// Note: likely temporary, possibly to be replaced by a fake "frontend"
	router.HandleFunc("GET /login", Login) 
	router.HandleFunc("GET /register", Register)

	http.ListenAndServe(":7102", router)
}
