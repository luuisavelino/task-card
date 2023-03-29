package service

import (
	"github.com/luuisavelino/task-card-cards/src/configuration/database"
	"github.com/luuisavelino/task-card-cards/src/configuration/logger"
	"github.com/luuisavelino/task-card-cards/src/configuration/rest_err"
	"github.com/luuisavelino/task-card-cards/src/models"
	"go.uber.org/zap"
)

func (c *cardDomainService) DeleteCard(cardId int, actionDomain models.ActionDomainInterface) *rest_err.RestErr {
	logger.Info("Init deleteCard model",
		zap.String("journey", "deleteCard"),
	)

	db := database.ConnectsWithDatabase()

	isManager, err := manager(db, actionDomain.GetUserId())
	if err != nil {
		logger.Error("Error trying to prepare query", err)
		return rest_err.NewForbiddenError("Error when deleting card")
	}

	if isManager {
		deleteCard, err := db.Prepare("delete from cards where id = ?")
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
