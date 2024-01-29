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
	mux.HandleFunc("/api/list", handleList)
	mux.HandleFunc("/api/add", handleAdd)
	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Index Endpoint")
}

func handleList(w http.ResponseWriter, r *http.Request) {
	data := "LIST endpoint accessed"
	fmt.Fprint(w, data)
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	data := "ADD endpoint accessed"
	fmt.Fprint(w, data)
}
