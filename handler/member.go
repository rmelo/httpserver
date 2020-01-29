package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rmelo/httpserver/entity"
)

//Member is a handler of about
func Member(w http.ResponseWriter, r *http.Request, p interface{}) {

	log.Print("Executing /member handler.")

	m := p.(*entity.Member)

	out, _ := json.Marshal(m)

	w.Header().Add("Content-Type", "application/json")
	w.Write(out)
}
