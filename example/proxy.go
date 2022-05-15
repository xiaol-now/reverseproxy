package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"reverseproxy/proxy"
	"strconv"
)

func main() {
	var hosts []string
	for i := 10000; i < 10010; i++ {
		go StartLocalHTTPServer(i)
		hosts = append(hosts, "http://127.0.0.1:"+strconv.Itoa(i))
	}
	httpProxy, err := proxy.NewHTTPProxy(hosts, "round-robin")
	if err != nil {
		log.Fatalln(err)
	}
	router := mux.NewRouter()
	router.Handle("/abc", httpProxy)
	log.Fatalln(http.ListenAndServe(":7777", router))
}

func StartLocalHTTPServer(port int) {
	router := mux.NewRouter()
	router.HandleFunc("/abc", func(w http.ResponseWriter, r *http.Request) {
		WriteString(w, "hello mux. port: "+strconv.Itoa(port))
		WriteString(w, "remoteAddr: "+r.RemoteAddr)
		WriteString(w, fmt.Sprintf("header: \n%#v", r.Header))
	})
	log.Fatalln(http.ListenAndServe(":"+strconv.Itoa(port), router))
}

func WriteString(w http.ResponseWriter, text string) {
	_, _ = w.Write([]byte(text + "\n"))
}
