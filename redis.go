package main

import "github.com/garyburd/redigo/redis"

var pool *redis.Pool
var redisServer = "127.0.0.1:6379"

func createRedisPool() {
	pool = &redis.Pool{
		MaxIdle: 3,
		Dial: func() (c redis.Conn, err error) {
			c, err = redis.Dial("tcp", redisServer)
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}
}

func getRedirectIPs(c redis.Conn, host string) ([]string, error) {
	return redis.Strings(c.Do("HKEYS", host))
}
