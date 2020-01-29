package http

import "net/http"

//Server returns a new server
type Server struct {
	Router http.Handler
}
