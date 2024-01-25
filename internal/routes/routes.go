package routes

import (
	"fmt"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// Define API endpoints
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/api/data", apiDataHandler)
	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Index Endpoint")
}

func apiDataHandler(w http.ResponseWriter, r *http.Request) {

}
