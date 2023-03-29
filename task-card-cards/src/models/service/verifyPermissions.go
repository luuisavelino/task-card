package service

import (
	"database/sql"
	"fmt"

	"github.com/luuisavelino/task-card-cards/src/configuration/logger"
)

func manager(db *sql.DB, userId int) (bool, error) {
	executeIsManager, err := db.Query(fmt.Sprintf("select True from users join roles on roles.id = users.role_id where users.id = %v and roles.role_  = 'manager'", userId))
	if err != nil {
		logger.Error("Error trying to prepare query", err)
		return false, err
	}

	var isManager bool
	for executeIsManager.Next() {
		executeIsManager.Scan(&isManager)
	}

	return isManager, nil
}

func cardOwner(db *sql.DB, cardId, userId int) (bool, error) {
	executeIsCardOwner, err := db.Query(fmt.Sprintf("select True from cards where id = %v and user_id = %v", cardId, userId))
	if err != nil {
		logger.Error("Error trying to prepare query", err)
		return false, err
	}

	var isCardOwner bool
	for executeIsCardOwner.Next() {
		executeIsCardOwner.Scan(&isCardOwner)
	}

	return isCardOwner, nil
}

func getCardStatus(db *sql.DB, cardId int) (string, error) {
	//

	executeGetCardStatus, err := db.Query(fmt.Sprintf("select card_status from cards where id = %v", cardId))
	if err != nil {
		logger.Error("Error trying to prepare query", err)
		return "", err
	}

	var cardStatus string
	for executeGetCardStatus.Next() {
		executeGetCardStatus.Scan(&cardStatus)
	}

	return cardStatus, nil
}
