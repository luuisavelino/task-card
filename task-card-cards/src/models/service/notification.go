package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/luuisavelino/task-card-cards/src/configuration/database"
)

var server = os.Getenv("server")

type CardNotification struct {
	Id            int               `json:"id"`
	Title         string            `json:"title"`
	CardStatus    string            `json:"card_status"`
	Username      string            `json:"username"`
	ManagersEmail map[string]string `json:"manager_email"`
}

func (c *cardDomainService) SendNotification(cardId int, event string) {

	var cardNotification CardNotification

	cardNotification.Id = cardId
	cardNotification.getCardInfo(cardId)
	cardNotification.getManagersEmail()

	switch event {
	case "update":
		cardNotificationBase64, err := json.Marshal(cardNotification)
		if err != nil {
			log.Println("error when trying to send notification")
			return
		}

		produce(cardNotificationBase64, "update")
	}
}

func (c *CardNotification) getCardInfo(cardId int) {
	db := database.ConnectsWithDatabase()
	selectCardInfo, err := db.Query(fmt.Sprintf("select cards.title, cards.card_status, users.username from users join cards on cards.user_id = users.id where cards.id = %v", cardId))
	if err != nil {
		log.Println("error to get card info")
	}

	var title, cardStatus, username string
	for selectCardInfo.Next() {
		if err = selectCardInfo.Scan(&title, &cardStatus, &username); err != nil {
			log.Println(err)
		}

		c.Title = title
		c.CardStatus = cardStatus
		c.Username = username
	}
}

func (c *CardNotification) getManagersEmail() {
	db := database.ConnectsWithDatabase()
	selectManagersEmail, err := db.Query("select users.username, users.email from roles join users on users.role_id = roles.id  where roles.role_ = 'manager'")
	if err != nil {
		log.Println("error to get managers email")
	}

	var name, email string
	managersEmail := make(map[string]string)

	for selectManagersEmail.Next() {
		if err = selectManagersEmail.Scan(&name, &email); err != nil {
			log.Println(err)
		}

		managersEmail[name] = email
	}

	c.ManagersEmail = managersEmail
}

func produce(msg []byte, topic string) {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": server,
	}
	kafkaProducer, err := kafka.NewProducer(configMap)
	if err != nil {
		fmt.Println("aqqqqq")
		panic(err)
	}

	kafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic: &topic, Partition: kafka.PartitionAny,
		},
		Value: []byte(base64.StdEncoding.EncodeToString(msg)),
	}, nil)
}
