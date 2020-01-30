package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rmelo/httpserver/handler"
	lib "github.com/rmelo/httpserver/http"
)

func routes() http.Handler {

	r := mux.NewRouter()

	r.Use(lib.WithCorrelationID)
	r.Use(lib.WithLogging)

	memberHandler := lib.WithDecode(NewMember, handler.Member)
	timeHandler := lib.WithDecode(NewTime, handler.Time)

	r.Handle("/time", timeHandler)
	r.Handle("/member", memberHandler)
	r.HandleFunc("/hello", handler.Hello)

	return r
}
