package proxy

import (
	"net"
	"net/http"
	"strings"
)

var (
	XRealIP       = http.CanonicalHeaderKey("X-Real-IP")
	XForwardedFor = http.CanonicalHeaderKey("X-Forwarded-For")
)

func GetClientIP(r *http.Request) string {
	clientIP, _, _ := net.SplitHostPort(r.RemoteAddr)
	if len(r.Header.Get(XForwardedFor)) != 0 {
		xff := r.Header.Get(XForwardedFor)
		s := strings.Index(xff, ", ")
		if s == -1 {
			s = len(r.Header.Get(XForwardedFor))
		}
		clientIP = xff[:s]
	} else if len(r.Header.Get(XRealIP)) != 0 {
		clientIP = r.Header.Get(XRealIP)
	}

	return clientIP
}
