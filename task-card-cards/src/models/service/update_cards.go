package service

import (
	"errors"
	"log"

	"github.com/luuisavelino/task-card-cards/src/configuration/database"
	"github.com/luuisavelino/task-card-cards/src/configuration/rest_err"
	"github.com/luuisavelino/task-card-cards/src/models"
)

func (c *cardDomainService) UpdateCardInfo(cardId int, cardDomain models.CardDomainInterface) *rest_err.RestErr {
	return nil
}

func (c *cardDomainService) MoveCard(cardId int, cardDomain models.CardDomainInterface) *rest_err.RestErr {
	return nil
}

func UpdateCardInfo(card Card) error {
	db := database.ConnectsWithDatabase()

	updateCard, err := db.Prepare("Update cards set title=?, summary=?, due_date=? where id=?")
	if err != nil {
		log.Println(err)
		return errors.New(genericErrToUser)
	}

	updateCard.Exec(card.Title, card.Summary, card.DueDate, card.Id)
	defer db.Close()
	return nil
}

func UpdateCardStatus(card Card) error {
	db := database.ConnectsWithDatabase()

	updateCard, err := db.Prepare("Update cards set card_status=? where id=?")
	if err != nil {
		log.Println(err)
		return errors.New(genericErrToUser)
	}

	updateCard.Exec(card.CardStatus, card.Id)
	defer db.Close()
	return nil
}
