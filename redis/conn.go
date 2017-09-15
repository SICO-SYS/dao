/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package redis

import (
	"github.com/garyburd/redigo/redis"
	"github.com/getsentry/raven-go"
)

func Pool(host, port, auth string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 2000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host+":"+port)
			if err != nil {
				raven.CaptureError(err, nil)
				return c, err
			}
			if auth != "" {
				_, err := c.Do("AUTH", auth)
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
