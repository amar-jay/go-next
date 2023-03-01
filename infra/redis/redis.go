package redis

import (
	"encoding/json"
	"errors"

	"github.com/amar-jay/go-api-boilerplate/pkg/config"
	"github.com/gomodule/redigo/redis"
)

type redisConn struct {
	addr string
	conn *redis.Conn
}

func Init(args ...string) *redisConn {
	var addr string
	var config = config.GetRedisConfig()
	switch {
	case len(args) > 1:
		panic("only address should be provided")
	case len(args) == 1:
		addr = args[0]
	default:
		addr = config.Address
	}

	return &redisConn{
		addr: addr,
		conn: nil,
	}
}

// connect to redis server
func (r *redisConn) Dial() error {
	conn, err := redis.Dial("tcp", r.addr)

	if err != nil {
		return err
	}

	r.conn = &conn
	return nil
}

type Value any

// set key value pair
func (r *redisConn) Set(key string, val Value) error {
	conn := *r.conn

	// json serialization
	b, err := json.Marshal(val)

	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, string(b))
	return err
}

// get valu e and store value in pointer
func (r *redisConn) Get(key string, val any) error {
	if r.conn == nil {
		msg := errors.New("error in redis connection: Ought to dial first")
		return msg
	}

	conn := *r.conn

	res, err := conn.Do("GET", key)
	if err != nil {
		return err
	}
	bytes, ok := res.([]byte)
	if !ok {
		return errors.New("unable to convert response to bytes")
	}
	err = json.Unmarshal(bytes, &val)
	return err

}
