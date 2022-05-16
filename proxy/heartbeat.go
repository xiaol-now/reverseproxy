package proxy

import (
	"fmt"
	"net"
	"time"
)

var ConnectionTimeout = 3 * time.Second

func (h *HTTPProxy) Heartbeat(interval uint) {
	for host := range h.alive {
		go h.heartbeat(host, interval)
	}
}

func (h *HTTPProxy) heartbeat(host string, interval uint) {
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	for range ticker.C {
		if IsBackendAlive(host) && !h.IsAlive(host) {
			h.SetAlive(host, true)
			h.bl.Add(host)
		} else if !IsBackendAlive(host) && h.IsAlive(host) {
			h.SetAlive(host, false)
			h.bl.Remove(host)
		}
	}
}

func (h *HTTPProxy) IsAlive(host string) bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.alive[host]
}

func (h *HTTPProxy) SetAlive(host string, alive bool) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.alive[host] = alive
}

func IsBackendAlive(host string) bool {
	addr, err := net.ResolveTCPAddr("tcp", host)
	if err != nil {
		return false
	}
	resolveAddr := fmt.Sprintf("%s:%d", addr.IP, addr.Port)
	conn, err := net.DialTimeout("tcp", resolveAddr, ConnectionTimeout)
	if err != nil {
		return false
	}
	_ = conn.Close()
	return true
}
