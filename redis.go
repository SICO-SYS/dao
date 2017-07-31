/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package dao

import (
	"github.com/garyburd/redigo/redis"
	"github.com/getsentry/raven-go"
)

var (
	RedisPool *redis.Pool
)

func RedisSetShort(k string, v interface{}, t int16) error {
	conn := RedisPool.Get()
	err = conn.Err()
	defer conn.Close()
	conn.Do("SET", k, v)
	conn.Do("EXPIRE", k, t)
	return err
}

func RedisSetLong(k string, v interface{}) error {
	conn := RedisPool.Get()
	err = conn.Err()
	defer conn.Close()
	conn.Do("SET", k, v)
	return err
}

func RedisGetValue(k string) (interface{}, error, error) {
	conn := RedisPool.Get()
	err = conn.Err()
	defer conn.Close()
	data, err2 := conn.Do("GET", k)
	return data, err, err2
}

func RedisBool(v interface{}) (bool, error) {
	return redis.Bool(v, err)
}

func RedisString(v interface{}) (string, error) {
	return redis.String(v, err)
}

func init() {
	RedisPool = &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.Redis.Default.Host+":"+config.Redis.Default.Port)
			if err != nil {
				raven.CaptureError(err, nil)
			}
			return c, err
		},
	}

	err = RedisPool.Get().Close()
	if err != nil {
		raven.CaptureError(err, nil)
	}
}
