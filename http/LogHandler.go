package http

import (
	gid "github.com/google/uuid"
	"log"
	"net/http"
)

//WithLogging logs everything
func WithLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cid := gid.New().String()
		log.Printf("Executing request (%v)", cid)
		next.ServeHTTP(w, r)
		log.Printf("Finishing request (%v)", cid)
	})
}
