package nats

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func Client() {
	// Подключение к серверу NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Println("Error connecting to NATS server:", err)
		return
	}
	defer nc.Close()

	// Отправка запроса на тему "api.request" и ожидание ответа
	response, err := nc.Request("api.request", []byte("Hello API!"), 1000*time.Millisecond)
	if err != nil {
		fmt.Println("Error sending request to API:", err)
		return
	}
	fmt.Println("Received response:", fmt.Sprintf("%s", response.Data))
}

func Server() {
	// Подключение к серверу NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Println("Error connecting to NATS server:", err)
		return
	}
	defer nc.Close()

	// Обработка запросов на тему "api.request"
	nc.Subscribe("api.request", func(msg *nats.Msg) {
		fmt.Println("Received request:", string(msg.Data))
		// Отправка ответа на тему, указанную в запросе
		nc.Publish(msg.Reply, []byte("Hello Client!"))
	})

	// Бесконечный цикл для ожидания сообщений
	select {}
}
