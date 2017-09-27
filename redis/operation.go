/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package redis

import (
	"github.com/garyburd/redigo/redis"
)

func Set(r *redis.Pool, key string, value interface{}, time int64) error {
	conn := r.Get()
	err := conn.Err()
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	if time != 0 {
		_, err = conn.Do("EXPIRE", key, time)
		if err != nil {
			return err
		}
	}
	return nil
}

func Get(r *redis.Pool, key string) (interface{}, error) {
	conn := r.Get()
	err := conn.Err()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	data, err := conn.Do("GET", key)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Hmset(r *redis.Pool, key string, value interface{}) error {
	conn := r.Get()
	err := conn.Err()
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(value)...)
	if err != nil {
		return err
	}
	return nil
}

func Hgetall(r *redis.Pool, key string) (map[string]string, error) {
	conn := r.Get()
	err := conn.Err()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	data, err := conn.Do("HGETALL", key)
	if err != nil {
		return nil, err
	}
	return redis.StringMap(data, nil)
}

func ExpiredAfterGet(r *redis.Pool, key string) (interface{}, error) {
	conn := r.Get()
	err := conn.Err()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	data, err := conn.Do("GET", key)
	if err != nil {
		return nil, err
	}
	_, err = conn.Do("DEL", key)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ValueIsBool(v interface{}) (bool, error) {
	var err error
	return redis.Bool(v, err)
}

func ValueIsString(v interface{}) (string, error) {
	var err error
	return redis.String(v, err)
}
