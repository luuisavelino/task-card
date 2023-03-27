package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var server = os.Getenv("server")
var topicConsume = []string{os.Getenv("topicConsume")}
var groupId = os.Getenv("groupId")
var autoOffSetReset = os.Getenv("autoOffSetReset")
var serviceAccount = os.Getenv("serviceAccount")

type Email struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Text    string   `json:"text"`
}

type CardNotification struct {
	Id            int               `json:"id"`
	Title         string            `json:"title"`
	CardStatus    string            `json:"card_status"`
	Username      string            `json:"username"`
	ManagersEmail map[string]string `json:"manager_email"`
}

func main() {
	msgChan := make(chan *kafka.Message)

	go Consume(topicConsume, msgChan)

	for msg := range msgChan {
		var email Email

		card := getCard(string(msg.Value))
		email.Build(card)
		email.Send()
	}
}

func (e Email) Send() {
	fmt.Println("\n:: ============== EMAIL ============== ::")
	fmt.Println("From:\t", e.From)
	fmt.Println("To:\t", e.To)
	fmt.Println("Subject:", e.Subject)
	fmt.Println("Text:\t", e.Text)
	fmt.Println("\n:: =================================== ::\n")
}

func (e *Email) Build(c CardNotification) {
	var emails []string
	for _, email := range c.ManagersEmail {
		emails = append(emails, email)
	}

	e.From = serviceAccount
	e.To = emails
	e.Subject = "Card update"
	e.Text = "Card \"" + c.Title + "\" (id: " + fmt.Sprint(c.Id) + ") has been moved to \"" + c.CardStatus + "\" by user " + c.Username
}

func getCard(data string) CardNotification {
	var card CardNotification

	rawDecodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(rawDecodedData, &card)

	return card
}

func Consume(topics []string, msgChan chan *kafka.Message) {
	kafkaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": server,
		"group.id":          groupId,
		"auto.offset.reset": autoOffSetReset,
	})
	if err != nil {
		panic(err)
	}

	if err := kafkaConsumer.SubscribeTopics(topics, nil); err != nil {
		panic(err)
	}

	for {
		msg, err := kafkaConsumer.ReadMessage(-1)
		if err == nil {
			msgChan <- msg
		}
	}
}
