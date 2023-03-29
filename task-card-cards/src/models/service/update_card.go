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

	updateCard, err := db.Prepare("Update cards set title=?, summary=?, due_date=? where id=?")

	if err != nil {
		logger.Error("Error trying to prepare query", err)
		return rest_err.NewForbiddenError("error to update card")
	}

	updateCard.Exec(cardDomain.GetTitle(), cardDomain.GetSummary(), cardDomain.GetDueDate(), cardId)
	defer db.Close()
	return nil
}
