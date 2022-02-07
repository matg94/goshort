package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/matg94/ezs3/ezs3lib"
	"gopkg.in/yaml.v2"
)

type RedisConfig struct {
	MaxIdle   int    `yaml:"MaxIdle"`
	MaxActive int    `yaml:"MaxActive"`
	Port      int    `yaml:"Port"`
	User      string `yaml:"Username"`
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

func DownloadConfig(origin string, target string) {
	s3Connection := ezs3lib.ConnectS3(
		os.Getenv("S3_BUCKET"),
		os.Getenv("S3_ENDPOINT"),
		os.Getenv("S3_REGION"),
	)
	err := s3Connection.DownloadFile(origin, target)
	if err != nil {
		fmt.Println(err)
	}
}

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

func GetRedisUser() string {
	return goShortConfig.RedisConf.User
}

func GetRedisPassword() string {
	return os.Getenv("REDIS_PASSWORD")
}
