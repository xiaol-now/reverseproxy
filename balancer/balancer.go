package balancer

type Balancer interface {
	Balance(string) (string, error)
	Add(host string)
	Remove(host string)
}

func Factory(mode string, hosts []string) (Balancer, error) {
	switch mode {
	case "round-robin":
		return NewRoundRobin(hosts), nil
	default:
		return nil, ErrBalanceModeNotSupported
	}
}
