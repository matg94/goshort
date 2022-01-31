package config

import (
	"fmt"
	"testing"
)

func TestConfigReadCorrectly(t *testing.T) {
	ReadConfig("test.yaml")

	if GetHashLength() != 4 {
		failLog := fmt.Sprintf("Config read failed, expected HashLength %d, but got %d", 4, GetHashLength())
		t.Error(failLog)
	}

	if GetURLPrefix() != "http://test/" {
		failLog := fmt.Sprintf("Config read failed, expected URLPrefix %s, but got %s", "http://test/", GetURLPrefix())
		t.Error(failLog)
	}

	if GetRedisPort() != 1 {
		failLog := fmt.Sprintf("Config read failed, expected RedisPort %d, but got %d", 1, GetRedisPort())
		t.Error(failLog)
	}

	if GetRedisURL() != "1.1.1.1" {
		failLog := fmt.Sprintf("Config read failed, expected RedisURL %s, but got %s", "1.1.1.1", GetRedisURL())
		t.Error(failLog)
	}
}
