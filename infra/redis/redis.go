package redis

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/amar-jay/go-api-boilerplate/utils/config"
)


// connect to redis server
func Dial() (redis.Conn, error) {
  var config = config.GetRedisConfig()
  conn, err := redis.Dial("tcp", config.Address)

  if err != nil {
    return nil, err
  }

  return conn, nil
}

type Value any

// set key value pair
func Set(key string, val Value) error {
  conn, err := Dial()
  if err != nil {
    msg := fmt.Sprintf("error in redis connection: %v", err)
    return errors.New(msg)
  }

  // json serialization
  b, err := json.Marshal(val)

  if err != nil {
    return err;
  }

  _, err = conn.Do("SET", key, string(b))
  return err

}

func Get(key string) (any, error) {
  conn, err := Dial()
  if err != nil {
    msg := fmt.Sprintf("error in redis connection: %v", err)
    return nil, errors.New(msg)
  }
  if err != nil {
    return nil, err;
  }

  res, err := conn.Do("GET", key)
  return res, err

}
