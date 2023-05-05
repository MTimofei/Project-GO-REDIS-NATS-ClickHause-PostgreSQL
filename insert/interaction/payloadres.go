package interaction

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

func CreatePayloadItemRes(rows *sql.Rows) (jsonBytes []byte, err error) {
	var item Item
	for rows.Next() {
		defer rows.Close()
		err = rows.Scan(&item.Id, &item.CampaignId, &item.Name, &item.Description, &item.Priority, &item.Removed, &item.CreatedAt)
		if err != nil {
			err = fmt.Errorf("Ошибка декодирования *sql.Rows:%q", err)
			return nil, err
		}
	}

	// log.Println(item)

	jsonBytes, err = json.Marshal(item)
	if err != nil {
		err = fmt.Errorf("Ошибка кодирования JSON:%q", err)
		return nil, err
	}
	return jsonBytes, err
}

func CreatePayloadDeleteRes(rows *sql.Rows) (jsonBytes []byte, err error) {
	var item Delete
	for rows.Next() {
		defer rows.Close()
		err = rows.Scan(&item.Id, &item.CampaignId, &item.Removed)
		if err != nil {
			err = fmt.Errorf("Ошибка декодирования *sql.Rows:%q", err)
			return nil, err
		}
	}

	// log.Println(item)

	jsonBytes, err = json.Marshal(item)
	if err != nil {
		err = fmt.Errorf("Ошибка кодирования JSON:%q", err)
		return nil, err
	}
	return jsonBytes, err
}

func CreatePayloadItemsRes(rows *sql.Rows) (jsonBytes []byte, err error) {
	var items []Item
	for rows.Next() {
		var item Item
		defer rows.Close()
		err = rows.Scan(&item.Id, &item.CampaignId, &item.Name, &item.Description, &item.Priority, &item.Removed, &item.CreatedAt)
		if err != nil {
			err = fmt.Errorf("Ошибка декодирования *sql.Rows:%q", err)
			return nil, err
		}
		items = append(items, item)
	}

	// log.Println(items)

	jsonBytes, err = json.Marshal(items)
	if err != nil {
		err = fmt.Errorf("Ошибка кодирования JSON:%q", err)
		return nil, err
	}
	return jsonBytes, err
}
