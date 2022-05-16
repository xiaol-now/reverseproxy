package reverseproxy

type Option func(*Config)

type Config struct {
	Port              uint16
	MaxAllowed        uint
	HeartbeatInterval uint
	ProxyHosts        []ProxyHost
}
type ProxyHost struct {
	Pattern     string
	Hosts       []string
	BalanceMode string
}

func WithPort(port uint16) Option {
	return func(config *Config) {
		config.Port = port
	}
}

func WithMaxAllowed(maxAllowed uint) Option {
	return func(config *Config) {
		config.MaxAllowed = maxAllowed
	}
}

func WithHeartbeatInterval(interval uint) Option {
	return func(config *Config) {
		config.HeartbeatInterval = interval
	}
}
func WithProxyHosts(hosts []ProxyHost) Option {
	return func(config *Config) {
		config.ProxyHosts = hosts
	}
}
func NewConfig(options ...Option) *Config {
	c := &Config{}
	for _, option := range options {
		option(c)
	}
	return c
}
