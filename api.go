package main

import (
	"log"
	"net/http"

)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{addr: addr}
}

func (s *APIServer) Start() error {
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	middlewareChain := MiddlewareChain(RequestLoggerMiddleware, RequireAuthMiddleware)

	server := http.Server{
		Addr:    s.addr,
		Handler: middlewareChain(router),
	}

	log.Println("Server is starting on", s.addr)
	return server.ListenAndServe()
}

func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("method =", r.Method, "url =", r.URL)
		next.ServeHTTP(w, r)
	})
}

func RequireAuthMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request has the correct authentication token
		if r.Header.Get("Authorization") != "Bearer my_secret_token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

type Middleware func(http.Handler) http.HandlerFunc

func MiddlewareChain(middlewareargs ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		for i := len(middlewareargs) - 1; i >= 0; i-- {
			next = middlewareargs[i](next)
		}
		return next.ServeHTTP
	}
}
