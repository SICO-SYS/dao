/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package redis

import (
	"github.com/garyburd/redigo/redis"
)

func SetWithExpire(r *redis.Pool, key string, value interface{}, time int) error {
	conn := r.Get()
	err := conn.Err()
	defer conn.Close()
	conn.Do("SET", key, value)
	conn.Do("EXPIRE", key, time)
	return err
}

func SetWithUnexpire(r *redis.Pool, key string, value interface{}) error {
	conn := r.Get()
	err := conn.Err()
	defer conn.Close()
	conn.Do("SET", key, value)
	return err
}

func Hmset(r *redis.Pool, key string, value interface{}) error {
	conn := r.Get()
	err := conn.Err()
	defer conn.Close()
	_, operr := conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(value)...)
	if operr != nil {
		return operr
	}
	return err
}

func Hgetall(r *redis.Pool, key string) (map[string]string, error) {
	conn := r.Get()
	err := conn.Err()
	defer conn.Close()
	data, err := conn.Do("HGETALL", key)
	if err != nil {
		return nil, err
	}
	return redis.StringMap(data, nil)
}

func GetWithKey(r *redis.Pool, key string) (interface{}, error) {
	conn := r.Get()
	err := conn.Err()
	defer conn.Close()
	if err != nil {
		return "", err
	}
	data, err := conn.Do("GET", key)
	return data, err
}

func ExpiredAfterGetWithKey(r *redis.Pool, key string) (interface{}, error) {
	conn := r.Get()
	err := conn.Err()
	defer conn.Close()
	if err != nil {
		return "", err
	}
	data, err := conn.Do("GET", key)
	if err == nil {
		conn.Do("DEL", key)
	}
	return data, err
}

func ValueIsBool(v interface{}) (bool, error) {
	var err error
	return redis.Bool(v, err)
}

func ValueIsString(v interface{}) (string, error) {
	var err error
	return redis.String(v, err)
}
