package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := router.NewRouter()

	port := 8000
	// Sprintf returns a string accompanying the specified formatter
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server listening on http://localhost%s\n", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}
