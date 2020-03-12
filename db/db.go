package db

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

// Package contains all db interactions

var (
	Conn redis.Conn
	err  error
)

func init() {
	Conn, err = redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
}
