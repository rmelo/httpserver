package http

import (
	"log"
	"net/http"
	gid "github.com/google/uuid"
)

//LogMiddleware logs everything
func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cid := gid.New().String()
		log.Printf("Executing request (%v)", cid)
		next.ServeHTTP(w, r)
		log.Printf("Finishing request (%v)", cid)
	})
}
