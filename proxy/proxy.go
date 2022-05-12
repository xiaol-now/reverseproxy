package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"reverseproxy/balancer"
)

type HTTPProxy struct {
	Heartbeat

	hostMap map[string]*httputil.ReverseProxy
	bl      balancer.Balancer
}

func NewHTTPProxy(proxyPass []string, balanceMode string) (*HTTPProxy, error) {
	hosts := make([]string, len(proxyPass))
	hostMap := make(map[string]*httputil.ReverseProxy, len(proxyPass))
	alive := make(map[string]bool, len(proxyPass))
	for _, host := range proxyPass {
		u, err := url.Parse(host)
		if err != nil {
			return nil, err
		}
		hostMap[u.Host] = NewReverseProxy(u)
		hosts = append(hosts, u.Host)
		alive[host] = true
	}
	return &HTTPProxy{
		Heartbeat: Heartbeat{alive: alive},
		hostMap:   hostMap,
		bl:        balancer.Factory(balanceMode, hosts),
	}, nil
}

func NewReverseProxy(u *url.URL) *httputil.ReverseProxy {
	return httputil.NewSingleHostReverseProxy(u)
}

func (h *HTTPProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 处理请求失败的情况
	host, err := h.bl.Balance(GetClientHost(r))
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		_, _ = w.Write([]byte("balance error: " + err.Error()))
	}
	h.hostMap[host].ServeHTTP(w, r)
}
