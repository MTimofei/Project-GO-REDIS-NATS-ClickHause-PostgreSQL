package interaction

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"git_p/test/pkg/errmy"
)

func TransactionPost(w http.ResponseWriter, db *sql.DB, campaignId int, payload Post) (rows *sql.Rows, err error) {

	tx, err := db.Begin()
	if err != nil {
		err = fmt.Errorf("open transaction %q", err)
		errmy.TransactionPost(w, tx)
		return nil, err
	}

	_, err = tx.Exec("LOCK TABLE items IN SHARE MODE")
	if err != nil {
		err = fmt.Errorf("blok %q", err)
		errmy.TransactionPost(w, tx)
		return nil, err
	}

	query := fmt.Sprintf("INSERT INTO items (campaign_id,name) VALUES (%d,'%s') RETURNING id;", campaignId, payload.Name)

	log.Println(query)

	result, err := tx.Query(query)
	if err != nil {
		err = fmt.Errorf("set %q", err)
		errmy.TransactionPost(w, tx)
		return nil, err

	}

	log.Println(result)

	err = tx.Commit()
	if err != nil {
		err = fmt.Errorf("close transaction %q", err)
		errmy.TransactionPost(w, tx)
		return nil, err
	}

	var id int
	result.Next()
	defer result.Close()
	result.Scan(&id)

	query = fmt.Sprintf("SELECT * FROM items WHERE id = %d;", id)
	rows, err = db.Query(query)

	log.Println(query)
	log.Println(result)

	if err != nil {
		err = fmt.Errorf("get %q", err)
		// errmy.TransactionPost(w, tx)
		return nil, err
	}

	// err = tx.Commit()
	// if err != nil {
	// 	err = fmt.Errorf("close transaction %q", err)
	// 	errmy.TransactionPost(w, tx)
	// 	return nil, err
	// }

	return rows, nil
}
