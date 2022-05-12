package balancer

type Balancer interface {
	Balance(string) (string, error)
}

func Factory(mode string, hosts []string) Balancer {
	return nil
}

type TODO struct{}

func (TODO) Balance(s string) (string, error) {
	return s, nil
}
