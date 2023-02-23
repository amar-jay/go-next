package redis

import (
	"github.com/amar-jay/go-api-boilerplate/utils/config"
	"github.com/google/uuid"
	"testing"
)

//var redigoConn = redigomock.NewConn()

const (
	mock_addr = "localhost:6379"
)

var (
	c = config.RedisConfig{
		Address: mock_addr,
	}
)

func Test_Redis(t *testing.T) {
	var conn = Init(mock_addr)
	key := uuid.New().String()
	val := uuid.New()
	t.Run("dial redis", func(t *testing.T) {
		if err := conn.Dial(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("set key value", func(t *testing.T) {
		if err := conn.Set(key, val); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("get value", func(t *testing.T) {
		var res uuid.UUID
		err := conn.Get(key, &res)
		if err != nil {
			t.Fatal(err)
		}
		//		t.Logf("set val: %s\ngot val: %s", res, val);
		if res.String() != val.String() {
			t.Fatal("value got dies not match")
		}
	})
}
