package balancer

import "sync"

type RoundRobin struct {
	hosts []string
	i     int
	sync.RWMutex
}

func NewRoundRobin(hosts []string) *RoundRobin {
	return &RoundRobin{hosts: hosts}
}

func (r *RoundRobin) Balance(_ string) (string, error) {
	r.RLock()
	defer r.RUnlock()
	if len(r.hosts) == 0 {
		return "", ErrHostEmpty
	}
	host := r.hosts[r.i%len(r.hosts)]
	r.i++
	return host, nil
}

func (r *RoundRobin) Add(host string) {
	r.Lock()
	defer r.Unlock()
	r.hosts = append(r.hosts, host)
}

func (r *RoundRobin) Remove(host string) {
	r.Lock()
	defer r.Unlock()
	for i, h := range r.hosts {
		if h == host {
			r.hosts = append(r.hosts[:i], r.hosts[i+1:]...)
			return
		}
	}
}
