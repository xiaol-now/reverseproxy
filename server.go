package reverseproxy

import (
	"github.com/gorilla/mux"
	"net/http"
	"reverseproxy/proxy"
	"strconv"
)

func StartProxyServer(config *Config) error {
	router := mux.NewRouter()
	for _, host := range config.ProxyHosts {
		p, err := proxy.NewHTTPProxy(host.Hosts, host.BalanceMode)
		if err != nil {
			return err
		}
		p.Heartbeat(config.HeartbeatInterval)
		router.Handle(host.Pattern, p)
	}
	if config.MaxAllowed > 0 {
		router.Use(MaxAllowedMiddleware)
	}
	return http.ListenAndServe(":"+strconv.Itoa(int(config.Port)), router)
}

func MaxAllowedMiddleware(handler http.Handler) http.Handler {
	return handler
}
