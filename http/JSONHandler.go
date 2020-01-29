package http

import (
	"encoding/json"
	"net/http"
)

//JSONHandler is a handler which creates a json
type JSONHandler func(w http.ResponseWriter, r *http.Request, p interface{})

//ServeHTTP just to be compliant with Handler interface.
func (j JSONHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

//JSONHandle handles request with json body.
func JSONHandle(constructor func() interface{}, h JSONHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := constructor()
		d := json.NewDecoder(r.Body)

		if err := d.Decode(p); err != nil {
			http.Error(w, "Malformed request.", 400)
			return
		}

		if h != nil {
			h(w, r, p)
			return
		}
	}
}
