/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package redis

import (
	"github.com/garyburd/redigo/redis"
)

func NewPool() *redis.Pool { return &redis.Pool{} }

func InitPool(host, port, auth string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 1000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host+":"+port)
			if err != nil {
				return c, err
			}
			if auth != "" {
				_, err := c.Do("AUTH", auth)
				if err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, nil
		},
	}
}
