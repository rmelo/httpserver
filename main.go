package main

import (
	"github.com/rmelo/httpserver/app"
	"github.com/rmelo/httpserver/http"
)

func main() {

	opts := http.ServeOpts{
		Addr:     ":8080",
		RootPath: "/",
	}

	s := http.NewServer(opts, routes())
	app.NewApp(s).Run()
}
