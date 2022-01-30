package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type GoShortConfig struct {
	URLPrefix  string `yaml:"url_prefix"`
	HashLength int    `yaml:"hash_length"`
}

func (c *GoShortConfig) Parse(data []byte) error {
	return yaml.Unmarshal(data, c)
}

var goShortConfig GoShortConfig

func ReadConfig(file string) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("Could not read Config ", err)
	}
	if err := goShortConfig.Parse(data); err != nil {
		log.Fatal("Could not parse Config ", err)
	}
}

func GetURLPrefix() string {
	return goShortConfig.URLPrefix
}

func GetHashLength() int {
	return goShortConfig.HashLength
}
