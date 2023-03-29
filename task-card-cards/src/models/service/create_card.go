package service

import (
	"github.com/luuisavelino/task-card-cards/src/configuration/database"
	"github.com/luuisavelino/task-card-cards/src/configuration/logger"
	"github.com/luuisavelino/task-card-cards/src/configuration/rest_err"
	"github.com/luuisavelino/task-card-cards/src/models"
	"go.uber.org/zap"
)

func (c *cardDomainService) CreateCard(cardDomain models.CardDomainInterface) *rest_err.RestErr {
	logger.Info("Init createCard model",
		zap.String("journey", "createCard"),
	)

	db := database.ConnectsWithDatabase()

	insertCardIntoDatabase, err := db.Prepare("insert into cards(title, summary, due_date, card_status, user_id) values(?, ?, ?, ?, ?)")
	if err != nil {
		logger.Error("Error trying to prepare query", err)
		return rest_err.NewForbiddenError("error to insert values")
	}

	insertCardIntoDatabase.Exec(cardDomain.GetTitle(), cardDomain.GetSummary(), cardDomain.GetDueDate(), cardDomain.GetCardStatus(), cardDomain.GetUserId())

	defer db.Close()
	return nil
}
