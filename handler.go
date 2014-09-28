// Http/net and Gorilla/mux Handlers for JSON API usage.

package httpUtil

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// set to true to send JSON in pretty format
var Pretty bool
var logger *log.Logger

// if set to true, logs are put into file. False by default
//var LogToFile bool = false

// logfile name. by default its http.log
//var Logfile string = "http.log"

// Type that contains all needed error data
type HandlerError struct {
	Error   error
	Message string
	Code    int
}

// run this function to log to file instead of stderr. just put name of the file.
// path example: "./logs"
// file example: "http.log"
func LogToFile(path string, file string) (*os.File, error) {
	filePath := fmt.Sprintf("%s/%s", path, file)
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	logger = log.New(f, "http-log: ", log.Lshortfile)
	return f, nil
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
		logger.Printf("ERROR: %v\n", err.Error)
		http.Error(w, fmt.Sprintf("error: %v", err.Error), err.Code)
		return
	}

	// check if response was empty
	if res == nil {
		logger.Printf("ERROR: response from server is nil\n")
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
	logger.Printf("%s %s %s %d", r.RemoteAddr, r.Method, r.URL, 200) // log request
	return
}

func init() {
	logger = log.New(os.Stderr, "", log.LstdFlags)
}
