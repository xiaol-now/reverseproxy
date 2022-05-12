package reverseproxy

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Port                uint16         `yaml:"port"`
	SSLCertificateKey   string      `yaml:"ssl_certificate_key"`
	SSLCertificate      string      `yaml:"ssl_certificate"`
	HeartbeatInterval uint        `yaml:"heartbeat_interval"`
	MaxAllowed          uint        `yaml:"max_allowed"`
	ProxyHosts            []ProxyHost `yaml:"proxy_hosts"`
}

type ProxyHost struct {
	Pattern     string   `yaml:"pattern"`
	Host   []string `yaml:"host"`
	BalanceMode string   `yaml:"balance_mode"`
}

func ReadConfig(filename string) (*Config, error) {
	in, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var config Config
	if err = yaml.Unmarshal(in, &config);err != nil {
		return nil, err
	}
	if err = config.Validation(); err != nil {
		return nil, err
	}
	return &config, nil
}

func (c *Config) Validation() error {
	return nil
}
