package handler

import (
	"net/http"

	http_resp "dtech.com/traktirku/pkg/http"
)

type AuthApiHandler struct {
	name string
}

func NewAuth() *AuthApiHandler {
	return &AuthApiHandler{}
}

func (a *AuthApiHandler) Login(w http.ResponseWriter, r *http.Request) {
	response := http_resp.NewResponse(w, r)

	response.JSON()
}
