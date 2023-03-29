package service

import (
	"github.com/luuisavelino/task-card-cards/src/configuration/database"
	"github.com/luuisavelino/task-card-cards/src/configuration/logger"
	"github.com/luuisavelino/task-card-cards/src/configuration/rest_err"
	"github.com/luuisavelino/task-card-cards/src/models"
	"go.uber.org/zap"
)

func (c *cardDomainService) MoveCard(cardId int, actionDomain models.ActionDomainInterface) *rest_err.RestErr {
	logger.Info("Init moveCard model",
		zap.String("journey", "moveCard"),
	)

	db := database.ConnectsWithDatabase()

	// Get atual Status
	// Move to next Status

	isManager, err := manager(db, actionDomain.GetUserId())
	if err != nil {
		logger.Error("Error trying to prepare query", err)
		return rest_err.NewForbiddenError("Error when deleting card")
	}

	isCardOwner, err := cardOwner(db, cardId, actionDomain.GetUserId())
	if err != nil {
		logger.Error("Error trying to prepare query", err)
		return rest_err.NewForbiddenError("Error when deleting card")
	}

	if isManager || isCardOwner {
		updateCard, err := db.Prepare("Update cards set card_status = ? where id = ?")
		if err != nil {
			logger.Error("Error trying to prepare query", err)
			return rest_err.NewForbiddenError("error to move card")
		}

		updateCard.Exec("to do", cardId)
		defer db.Close()
		return nil
	}

	logger.Error("No permission to delete card", err)
	return rest_err.NewForbiddenError("No permission to delete card")
}
