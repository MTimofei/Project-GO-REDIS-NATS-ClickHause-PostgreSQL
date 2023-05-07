package postgres

import (
	"database/sql"
	"fmt"
	"net/http"

	"git_p/test/pkg/errmy"
)

func TransactionPost(w http.ResponseWriter, db *sql.DB, campaignId int, payload Post) (rows *sql.Rows, err error) {

	tx, err := db.Begin()
	if err != nil {
		err = fmt.Errorf("open transaction %q", err)
		errmy.Transaction(w, tx)
		return nil, err
	}

	_, err = tx.Exec("LOCK TABLE items IN SHARE MODE")
	if err != nil {
		err = fmt.Errorf("blok %q", err)
		errmy.Transaction(w, tx)
		return nil, err
	}

	query := fmt.Sprintf("INSERT INTO items (campaign_id,name) VALUES (%d,'%s') RETURNING id;", campaignId, payload.Name)

	result, err := tx.Query(query)
	if err != nil {
		err = fmt.Errorf("set %q", err)
		errmy.Transaction(w, tx)
		return nil, err

	}

	var id int
	result.Next()
	defer result.Close()
	result.Scan(&id)

	err = tx.Commit()
	if err != nil {
		err = fmt.Errorf("close transaction %q", err)
		errmy.Transaction(w, tx)
		return nil, err
	}

	query = fmt.Sprintf("SELECT * FROM items WHERE id = %d;", id)
	rows, err = db.Query(query)

	if err != nil {
		err = fmt.Errorf("get %q", err)
		return nil, err
	}

	return rows, nil
}

func TransactionPatch(w http.ResponseWriter, db *sql.DB, campaignId int, id int, payload Update) (rows *sql.Rows, err error) {
	tx, err := db.Begin()
	if err != nil {
		err = fmt.Errorf("open transaction %q", err)
		errmy.Transaction(w, tx)
		return nil, err
	}

	var exists bool
	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM items WHERE id=%d AND campaign_id=%d);", id, campaignId)

	err = db.QueryRow(query).Scan(&exists)
	if err != nil {
		err = fmt.Errorf("chek record: %q", err)
		errmy.Transaction(w, tx)
		return nil, err
	}
	if !exists {
		err = fmt.Errorf("chek record:The record does not exist")
		errmy.TransactionNotFound(w, tx)
		return nil, err
	}

	query = fmt.Sprintf("SELECT * FROM items WHERE id=%d AND campaign_id=%d FOR UPDATE;", id, campaignId)

	_, err = tx.Exec(query)
	if err != nil {
		err = fmt.Errorf("blok %q", err)
		errmy.Transaction(w, tx)
		return nil, err
	}

	query = fmt.Sprintf("UPDATE items SET name='%s', description='%q' WHERE id=%d AND campaign_id=%d;", payload.Name, payload.Description, id, campaignId)

	_, err = tx.Exec(query)
	if err != nil {
		err = fmt.Errorf("set %q", err)
		errmy.Transaction(w, tx)
		return nil, err

	}

	err = tx.Commit()
	if err != nil {
		err = fmt.Errorf("close transaction %q", err)
		errmy.Transaction(w, tx)
		return nil, err
	}

	query = fmt.Sprintf("SELECT * FROM items WHERE id = %d;", id)

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
		errmy.Transaction(w, tx)
		return nil, err
	}

	var exists bool
	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM items WHERE id=%d AND campaign_id=%d);", id, campaignId)

	err = db.QueryRow(query).Scan(&exists)
	if err != nil {
		err = fmt.Errorf("chek record: %q", err)
		errmy.Transaction(w, tx)
		return nil, err
	}
	if !exists {
		err = fmt.Errorf("chek record:The record does not exist")
		errmy.TransactionNotFound(w, tx)
		return nil, err
	}

	query = fmt.Sprintf("SELECT * FROM items WHERE id=%d AND campaign_id=%d FOR UPDATE;", id, campaignId)

	_, err = tx.Exec(query)
	if err != nil {
		err = fmt.Errorf("blok %q", err)
		errmy.Transaction(w, tx)
		return nil, err
	}

	query = fmt.Sprintf("UPDATE items SET  removed=true WHERE id=%d AND campaign_id=%d;", id, campaignId)

	_, err = tx.Exec(query)
	if err != nil {
		err = fmt.Errorf("set %q", err)
		errmy.Transaction(w, tx)
		return nil, err

	}

	err = tx.Commit()
	if err != nil {
		err = fmt.Errorf("close transaction %q", err)
		errmy.Transaction(w, tx)
		return nil, err
	}

	query = fmt.Sprintf("SELECT id,campaign_Id,removed FROM items WHERE id = %d;", id)

	rows, err = db.Query(query)

	if err != nil {
		err = fmt.Errorf("get %q", err)
		return nil, err
	}

	return rows, nil
}

func TransactionGet(w http.ResponseWriter, db *sql.DB) (rows *sql.Rows, err error) {

	tx, err := db.Begin()
	if err != nil {
		err = fmt.Errorf("open transaction %q", err)
		errmy.Transaction(w, tx)
		return nil, err
	}

	_, err = tx.Exec("LOCK TABLE items IN SHARE MODE")
	if err != nil {
		err = fmt.Errorf("blok %q", err)
		errmy.Transaction(w, tx)
		return nil, err
	}

	query := "SELECT * FROM items WHERE removed=false;"

	rows, err = db.Query(query)
	if err != nil {
		err = fmt.Errorf("get %q", err)
		errmy.Transaction(w, tx)
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		err = fmt.Errorf("close transaction %q", err)
		errmy.Transaction(w, tx)
		return nil, err
	}

	return rows, nil
}

func TransactionMigartion(db *sql.DB) (rows *sql.Rows, err error) {

	tx, err := db.Begin()
	if err != nil {
		err = fmt.Errorf("open transaction %q", err)
		tx.Rollback()
		return nil, err
	}

	_, err = tx.Exec("LOCK TABLE items IN SHARE MODE")
	if err != nil {
		err = fmt.Errorf("blok %q", err)
		tx.Rollback()
		return nil, err
	}

	_, err = tx.Exec("LOCK TABLE campaigns IN SHARE MODE")
	if err != nil {
		err = fmt.Errorf("blok %q", err)
		tx.Rollback()
		return nil, err
	}

	query := "SELECT * FROM items;"

	rows, err = db.Query(query)
	if err != nil {
		err = fmt.Errorf("get %q", err)
		tx.Rollback()
		return nil, err
	}

	_, err = tx.Exec("INSERT INTO campaigns (name) VALUES ('new campaign')")
	if err != nil {
		err = fmt.Errorf("set %q", err)
		tx.Rollback()
		return nil, err

	}

	err = tx.Commit()
	if err != nil {
		err = fmt.Errorf("close transaction %q", err)
		tx.Rollback()
		return nil, err
	}

	return rows, nil
}
