package interaction

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func ParseRequestPost(r *http.Request) (campaignId int, payload Post, err error) {

	campaignId, err = strconv.Atoi(r.URL.Query().Get("campaignId"))
	if err != nil {
		err = fmt.Errorf("Ошибка чтения запроса:%q", err)
		return 0, Post{}, err
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		err = fmt.Errorf("Ошибка чтения тела запроса:%q", err)
		return 0, Post{}, err
	}

	err = json.Unmarshal(body, &payload)
	if err != nil {
		err = fmt.Errorf("Ошибка декодирования JSON:%q", err)
		return 0, Post{}, err
	}

	return campaignId, payload, nil
}

func ParseRequestPatch(r *http.Request) (campaignId int, id int, payload Update, err error) {

	campaignId, err = strconv.Atoi(r.URL.Query().Get("campaignId"))
	if err != nil {
		err = fmt.Errorf("Ошибка чтения campaignId: %q", err)
		return 0, 0, Update{}, err
	}
	id, err = strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		err = fmt.Errorf("Ошибка чтения id: %q", err)
		return 0, 0, Update{}, err
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		err = fmt.Errorf("Ошибка чтения тела запроса: %q", err)
		return 0, 0, Update{}, err
	}

	err = json.Unmarshal(body, &payload)
	if err != nil {
		err = fmt.Errorf("Ошибка декодирования JSON: %q", err)
		return 0, 0, Update{}, err
	}
	if payload.Name == "" {
		err = fmt.Errorf("Ошибка декодирования JSON поля Name: %q", err)
		return 0, 0, Update{}, err
	}
	return campaignId, id, payload, nil
}

func ParseRequestDelete(r *http.Request) (campaignId int, id int, err error) {

	campaignId, err = strconv.Atoi(r.URL.Query().Get("campaignId"))
	if err != nil {
		err = fmt.Errorf("Ошибка чтения campaignId: %q", err)
		return 0, 0, err
	}
	id, err = strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		err = fmt.Errorf("Ошибка чтения id: %q", err)
		return 0, 0, err
	}

	return campaignId, id, nil
}
