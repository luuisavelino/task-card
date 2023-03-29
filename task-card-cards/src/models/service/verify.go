package service

import (
	"database/sql"
	"fmt"

	"github.com/luuisavelino/task-card-cards/src/configuration/logger"
)

func manager(db *sql.DB, userId int) (bool, error) {
	executeIsManager, err := db.Query(fmt.Sprint("select True from users join roles on roles.id = users.role_id where users.id = %s and roles.role_  = 'manager'", userId))
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
	executeIsCardOwner, err := db.Query(fmt.Sprint("select True from cards where id = %s and user_id = %s", cardId, userId))
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
