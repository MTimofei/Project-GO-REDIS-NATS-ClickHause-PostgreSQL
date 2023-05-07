package postgres

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

// создает json обьект обрабатывая ответ CreateItem
func CreatePayloadItemRes(item *Item) (jsonBytes []byte, err error) {
	jsonBytes, err = json.Marshal(*item)
	if err != nil {
		err = fmt.Errorf("Ошибка кодирования JSON:%q", err)
		return nil, err
	}
	return jsonBytes, err
}

// создает json обьект обрабатывая ответ CreateDelete
func CreatePayloadDeleteRes(item *Delete) (jsonBytes []byte, err error) {
	jsonBytes, err = json.Marshal(*item)
	if err != nil {
		err = fmt.Errorf("Ошибка кодирования JSON:%q", err)
		return nil, err
	}
	return jsonBytes, err
}

// создает json обьект обрабатывая ответ CreateItems
func CreatePayloadItemsRes(items *[]Item) (jsonBytes []byte, err error) {
	jsonBytes, err = json.Marshal(items)
	if err != nil {
		err = fmt.Errorf("Ошибка кодирования JSON:%q", err)
		return nil, err
	}
	return jsonBytes, err
}

// разшифровывает ответ от TransactionGet и TransactionMigartion
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

// разшифровывает ответ от TransactionDelete
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

// разшифровывает ответ от TransactionPost и TransactionPatch
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
