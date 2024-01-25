package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() http.Handler {
	mux := mux.NewRouter()

	// Define API endpoints
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/api/data", apiDataHandler)
	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Index Endpoint")
}

func apiDataHandler(w http.ResponseWriter, r *http.Request) {
	data := "Some Data from the API"
	fmt.Fprint(w, data)
}
