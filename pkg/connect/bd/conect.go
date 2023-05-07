package bd

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ClickHouse/ch-go"
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
	fmt.Println("postgres connect readi")
	return db, nil
}

func ConectClickhause(addr *string) (conn *ch.Client, err error) {
	conn, err = ch.Dial(context.Background(), ch.Options{Address: *addr, Database: "default"})
	if err != nil {
		return nil, err
	}
	err = conn.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return conn, nil
}
