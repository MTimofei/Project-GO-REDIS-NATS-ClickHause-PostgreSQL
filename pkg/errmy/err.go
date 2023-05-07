package errmy

import (
	"bytes"
	"database/sql"
	"io"
	"log"
	"net/http"
	"os"
)

// откат транзакции и ответ клиетну статусом 500
func Transaction(w http.ResponseWriter, tx *sql.Tx) {
	tx.Rollback()
	w.WriteHeader(http.StatusInternalServerError)
}

// откат транзакции и ответ клиетну статусом 404 если нет записи
func TransactionNotFound(w http.ResponseWriter, tx *sql.Tx) {
	tx.Rollback()
	w.Header().Add("code", "3")
	w.Header().Add("massege", "errors.item.notFound")
	w.Header().Add("details", "{}")
	w.WriteHeader(http.StatusNotFound)

}

// кастомные логи с запаписью в терминал и в буфер для хранения
func Log(buf *bytes.Buffer) (Log *log.Logger) {
	var logM = log.New(io.MultiWriter(buf, os.Stdout), "API", log.LstdFlags)
	return logM
}
