package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func StartLocalHTTPServer(port int) {
	router := mux.NewRouter()
	router.HandleFunc("/abc", func(w http.ResponseWriter, r *http.Request) {
		WriteString(w, "hello mux. port: " + strconv.Itoa(port))
		WriteString(w, "remoteAddr: "+ r.RemoteAddr)
		WriteString(w, fmt.Sprintf("header: \n%#v", r.Header))
	})
	log.Fatalln(http.ListenAndServe(":"+strconv.Itoa(port), router))
}

func WriteString(w http.ResponseWriter, text string) {
	_, _ = w.Write([]byte(text + "\n"))
}