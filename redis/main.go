/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package redis

import (
	"github.com/garyburd/redigo/redis"
	"github.com/getsentry/raven-go"

	"github.com/SiCo-Ops/cfg"
)

var (
	config = cfg.Config

	PublicPool *redis.Pool
)

func RedisSetWithExpire(r *redis.Pool, key string, value interface{}, time int16) error {
	conn := r.Get()
	err := conn.Err()
	defer conn.Close()
	conn.Do("SET", key, value)
	conn.Do("EXPIRE", key, time)
	return err
}

func RedisSetWithUnexpire(r *redis.Pool, key string, value interface{}) error {
	conn := r.Get()
	err := conn.Err()
	defer conn.Close()
	conn.Do("SET", key, value)
	return err
}

func RedisGetWithKey(r *redis.Pool, key string) (interface{}, error, error) {
	conn := r.Get()
	err := conn.Err()
	defer conn.Close()
	data, err2 := conn.Do("GET", key)
	return data, err, err2
}

func RedisValueIsBool(v interface{}) (bool, error) {
	var err error
	return redis.Bool(v, err)
}

func RedisValueIsString(v interface{}) (string, error) {
	var err error
	return redis.String(v, err)
}

func init() {
	PublicPool = &redis.Pool{
		MaxIdle:   80,
		MaxActive: 2000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.Redis.Public.Host+":"+config.Redis.Public.Port)
			if err != nil {
				raven.CaptureError(err, nil)
				return c, err
			}
			if config.Redis.Public.Auth != "" {
				_, err := c.Do("AUTH", config.Redis.Public.Auth)
				if err != nil {
					raven.CaptureError(err, nil)
					c.Close()
					return nil, err
				}
			}
			return c, err

		},
	}
}
