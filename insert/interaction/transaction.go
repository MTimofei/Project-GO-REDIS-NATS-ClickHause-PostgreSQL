package interaction

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"git_p/test/pkg/errmy"
)

// func NewCampaign(db *sql.DB, name string) (rows *sql.Rows, err error) {
// 	query := fmt.Sprintf("INSERT INTO campaingns (name) VALUES ('%s') RETURNING id;", name)
// 	rows, err = db.Query(query)
// 	return
// }

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

	// log.Println(query)

	result, err := tx.Query(query)
	if err != nil {
		err = fmt.Errorf("set %q", err)
		errmy.TransactionPost(w, tx)
		return nil, err

	}

	// log.Println(result)

	var id int
	result.Next()
	defer result.Close()
	result.Scan(&id)

	err = tx.Commit()
	if err != nil {
		err = fmt.Errorf("close transaction %q", err)
		errmy.TransactionPost(w, tx)
		return nil, err
	}

	query = fmt.Sprintf("SELECT * FROM items WHERE id = %d;", id)
	rows, err = db.Query(query)

	// log.Println(query)
	// log.Println(result)

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

func TransactionPatch(w http.ResponseWriter, db *sql.DB, campaignId int, id int, payload Update) (rows *sql.Rows, err error) {
	tx, err := db.Begin()
	if err != nil {
		err = fmt.Errorf("open transaction %q", err)
		errmy.TransactionPost(w, tx)
		return nil, err
	}

	var exists bool
	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM items WHERE id=%d AND campaign_id=%d);", id, campaignId)
	log.Println(query)

	err = db.QueryRow(query).Scan(&exists)
	if err != nil {
		err = fmt.Errorf("chek record: %q", err)
		errmy.TransactionPost(w, tx)
		return nil, err
	}
	if !exists {
		err = fmt.Errorf("chek record:The record does not exist")
		errmy.TransactionPost(w, tx)
		return nil, err
	}

	query = fmt.Sprintf("SELECT * FROM items WHERE id=%d AND campaign_id=%d FOR UPDATE;", id, campaignId)
	log.Println(query)

	_, err = tx.Exec(query)
	if err != nil {
		err = fmt.Errorf("blok %q", err)
		errmy.TransactionPost(w, tx)
		return nil, err
	}

	query = fmt.Sprintf("UPDATE items SET name='%s', description='%q' WHERE id=%d AND campaign_id=%d;", payload.Name, payload.Description, id, campaignId)
	log.Println(query)

	_, err = tx.Exec(query)
	if err != nil {
		err = fmt.Errorf("set %q", err)
		errmy.TransactionPost(w, tx)
		return nil, err

	}

	err = tx.Commit()
	if err != nil {
		err = fmt.Errorf("close transaction %q", err)
		errmy.TransactionPost(w, tx)
		return nil, err
	}

	query = fmt.Sprintf("SELECT * FROM items WHERE id = %d;", id)
	log.Println(query)

	rows, err = db.Query(query)

	if err != nil {
		err = fmt.Errorf("get %q", err)
		return nil, err
	}

	return rows, nil
}

func TransactionDelete(w http.ResponseWriter, db *sql.DB, campaignId int, id int) (rows *sql.Rows, err error) {
	tx, err := db.Begin()
	if err != nil {
		err = fmt.Errorf("open transaction %q", err)
		errmy.TransactionPost(w, tx)
		return nil, err
	}

	var exists bool
	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM items WHERE id=%d AND campaign_id=%d);", id, campaignId)
	log.Println(query)

	err = db.QueryRow(query).Scan(&exists)
	if err != nil {
		err = fmt.Errorf("chek record: %q", err)
		errmy.TransactionPost(w, tx)
		return nil, err
	}
	if !exists {
		err = fmt.Errorf("chek record:The record does not exist")
		errmy.TransactionPost(w, tx)
		return nil, err
	}

	query = fmt.Sprintf("SELECT * FROM items WHERE id=%d AND campaign_id=%d FOR UPDATE;", id, campaignId)
	log.Println(query)

	_, err = tx.Exec(query)
	if err != nil {
		err = fmt.Errorf("blok %q", err)
		errmy.TransactionPost(w, tx)
		return nil, err
	}

	query = fmt.Sprintf("UPDATE items SET  removed=true WHERE id=%d AND campaign_id=%d;", id, campaignId)
	log.Println(query)

	_, err = tx.Exec(query)
	if err != nil {
		err = fmt.Errorf("set %q", err)
		errmy.TransactionPost(w, tx)
		return nil, err

	}

	err = tx.Commit()
	if err != nil {
		err = fmt.Errorf("close transaction %q", err)
		errmy.TransactionPost(w, tx)
		return nil, err
	}

	query = fmt.Sprintf("SELECT id,campaign_Id,removed FROM items WHERE id = %d;", id)
	log.Println(query)

	rows, err = db.Query(query)

	if err != nil {
		err = fmt.Errorf("get %q", err)
		return nil, err
	}

	return rows, nil
}
