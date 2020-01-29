package main

import (
	"log"
	"net/http"

	h "github.com/rmelo/httpserver/http"
)

func main() {

	server := startServer()

	log.Println("Listening...")
	http.ListenAndServe(":8080", server)
}

func startServer() http.Handler {
	mux := http.NewServeMux()
	s := h.Server{
		Router: routes(),
	}
	mux.Handle("/", s.Router)
	return mux
}
