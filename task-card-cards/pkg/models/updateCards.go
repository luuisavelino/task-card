package models

import (
	"errors"
	"log"

	"github.com/luuisavelino/task-card-cards/pkg/database"
)

func UpdateCard(card Card) error {
	db := database.ConnectsWithDatabase()

	updateCard, err := db.Prepare("Update cards set title=?, summary=?, card_status=?, due_date=? where id=?")
	if err != nil {
		log.Println(err)
		return errors.New(genericErrToUser)
	}

	updateCard.Exec(card.Title, card.Summary, card.CardStatus, card.DueDate, card.Id)
	defer db.Close()
	return nil
}
