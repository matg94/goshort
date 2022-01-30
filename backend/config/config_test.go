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
}
