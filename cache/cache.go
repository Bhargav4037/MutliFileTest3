package cache

import (
	comeon "github.com/gomodule/redigo/redis"
	"todo/maindb"
)

func Chack() {
	var tyg comeon.Conn
	comeon.NewPool(func() (comeon.Conn, error) {
		return tyg, nil
	}, 10)
	maindb.Check()
}
