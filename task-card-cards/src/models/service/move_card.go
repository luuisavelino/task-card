package service

import (
	"github.com/luuisavelino/task-card-cards/src/configuration/database"
	"github.com/luuisavelino/task-card-cards/src/configuration/logger"
	"github.com/luuisavelino/task-card-cards/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (c *cardDomainService) MoveCard(cardId int) *rest_err.RestErr {
	logger.Info("Init moveCard model",
		zap.String("journey", "moveCard"),
	)

	db := database.ConnectsWithDatabase()

	// Get atual Status
	// Move to next Status


	updateCard, err := db.Prepare("Update cards set card_status = ? where id = ?")
	if err != nil {
		logger.Error("Error trying to prepare query", err)
		return rest_err.NewForbiddenError("error to move card")
	}

	updateCard.Exec("", cardId)
	defer db.Close()
	return nil
}
