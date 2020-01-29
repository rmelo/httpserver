package http

import "net/http"

//ServeOpts represents the options of server.
type ServeOpts struct {
	Addr     string
	RootPath string
}

//Server returns a new server
type Server struct {
	Router http.Handler
	Mux    *http.ServeMux
	Opts   ServeOpts
}

//NewServer returns a new server
func NewServer(opts ServeOpts, r http.Handler) *Server {
	mux := http.NewServeMux()
	mux.Handle(opts.RootPath, r)
	return &Server{
		Router: r,
		Mux:    mux,
		Opts:   opts,
	}
}

//ListenAndServe delegates to http.ListenAndServe func.
func (s *Server) ListenAndServe() error {
	return http.ListenAndServe(s.Opts.Addr, s.Router)
}
