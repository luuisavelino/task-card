package service

import (
	"errors"
	"log"

	"github.com/luuisavelino/task-card-cards/src/configuration/database"
	"github.com/luuisavelino/task-card-cards/src/configuration/rest_err"
	"github.com/luuisavelino/task-card-cards/src/models"
)

func CreateNewCard(card Card) error {
	db := database.ConnectsWithDatabase()

	insertCardIntoDatabase, err := db.Prepare("insert into cards(title, summary, due_date, card_status, user_id) values(?, ?, ?, ?, ?)")
	if err != nil {
		log.Println(err)
		return errors.New(genericErrToUser)
	}

	insertCardIntoDatabase.Exec(card.Title, card.Summary, card.DueDate, card.CardStatus, card.UserId)

	defer db.Close()
	return nil
}

func (c *cardDomainService) CreateCard(cardDomain models.CardDomainInterface) *rest_err.RestErr {
	return nil
}
