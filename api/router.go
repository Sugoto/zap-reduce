package api

import (
	"github.com/gorilla/mux"
)

// NewRouter creates a new router for the API.
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/process", ProcessText).Methods("POST")
	return router
}
