package http

import (
	"encoding/json"
	"net/http"
)

//DecodeHandlerFunc is a func
type DecodeHandlerFunc func(p interface{}) http.Handler

//ConstructorFunc representing a constructor of any type.
type ConstructorFunc func() interface{}

//DecoderHandler decodes payload caming from requests.
type decoderHandler struct {
	handler     DecodeHandlerFunc
	constructor ConstructorFunc
}

//ServerHTTP to satisfy contract
func (d *decoderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s := d.constructor()
	if err := json.NewDecoder(r.Body).Decode(s); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	d.handler(s).ServeHTTP(w, r)
}

//WithDecode return a new instance of decoderHandler
func WithDecode(c ConstructorFunc, h DecodeHandlerFunc) http.Handler {
	return &decoderHandler{
		handler:     h,
		constructor: c,
	}
}
