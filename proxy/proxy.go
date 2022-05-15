package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"reverseproxy/balancer"
	"sync"
)

type HTTPProxy struct {
	hostMap map[string]*httputil.ReverseProxy
	bl      balancer.Balancer

	// Heartbeat check
	alive map[string]bool
	mu    sync.RWMutex
}

func NewHTTPProxy(proxyPass []string, balanceMode string) (*HTTPProxy, error) {
	hosts := make([]string, len(proxyPass))
	hostMap := make(map[string]*httputil.ReverseProxy, len(proxyPass))
	alive := make(map[string]bool, len(proxyPass))
	for i, host := range proxyPass {
		u, err := url.Parse(host)
		if err != nil {
			return nil, err
		}
		hostMap[u.Host] = NewReverseProxy(u)
		hosts[i] = u.Host
		alive[host] = true
	}
	bl, err := balancer.Factory(balanceMode, hosts)
	if err != nil {
		return nil, err
	}
	return &HTTPProxy{
		alive:   alive,
		hostMap: hostMap,
		bl:      bl,
	}, nil
}

func NewReverseProxy(u *url.URL) *httputil.ReverseProxy {
	return httputil.NewSingleHostReverseProxy(u)
}

func (h *HTTPProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("proxy backend panic: %s", err)
			w.WriteHeader(http.StatusBadGateway)
			_, _ = w.Write([]byte(err.(error).Error()))
		}
	}()
	host, err := h.bl.Balance(GetClientIP(r))
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		_, _ = w.Write([]byte("balance error: " + err.Error()))
	}
	h.hostMap[host].ServeHTTP(w, r)
}
