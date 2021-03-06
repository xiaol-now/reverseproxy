package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/xiaol-now/reverseproxy"
	"log"
	"net/http"
	"strconv"
)

func main() {
	var hosts []string
	for i := 10000; i < 10010; i++ {
		go StartLocalHTTPServer(i)
		hosts = append(hosts, "http://127.0.0.1:"+strconv.Itoa(i))
	}
	config := reverseproxy.NewConfig(
		reverseproxy.WithPort(9999),
		reverseproxy.WithHeartbeatInterval(5),
		reverseproxy.WithMaxAllowed(100),
		reverseproxy.WithProxyHosts([]reverseproxy.ProxyHost{
			{Pattern: "/abc", Hosts: hosts, BalanceMode: "round-robin"},
		}),
	)
	log.Fatalln(reverseproxy.StartProxyServer(config))
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
