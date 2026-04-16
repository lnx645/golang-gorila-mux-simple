package api

import (
	"dtech.com/traktirku/internal/handler"
	"github.com/gorilla/mux"
)

func NewApi(s *mux.Router) {
	auth := handler.NewAuth()
	s.HandleFunc("/login", auth.Login).Methods("GET")
}
