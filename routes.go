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

	jsonChain := func(f func() interface{}, next common.JSONHandler) http.Handler {
		return common.LogMiddleware(common.JSONHandle(f, next))
	}

	r.Handle("/time", jsonChain(entity.NewTime, handler.Time))
	r.Handle("/member", common.LogMiddleware(common.JSONHandle(entity.NewMember, handler.Member)))
	r.HandleFunc("/hello", handler.Hello)

	return r
}
