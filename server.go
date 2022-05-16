package reverseproxy

import (
	"github.com/gorilla/mux"
	"github.com/xiaol-now/reverseproxy/proxy"
	"net/http"
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
		router.Use(MaxAllowedMiddleware(config.MaxAllowed))
	}
	return http.ListenAndServe(":"+strconv.Itoa(int(config.Port)), router)
}

func MaxAllowedMiddleware(n uint) mux.MiddlewareFunc {
	max := make(chan struct{}, n)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() { <-max }()
			max <- struct{}{}
			next.ServeHTTP(w, r)
		})
	}
}
