// Http/net and Gorilla/mux Handlers for JSON API usage.

package httpUtil

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var (
	Pretty bool
)

// Type that contains all needed error data
type HandlerError struct {
	Error   error
	Message string
	Code    int
}

// custom handler type.
// use with http/net or gorilla/mux Handle function, not HandleFunc
type Handler func(w http.ResponseWriter, r *http.Request) (interface{}, *HandlerError)

// method that satisfies Handle interface and let http call it.
//It should set Header and write to ResponseWriter, then return.
func (fn Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//call the Handler
	res, err := fn(w, r)

	// check if there was error
	if err != nil {
		log.Printf("ERROR: %v\n", err.Error)
		http.Error(w, fmt.Sprintf("error: %v", err.Error), err.Code)
		return
	}

	// check if response was empty
	if res == nil {
		log.Printf("ERROR: response from server is nil\n")
		http.Error(w, fmt.Sprintf("Internal server error"), http.StatusInternalServerError)
		return
	}

	// encode to json
	var j []byte
	var e error
	if Pretty {
		j, e = json.MarshalIndent(res, "", "  ")
	} else {
		j, e = json.Marshal(res)
	}

	if e != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}
	// set Headers
	w.Header().Set("Content-Type", "application/json")
	// write to writer
	w.Write(j)
	log.Printf("%s %s %s %d", r.RemoteAddr, r.Method, r.URL, 200) // log request
	return
}
