package bd

import (
	"database/sql"
	"fmt"

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
