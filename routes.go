package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rmelo/httpserver/entity"
	"github.com/rmelo/httpserver/handler"
	common "github.com/rmelo/httpserver/http"
)

func routes() http.Handler {

	r := mux.NewRouter()
	r.Use(common.WithLogging)

	memberHandler := common.WithDecode(func() interface{} { return &entity.Member{} }, handler.Member)
	timeHandler := common.WithDecode(func() interface{} { return entity.NewTime() }, handler.Time)

	r.Handle("/time", timeHandler)
	r.Handle("/member", memberHandler)
	r.HandleFunc("/hello", handler.Hello)

	return r
}
