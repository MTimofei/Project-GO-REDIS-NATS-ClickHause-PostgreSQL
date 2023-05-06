package natspkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/nats-io/nats.go"
)

type logMasseg struct {
	Id  int    `json:"id"`
	Log string `json:"log"`
}

// Подключение к серверу NATS
func Connect() (nc *nats.Conn, err error) {
	nc, err = nats.Connect(nats.DefaultURL)
	if err != nil {
		err = fmt.Errorf("Error connecting to NATS server:%q", err)
		return nil, err
	}

	fmt.Println("connect nats")

	return nc, nil
}

// Отправка "log"
func SetLog(idlog *int, buf *bytes.Buffer, nc *nats.Conn) (err error) {
	logsStrings := buf.String()
	if reflect.DeepEqual(logsStrings, "") {
		log.Println("nil buf")
		return
	}

	logs := strings.Split(logsStrings, "\n")
	buf.Reset()

	for _, i := range logs {
		*idlog++

		logm := logMasseg{
			Id:  *idlog,
			Log: i,
		}

		log.Println(logm)

		json, _ := json.Marshal(&logm)
		if err != nil {
			err = fmt.Errorf("Error JSON API:%q", err)
			log.Println(err)
			return
		}
		log.Printf("i %s\n", json)
		_, err = nc.Request("log", json, 0)
		if err != nil {
			err = fmt.Errorf("Error sending request to API:%q", err)
			log.Println(err)
		}
	}

	buf.Reset()
	log.Println("out", fmt.Sprintf("%s", buf.Bytes()))
	return nil
}

// func Server() {
// 	// Подключение к серверу NATS
// 	nc, err := nats.Connect(nats.DefaultURL)
// 	if err != nil {
// 		fmt.Println("Error connecting to NATS server:", err)
// 		return
// 	}
// 	defer nc.Close()

// 	// Обработка запросов на тему "api.request"
// 	nc.Subscribe("api.request", func(msg *nats.Msg) {
// 		fmt.Println("Received request:", string(msg.Data))
// 		// Отправка ответа на тему, указанную в запросе
// 		nc.Publish(msg.Reply, []byte("Hello Client!"))
// 	})

// 	// Бесконечный цикл для ожидания сообщений
// 	select {}
// }
