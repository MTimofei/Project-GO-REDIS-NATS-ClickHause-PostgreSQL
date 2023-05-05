package errmy

import (
	"database/sql"
	"net/http"
)

func TransactionPost(w http.ResponseWriter, tx *sql.Tx) {
	tx.Rollback()
	w.WriteHeader(http.StatusInternalServerError)
}
