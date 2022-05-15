package balancer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRoundRobin(t *testing.T) {
	tests := []struct {
		name  string
		hosts []string
	}{
		{name: "round_robin01", hosts: []string{"127.0.0.1", "127.0.0.2", "127.0.0.3"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			counts := make(map[string]int)
			b := NewRoundRobin(tt.hosts)
			for i := 0; i < 99; i++ {
				host, err := b.Balance("")
				if err != nil {
					assert.Error(t, err)
				}
				counts[host]++
			}
			for _, i := range counts {
				assert.Equal(t, i, 33)
			}
		})
	}
}
