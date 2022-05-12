package proxy

import "sync"

type Heartbeat struct {
	alive map[string]bool
	sync.RWMutex
}