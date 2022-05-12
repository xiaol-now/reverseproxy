package main

import (
	"reverseproxy"
	"testing"
)

func TestConfig(t *testing.T) {
	config, err := reverseproxy.ReadConfig("../config.yaml")
	//for _, location := range config.Location {
	//	t.Log(location.BalanceMode)
	//	t.Logf("%#v", location.ProxyHost)
	//}
	t.Logf("%+v\n%#v", config, err)
}
