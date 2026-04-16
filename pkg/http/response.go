package http

import "net/http"

type Response struct {
	w http.ResponseWriter
	r *http.Request
}

func NewResponse(w http.ResponseWriter, r *http.Request) *Response {
	return &Response{
		w: w,
		r: r,
	}
}

func (w *Response) JSON() {
	w.w.Header().Set("content-type", "application/json")
	w.w.WriteHeader(http.StatusOK)
	w.w.Write([]byte("WKWKWKW"))
}
