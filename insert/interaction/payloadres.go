package interaction

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

func CreatePayloadItemRes(item *Item) (jsonBytes []byte, err error) {
	jsonBytes, err = json.Marshal(*item)
	if err != nil {
		err = fmt.Errorf("Ошибка кодирования JSON:%q", err)
		return nil, err
	}
	return jsonBytes, err
}

func CreatePayloadDeleteRes(item *Delete) (jsonBytes []byte, err error) {
	jsonBytes, err = json.Marshal(*item)
	if err != nil {
		err = fmt.Errorf("Ошибка кодирования JSON:%q", err)
		return nil, err
	}
	return jsonBytes, err
}

func CreatePayloadItemsRes(items *[]Item) (jsonBytes []byte, err error) {
	jsonBytes, err = json.Marshal(items)
	if err != nil {
		err = fmt.Errorf("Ошибка кодирования JSON:%q", err)
		return nil, err
	}
	return jsonBytes, err
}

func CreateItems(rows *sql.Rows) (items []Item, err error) {
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
	return items, err
}

func CreateDelete(rows *sql.Rows) (item Delete, err error) {
	for rows.Next() {
		defer rows.Close()
		err = rows.Scan(&item.Id, &item.CampaignId, &item.Removed)
		if err != nil {
			err = fmt.Errorf("Ошибка декодирования *sql.Rows:%q", err)
			return Delete{}, err
		}
	}

	return item, nil
}

func CreateItem(rows *sql.Rows) (item Item, err error) {
	for rows.Next() {
		defer rows.Close()
		err = rows.Scan(&item.Id, &item.CampaignId, &item.Name, &item.Description, &item.Priority, &item.Removed, &item.CreatedAt)
		if err != nil {
			err = fmt.Errorf("Ошибка декодирования *sql.Rows:%q", err)
			return Item{}, err
		}
	}

	return item, nil
}
