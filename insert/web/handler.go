package web

import (
	"fmt"
	"log"
	"net/http"

	"git_p/test/insert/db/postgres"
	"git_p/test/insert/db/redispkg"
)

func (cesh *Cesh) handlerPost(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	campaignId, payload, err := postgres.ParseRequestPost(r)
	if err != nil {
		cesh.Log.Println(err)
		return
	}

	rows, err := postgres.TransactionPost(w, cesh.PostgreasQL, campaignId, payload)
	if err != nil {
		cesh.Log.Println("ERROR TransactionPost:", err)
		return
	}

	item, err := postgres.CreateItem(rows)
	if err != nil {
		cesh.Log.Println("ERROR CreateItem:", err)
		return
	}

	jsonBytes, err := postgres.CreatePayloadItemRes(&item)
	if err != nil {
		cesh.Log.Println("Ошибка записи JSON:", err)
		return
	}

	_, err = w.Write(jsonBytes)
	if err != nil {
		cesh.Log.Println(fmt.Errorf("Ошибка записи JSON:%q", err))
	}
	cesh.Log.Println("client:", r.UserAgent(), "method:", r.Method, "id:", item.Id, "campaignId:", campaignId, "new name:", payload.Name)
}

func (cesh *Cesh) handlerPatch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	campaignId, id, payload, err := postgres.ParseRequestPatch(r)
	if err != nil {
		cesh.Log.Println("ParseRequestPatch:", err)
		return
	}

	rows, err := postgres.TransactionPatch(w, cesh.PostgreasQL, campaignId, id, payload)
	if err != nil {
		cesh.Log.Println("ERROR TransactionPatch:", err)
		return
	}

	item, err := postgres.CreateItem(rows)
	if err != nil {
		cesh.Log.Println("ERROR CreateItem:", err)
		return
	}

	jsonBytes, err := postgres.CreatePayloadItemRes(&item)
	if err != nil {
		cesh.Log.Println("Ошибка записи JSON:", err)
		return
	}

	_, err = w.Write(jsonBytes)
	if err != nil {
		cesh.Log.Println(fmt.Errorf("Ошибка записи JSON:%q", err))
		return
	}

	cesh.Log.Println("client:", r.UserAgent(), r.Method, "id:", id, "campaignId:", campaignId, "new nate", payload.Name, "description", payload.Description)
}

func (cesh *Cesh) handlerDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	campaignId, id, err := postgres.ParseRequestDelete(r)
	if err != nil {
		cesh.Log.Println("ParseRequestDelete:", err)
		return
	}

	rows, err := postgres.TransactionDelete(w, cesh.PostgreasQL, campaignId, id)
	if err != nil {
		cesh.Log.Println("ERROR TransactionDelete:", err)
		return
	}

	del, err := postgres.CreateDelete(rows)
	if err != nil {
		cesh.Log.Println("ERROR CreateDelete:", err)
		return
	}

	jsonBytes, err := postgres.CreatePayloadDeleteRes(&del)
	if err != nil {
		cesh.Log.Println("Ошибка записи JSON:", err)
		return
	}

	_, err = w.Write(jsonBytes)
	if err != nil {
		log.Println(fmt.Errorf("Ошибка записи JSON:%q", err))
	}

	cesh.Log.Println("client:", r.UserAgent(), r.Method, "id:", id, "campaignId:", campaignId, "remuved:", del.Removed)

	rdb := redispkg.ConectRedis()

	rows, err = postgres.TransactionGet(w, cesh.PostgreasQL)
	if err != nil {
		cesh.Log.Println("ERROR TransactionGet:", err)
		return
	}

	items, err := postgres.CreateItems(rows)
	if err != nil {
		cesh.Log.Println("ERROR CreateItems:", err)
		return
	}

	jsonBytes, err = postgres.CreatePayloadItemsRes(&items)
	if err != nil {
		cesh.Log.Println("Ошибка записи JSON:", err)
		return
	}

	err = redispkg.Set(rdb, jsonBytes)
	if err != nil {
		cesh.Log.Println("Ошибка записи", err)
		return
	}
}

func (cesh *Cesh) handlerGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	var jsonBytes []byte

	rdb := redispkg.ConectRedis()
	result, err := redispkg.Get(rdb)

	cesh.Log.Println("client:", r.UserAgent(), r.Method)

	if err != nil {
		cesh.Log.Println(err)
		rows, err := postgres.TransactionGet(w, cesh.PostgreasQL)
		if err != nil {
			cesh.Log.Println("ERROR TransactionGet:", err)
			return
		}

		items, err := postgres.CreateItems(rows)
		if err != nil {
			cesh.Log.Println("ERROR CreateItems:", err)
			return
		}

		jsonBytes, err = postgres.CreatePayloadItemsRes(&items)
		if err != nil {
			cesh.Log.Println("Ошибка записи JSON:", err)
			return
		}

		err = redispkg.Set(rdb, jsonBytes)
		if err != nil {
			cesh.Log.Println("Ошибка записи в redis:", err)
			return
		}
	} else {
		jsonBytes = result
	}

	_, err = w.Write(jsonBytes)
	if err != nil {
		cesh.Log.Println(fmt.Errorf("Ошибка записи JSON:%q", err))

	}
	cesh.Log.Println("client:", r.UserAgent(), r.Method)
}
