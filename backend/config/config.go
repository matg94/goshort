package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type RedisConfig struct {
	maxIdle   int    `yaml:"maxIdle"`
	maxActive int    `yaml:"maxActive"`
	port      int    `yaml:"port"`
	URL       string `yaml:"URL"`
}

type GoShortConfig struct {
	URLPrefix   string      `yaml:"url_prefix"`
	HashLength  int         `yaml:"hash_length"`
	RedisConfig RedisConfig `yaml:"redis"`
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

func GetRedisMaxActive() int {
	return goShortConfig.RedisConfig.maxActive
}

func GetRedisMaxIdle() int {
	return goShortConfig.RedisConfig.maxIdle
}

func GetRedisPort() int {
	return goShortConfig.RedisConfig.port
}

func GetRedisURL() string {
	return goShortConfig.RedisConfig.URL
}
