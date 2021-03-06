package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/rmelo/httpserver/entity"
)

//Time is a handler of timne
func Time(p interface{}) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Executing /time handler.")
		t := p.(*entity.Time)
		time.Sleep(time.Duration(t.Sleep) * time.Second)
		out, _ := json.Marshal(p)
		w.Header().Add("Content-Type", "application/json")
		w.Write(out)
	})
}
