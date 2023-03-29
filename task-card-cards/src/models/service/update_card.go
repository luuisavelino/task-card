package service

import (
	"github.com/luuisavelino/task-card-cards/src/configuration/database"
	"github.com/luuisavelino/task-card-cards/src/configuration/logger"
	"github.com/luuisavelino/task-card-cards/src/configuration/rest_err"
	"github.com/luuisavelino/task-card-cards/src/models"
	"go.uber.org/zap"
)

func (c *cardDomainService) UpdateCardInfo(cardId int, cardDomain models.CardDomainInterface) *rest_err.RestErr {
	logger.Info("Init updateCard model",
		zap.String("journey", "updateCard"),
	)

	db := database.ConnectsWithDatabase()

	isManager, err := manager(db, cardDomain.GetUserId())
	if err != nil {
		logger.Error("Error trying to prepare query", err)
		return rest_err.NewForbiddenError("Error when deleting card")
	}

	isCardOwner, err := cardOwner(db, cardId, cardDomain.GetUserId())
	if err != nil {
		logger.Error("Error trying to prepare query", err)
		return rest_err.NewForbiddenError("Error when deleting card")
	}

	if isManager || isCardOwner {

		updateCard, err := db.Prepare("Update cards set title=?, summary=?, due_date=? where id=?")

		if err != nil {
			logger.Error("Error trying to prepare query", err)
			return rest_err.NewForbiddenError("error to update card")
		}

		updateCard.Exec(cardDomain.GetTitle(), cardDomain.GetSummary(), cardDomain.GetDueDate(), cardId)
		defer db.Close()
		return nil
	}

	logger.Error("No permission to delete card", err)
	return rest_err.NewForbiddenError("No permission to delete card")
}
