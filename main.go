package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("=== PYE ===")

	router := http.NewServeMux()

	router.HandleFunc("GET /public-key", publicKey)

	router.HandleFunc("POST /register", Register)
	router.HandleFunc("POST /login", Login)

	http.ListenAndServe(":7102", router)
}
