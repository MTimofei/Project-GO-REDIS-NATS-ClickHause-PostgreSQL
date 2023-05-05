package errmy

import (
	"database/sql"
	"net/http"
)

func Transaction(w http.ResponseWriter, tx *sql.Tx) {
	tx.Rollback()
	w.WriteHeader(http.StatusInternalServerError)
}

func TransactionNotFound(w http.ResponseWriter, tx *sql.Tx) {
	tx.Rollback()
	w.Header().Add("code", "3")
	w.Header().Add("massege", "errors.item.notFound")
	w.Header().Add("details", "{}")
	w.WriteHeader(http.StatusNotFound)

}
