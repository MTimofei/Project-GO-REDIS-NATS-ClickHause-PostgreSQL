package natspkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
)

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

// Отправка запроса на тему "log" и ожидание ответа
func SetLog(buf *bytes.Buffer, mu *sync.Mutex, nc *nats.Conn) (err error) {
	if reflect.DeepEqual(*buf, bytes.Buffer{}) {
		return fmt.Errorf("nil buf")
	}
	log.Println("in", *buf)

	json, err := json.Marshal(buf)
	if err != nil {
		err = fmt.Errorf("Error JSON API:%q", err)
		log.Println(err)
	}

	_, err = nc.Request("log", json, 1000*time.Millisecond)
	if err != nil {
		err = fmt.Errorf("Error sending request to API:%q", err)
		log.Println(err)
	}
	buf.Reset()
	log.Println("out", *buf)
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
