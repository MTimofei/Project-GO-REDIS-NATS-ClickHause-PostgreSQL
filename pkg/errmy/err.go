package errmy

import (
	"bytes"
	"database/sql"
	"io"
	"log"
	"net/http"
	"os"
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

func Log(buf *bytes.Buffer) (Log *log.Logger) {
	// var logmyS = log.New(os.Stdout|buf, "API Log", log.LstdFlags)
	// var logmyB = log.New(buf, "API Log", log.LstdFlags)

	var logM = log.New(io.MultiWriter(buf, os.Stdout), "", log.LstdFlags)
	return logM
}
