package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ServerInstance struct {
	Port        string
	Host        string
	Router      *mux.Router
	V1ApiRouter *mux.Router
}

func (s *ServerInstance) NewServer() {
	s.Router = mux.NewRouter()
	s.V1ApiRouter = s.Router.PathPrefix("/v1/api").Subrouter()
}

func (s *ServerInstance) RunServer() {
	go func(s *ServerInstance) {
		fmt.Printf("🚀 Server siap di port %s\n", s.Port)
		if err := http.ListenAndServe(s.Port, s.Router); err != nil {
			log.Fatalf("Ada kesalahan saat jalanin server!: %s", err)
		}
		fmt.Println("Server Berjalan di port:")
	}(s)
	fmt.Println("Program utama terus berjalan...")
	select {}
}
