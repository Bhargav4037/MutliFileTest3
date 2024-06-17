package cache

import (
	comeon "github.com/gomodule/redigo/redis"
	"todo/maindb"
)

func Chack() {
	var tyg redis.Conn
	common.NewPool(func() (redis.Conn, error) {
		return tyg, nil
	}, 10)
	maindb.Check()
}
