package repo

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/matg94/goshort/config"
)

type Redis struct {
	connPool *redis.Pool
}

var redisConn Redis

func InitRedis() {
	r := Redis{}
	r.connPool = &redis.Pool{
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
	redisConn = r
}

func (r *Redis) GetConnection() redis.Conn {
	return r.connPool.Get()
}

func (r *Redis) GET(key string) string {
	conn := r.GetConnection()
	val, err := conn.Do("GET", key)
	conn.Close()

	if err != nil || val == nil {
		return ""
	}
	return fmt.Sprintf("%s", val)
}

func (r *Redis) SET(key string, value string) error {
	fmt.Println("SET", key, "TO", value)
	conn := r.GetConnection()
	_, err := conn.Do("SET", key, value)
	conn.Close()

	if err != nil {
		return err
	}
	return nil
}
