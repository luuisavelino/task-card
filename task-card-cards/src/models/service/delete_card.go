package service

import (
	"fmt"

	"github.com/luuisavelino/task-card-cards/src/configuration/database"
	"github.com/luuisavelino/task-card-cards/src/configuration/logger"
	"github.com/luuisavelino/task-card-cards/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (c *cardDomainService) DeleteCard(cardId, userId int) *rest_err.RestErr {
	logger.Info("Init deleteCard model",
		zap.String("journey", "deleteCard"),
	)

	db := database.ConnectsWithDatabase()

	executeIsManager, err := db.Query("select True from users join roles on roles.id = users.role_id where users.id = " + fmt.Sprint(userId) + " and roles.role_  = 'manager'")
	if err != nil {
		logger.Error("Error trying to prepare query", err)
		return rest_err.NewForbiddenError("Error when deleting card")
	}

	var isManager bool
	for executeIsManager.Next() {
		executeIsManager.Scan(&isManager)
	}

	if isManager {
		deleteCard, err := db.Prepare("delete from cards where id=?")
		if err != nil {
			logger.Error("Error trying to prepare query", err)
			return rest_err.NewForbiddenError("Error when deleting card")
		}

		deleteCard.Exec(cardId)
		defer db.Close()
		return nil
	}

	logger.Error("No permission to delete card", err)
	return rest_err.NewForbiddenError("No permission to delete card")
}
