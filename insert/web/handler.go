package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gomodule/redigo/redis"

	"git_p/test/insert/interaction"
	"git_p/test/pkg/connect/bd"
)

func (cesh *Cesh) handlerPost(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	campaignId, payload, err := interaction.ParseRequestPost(r)
	if err != nil {
		log.Println(err)
		return
	}

	rows, err := interaction.TransactionPost(w, cesh.PostgreasQL, campaignId, payload)
	if err != nil {
		log.Println("ERROR TransactionPost:", err)
		return
	}

	item, err := interaction.CreateItem(rows)
	if err != nil {
		log.Println("ERROR CreateItem:", err)
		return
	}

	jsonBytes, err := interaction.CreatePayloadItemRes(&item)
	if err != nil {
		log.Println("Ошибка записи JSON:", err)
		return
	}

	_, err = w.Write(jsonBytes)
	if err != nil {
		log.Println(fmt.Errorf("Ошибка записи JSON:%q", err))
		return
	}
}

func (cesh *Cesh) handlerPatch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	campaignId, id, payload, err := interaction.ParseRequestPatch(r)
	if err != nil {
		log.Println("ParseRequestPatch:", err)
		return
	}

	rows, err := interaction.TransactionPatch(w, cesh.PostgreasQL, campaignId, id, payload)
	if err != nil {
		log.Println("ERROR TransactionPatch:", err)
		return
	}

	item, err := interaction.CreateItem(rows)
	if err != nil {
		log.Println("ERROR CreateItem:", err)
		return
	}

	jsonBytes, err := interaction.CreatePayloadItemRes(&item)
	if err != nil {
		log.Println("Ошибка записи JSON:", err)
		return
	}

	_, err = w.Write(jsonBytes)
	if err != nil {
		log.Println(fmt.Errorf("Ошибка записи JSON:%q", err))
		return
	}
}

func (cesh *Cesh) handlerDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	campaignId, id, err := interaction.ParseRequestDelete(r)
	if err != nil {
		log.Println("ParseRequestDelete:", err)
		return
	}

	rows, err := interaction.TransactionDelete(w, cesh.PostgreasQL, campaignId, id)
	if err != nil {
		log.Println("ERROR TransactionDelete:", err)
		return
	}

	del, err := interaction.CreateDelete(rows)
	if err != nil {
		log.Println("ERROR CreateDelete:", err)
		return
	}

	jsonBytes, err := interaction.CreatePayloadDeleteRes(&del)
	if err != nil {
		log.Println("Ошибка записи JSON:", err)
		return
	}

	conn, err := bd.ConectRedis(cesh.RediaAddr)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	rows, err = interaction.TransactionGet(w, cesh.PostgreasQL)
	if err != nil {
		log.Println("ERROR TransactionGet:", err)
		return
	}

	items, err := interaction.CreateItems(rows)
	if err != nil {
		log.Println("ERROR CreateItems:", err)
		return
	}

	jsonBytes, err = interaction.CreatePayloadItemsRes(&items)
	if err != nil {
		log.Println("Ошибка записи JSON:", err)
		return
	}
	conn.Do("SET", "get", jsonBytes)

	_, err = w.Write(jsonBytes)
	if err != nil {
		log.Println(fmt.Errorf("Ошибка записи JSON:%q", err))
		return
	}

}

func (cesh *Cesh) handlerGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	conn, err := bd.ConectRedis(cesh.RediaAddr)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	var jsonBytes []byte

	result, err := redis.Bytes(conn.Do("GET", "get"))
	if err != nil {
		log.Println(err)
		rows, err := interaction.TransactionGet(w, cesh.PostgreasQL)
		if err != nil {
			log.Println("ERROR TransactionGet:", err)
			return
		}

		items, err := interaction.CreateItems(rows)
		if err != nil {
			log.Println("ERROR CreateItems:", err)
			return
		}

		jsonBytes, err = interaction.CreatePayloadItemsRes(&items)
		if err != nil {
			log.Println("Ошибка записи JSON:", err)
			return
		}
		conn.Do("SET", "get", jsonBytes)
	} else {
		jsonBytes = result
	}

	_, err = w.Write(jsonBytes)
	if err != nil {
		log.Println(fmt.Errorf("Ошибка записи JSON:%q", err))
		return
	}
}
