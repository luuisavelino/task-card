package models

import (
	"errors"
	"fmt"
	"log"

	"github.com/luuisavelino/task-card-cards/pkg/database"
)

func DeleteCard(cardId, userId int) error {
	db := database.ConnectsWithDatabase()

	executeIsManager, err := db.Query("select True from users join roles on roles.id = users.role_id where users.id = " + fmt.Sprint(userId) + " and roles.role_  = 'manager'")
	if err != nil {
		log.Println(err)
		return errors.New("error when deleting card")
	}

	var isManager bool
	for executeIsManager.Next() {
		executeIsManager.Scan(&isManager)
	}

	if isManager {
		deleteCard, err := db.Prepare("delete from cards where id=?")
		if err != nil {
			log.Println(err)
			return errors.New("error when deleting card")
		}

		deleteCard.Exec(cardId)
		defer db.Close()
		return nil
	}

	return errors.New("no permission to delete card")
}
