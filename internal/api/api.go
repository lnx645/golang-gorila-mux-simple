package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewApi(s *mux.Router) {
	s.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("WKWKWK"))
	}).Methods("GET")
}
