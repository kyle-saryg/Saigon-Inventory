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
	mux.HandleFunc("/api/list", apiDataHandler)
	mux.HandleFunc("/api/add", apiAddItem)
	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Index Endpoint")
}

func apiAddItem(w http.ResponseWriter, r *http.Request) {

}

func apiDataHandler(w http.ResponseWriter, r *http.Request) {
	data := "Some Data from the API"
	fmt.Fprint(w, data)
}
