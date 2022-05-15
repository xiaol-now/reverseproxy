package reverseproxy

import (
	"github.com/gorilla/mux"
	"reverseproxy/proxy"
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
	return nil
}
