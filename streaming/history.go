package main

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type viewedMessage struct {
	VideoPath string `json:"videoPath"`
}

func sendViewedMessage(videoPath string) {
	msg, err := json.Marshal(viewedMessage{VideoPath: videoPath})
	if err != nil {
		log.Print(err)
		return
	}
	messageChannel.ExchangeDeclare("viewed", "fanout", false, false, false, false, nil)

	err = messageChannel.Publish("viewed", "", false, false,
		amqp.Publishing{ContentType: "text/plain", Body: msg})
	if err != nil {
		log.Print(err)
		return
	}
}
