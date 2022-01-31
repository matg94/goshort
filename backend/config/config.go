package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type RedisConfig struct {
	MaxIdle   int    `yaml:"MaxIdle"`
	MaxActive int    `yaml:"MaxActive"`
	Port      int    `yaml:"Port"`
	URL       string `yaml:"URL"`
}

type GoShortConfig struct {
	URLPrefix  string      `yaml:"url_prefix"`
	HashLength int         `yaml:"hash_length"`
	RedisConf  RedisConfig `yaml:"redis"`
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
	return goShortConfig.RedisConf.MaxActive
}

func GetRedisMaxIdle() int {
	return goShortConfig.RedisConf.MaxIdle
}

func GetRedisPort() int {
	return goShortConfig.RedisConf.Port
}

func GetRedisURL() string {
	return goShortConfig.RedisConf.URL
}
