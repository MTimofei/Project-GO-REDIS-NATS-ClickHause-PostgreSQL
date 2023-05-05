package web

import (
	"fmt"
	"log"
	"net/http"

	"git_p/test/insert/interaction"
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

	// log.Println(campaignId)
	// log.Println(payload.Name)

	rows, err := interaction.TransactionPost(w, cesh.PostgreasQL, campaignId, payload)
	if err != nil {
		log.Println("ERROR TransactionPost:", err)
		return
	}

	jsonBytes, err := interaction.CreatePayloadItemRes(rows)
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
		log.Println(err)
		return
	}

	log.Println(campaignId, id, payload.Name, payload.Description, err)

	rows, err := interaction.TransactionPatch(w, cesh.PostgreasQL, campaignId, id, payload)
	if err != nil {
		log.Println("ERROR TransactionPost:", err)
		return
	}

	jsonBytes, err := interaction.CreatePayloadItemRes(rows)
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
}

func (cesh *Cesh) handlerGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
