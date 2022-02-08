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

var redisConn Redis

func InitRedis() {
	r := Redis{}
	var pool = &redis.Pool{
		MaxIdle:   config.GetRedisMaxIdle(),
		MaxActive: config.GetRedisMaxActive(),
		Dial: func() (redis.Conn, error) {
			var c redis.Conn
			var err error
			if config.GetRedisPassword() == "" {
				c, err = redis.Dial(
					"tcp",
					fmt.Sprintf(
						"%s:%d",
						config.GetRedisURL(),
						config.GetRedisPort(),
					),
				)
			} else {
				c, err = redis.Dial(
					"tcp",
					fmt.Sprintf(
						"%s:%d",
						config.GetRedisURL(),
						config.GetRedisPort(),
					),
					redis.DialPassword(config.GetRedisPassword()),
					redis.DialUseTLS(true),
				)
			}
			if err != nil {
				log.Fatal(err.Error())
			}
			return c, err
		},
	}
	r.client = pool.Get()
	redisConn = r
}

func (r *Redis) GET(key string) string {
	fmt.Println("GET ", key)
	val, err := r.client.Do("GET", key)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s", val)
}

func (r *Redis) SET(key string, value string) {
	fmt.Println("SET ", key)
	r.client.Do("SET", key, value)
}
