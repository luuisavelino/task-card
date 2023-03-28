package models

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/luuisavelino/task-card-cards/pkg/database"
)

var server = os.Getenv("server")

type CardNotification struct {
	Id            int               `json:"id"`
	Title         string            `json:"title"`
	CardStatus    string            `json:"card_status"`
	Username      string            `json:"username"`
	ManagersEmail map[string]string `json:"manager_email"`
}

func SendNotification(id int, event string) {
	card, err := GetCard(id)
	if err != nil {
		log.Fatalln("error getting card information")
		return
	}
	
	var cardNotification CardNotification
	cardNotification.Infos(card)

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

func (c *CardNotification) Infos(card Card) {
	c.Id = card.Id
	c.Title = card.Title
	c.CardStatus = card.CardStatus

	username, err := getUsername(card.UserId)
	if err != nil {
		return
	}
	c.Username = username

	managerEmail, err := getManagersEmail()
	if err != nil {
		return
	}
	c.ManagersEmail = managerEmail
}



func getManagersEmail() (map[string]string, error) {
	db := database.ConnectsWithDatabase()
	selectManagersEmail, err := db.Query("select users.username, users.email from roles join users on users.role_id = roles.id  where roles.role_ = 'manager'")
	if err != nil {
		log.Println("error to get managers email")
		return nil, err
	}

	managersEmail := make(map[string]string)
	var name, email string
	for selectManagersEmail.Next() {
		if err = selectManagersEmail.Scan(&name, &email); err != nil {
			log.Println(err)
			return nil, errors.New("error to get managers email")
		}

		managersEmail[name] = email
	}

	return managersEmail, nil
}

func getUsername(userId int) (string, error) {
	db := database.ConnectsWithDatabase()
	selectUsername, err := db.Query("select username from users where id = " + fmt.Sprint(userId))
	if err != nil {
		log.Println("error to get username")
		return "", err
	}

	var username string
	for selectUsername.Next() {
		if err = selectUsername.Scan(&username); err != nil {
			log.Println(err)
			return "", errors.New("error to get username")
		}
	}

	return username, nil
}

func produce(msg []byte, topic string) {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": server,
	}
	kafkaProducer, err := kafka.NewProducer(configMap)
	if err != nil {
		panic(err)
	}

	kafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic: &topic, Partition: kafka.PartitionAny,
		},
		Value: []byte(base64.StdEncoding.EncodeToString(msg)),
	}, nil)
}
