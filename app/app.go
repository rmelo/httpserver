package app

import (
	"github.com/rmelo/httpserver/http"
)

//App represents a app
type App struct {
	Server *http.Server
}

//NewApp returns an instance of App.
func NewApp(s *http.Server) *App {
	return &App{
		Server: s,
	}
}

//Run runs an App
func (a *App) Run() {
	a.Server.ListenAndServe()
}
