package interaction

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
)

func CreatePayloadPostRes(rows *sql.Rows) (jsonBytes []byte, err error) {
	var item Item
	for rows.Next() {
		defer rows.Close()
		err = rows.Scan(&item.Id, &item.CampaignId, &item.Name, &item.Description, &item.Priority, &item.Removed, &item.CreatedAt)
		if err != nil {
			err = fmt.Errorf("Ошибка декодирования *sql.Rows:%q", err)
			return nil, err
		}
	}

	log.Println(item)

	jsonBytes, err = json.Marshal(item)
	if err != nil {
		err = fmt.Errorf("Ошибка кодирования JSON:%q", err)
		return nil, err
	}
	return jsonBytes, err
}
