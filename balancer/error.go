package balancer

import "errors"

var (
	ErrHostEmpty               = errors.New("host empty")
	ErrBalanceModeNotSupported = errors.New("balance mode not supported")
)
