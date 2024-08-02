package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux" // go install github.com/gorilla/mux
)

type APIServer struct {
	store Store
	addr  string
}

func NewAPIServer(addr string, store Store) *APIServer {
	return &APIServer{
		addr:  addr,
		store: store,
	}
}

func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// registering our services
	taskService := NewTaskService(s.store)
	taskService.RegisterRoutes(subrouter)

	log.Println("Starting server on http://localhost" +
		s.addr + "/api/v1")

	log.Fatal(http.ListenAndServe(s.addr, subrouter))

}
