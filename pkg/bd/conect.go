package bd

import (
	"database/sql"

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

	return db, nil
}
