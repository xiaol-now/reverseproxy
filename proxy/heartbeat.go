package proxy

import "time"

var ConnectionTimeout = 3 * time.Second

func (h *HTTPProxy) Heartbeat(interval uint) {
	for host := range h.alive {
		go h.heartbeat(host, interval)
	}
}

func (h *HTTPProxy) heartbeat(host string, interval uint) {

}
