package migration

import (
	"database/sql"
	"log"

	"git_p/test/insert/db/clickhause"
	"git_p/test/insert/db/postgres"
	"git_p/test/pkg/connect/bd"
)

func Migration(db *sql.DB, addr *string) {
	conCH, err := bd.ConectClickhause(addr)
	if err != nil {
		log.Println("ERROR Migration:", err)
		return
	}
	defer conCH.Close()

	resolt, err := postgres.TransactionMigartion(db)
	if err != nil {
		log.Println("ERROR Migration:", err)
		return
	}

	payloadres, err := postgres.CreateItems(resolt)
	if err != nil {
		log.Println("ERROR Migration:", err)
		return
	}

	err = clickhause.SetMigrationDates(conCH, payloadres)
	if err != nil {
		log.Println("ERROR Migration:", err)
		return
	}
}
