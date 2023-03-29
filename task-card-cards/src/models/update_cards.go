package models

import (
	"errors"
	"log"

	"github.com/luuisavelino/task-card-cards/src/configuration/database"
)

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
