package interaction

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func ParseRequestPost(r *http.Request) (campaignId int, payload Post, err error) {

	qs := strings.Split(r.URL.RawQuery, "=")
	campaignId, err = strconv.Atoi(qs[1])
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
