package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

var(
	pool *redis.Pool
	redisHost = "127.0.0.1:6379"
	redisPass = "testupload"
)

//创建Redis连接池
func newRedisPool() *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			//1.打开链接
			conn, err := redis.Dial("tcp", redisHost)
			if err != nil{
				fmt.Println(err)
				return nil, err
			}

			//2.访问认证
			if _, err = conn.Do("AUTH", redisPass); err != nil {
				conn.Close()
				return nil, err
			}
			return conn, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute{
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:      50,
		MaxActive:    30,
		IdleTimeout:  300 * time.Second,
		Wait:         false,
	}
}

func init() {
	pool = newRedisPool()
}

func RedisPool() *redis.Pool {
	return pool
}