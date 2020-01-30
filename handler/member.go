package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rmelo/httpserver/entity"
)

//Member is a handler of about
func Member(p interface{}) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Executing /member handler.")
		m := p.(*entity.Member)
		out, _ := json.Marshal(m)
		w.Header().Add("Content-Type", "application/json")
		w.Write(out)
	})
}
