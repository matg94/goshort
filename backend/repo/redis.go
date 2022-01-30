package repo

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/matg94/goshort/config"
)

type Redis struct {
	client redis.Conn
}

func (r *Redis) Init() {
	var pool = &redis.Pool{
		MaxIdle:   config.GetRedisMaxIdle(),
		MaxActive: config.GetRedisMaxActive(),
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", config.GetRedisURL(), config.GetRedisPort()))
			if err != nil {
				log.Fatal(err.Error())
			}
			return c, err
		},
	}
	r.client = pool.Get()
}

func (r *Redis) GET(key string) interface{} {
	val, err := r.client.Do("GET", key)
	if err != nil {
		return ""
	}
	return val
}

func (r *Redis) SET(key string, value string) {
	r.client.Do("SET", key, value)
}
