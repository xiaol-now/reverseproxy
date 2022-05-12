package main

import (
	"github.com/gorilla/mux"
	"log"
	"reverseproxy/proxy"
	"strconv"
)

func main() {
	var hosts []string
	for i := 10000; i < 10010; i++ {
		go StartLocalHTTPServer(i)
		hosts = append(hosts, strconv.Itoa(i))
	}
	httpProxy, err := proxy.NewHTTPProxy(hosts, "")
	if err != nil {
		log.Fatalln(err)
	}
	router := mux.NewRouter()
	router.Handle("/abc", httpProxy)
	
}