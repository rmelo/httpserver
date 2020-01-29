package handler

import (
	"fmt"
	"log"
	"net/http"
)

//Hello is a handler of hello world
func Hello(w http.ResponseWriter, r *http.Request) {

	log.Print("Executing /hello handler.")
	fmt.Fprint(w, "Hello World!")

}
