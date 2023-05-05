package bd

import (
	"database/sql"
	"log"

	"github.com/gomodule/redigo/redis"
	_ "github.com/lib/pq"
)

func ConectPostgras(connStr *string) (db *sql.DB, err error) {
	db, err = sql.Open("postgres", *connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Println("postgres connect readi")
	return db, nil
}

func ConectRedis(redisAddr *string) (conn redis.Conn, err error) {
	conn, err = redis.Dial("tcp", *redisAddr)
	if err != nil {
		return nil, err
	}
	log.Println("redis connect readi")
	return conn, err
}
