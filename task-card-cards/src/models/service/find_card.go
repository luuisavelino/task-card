package service

import (
	"fmt"

	"github.com/luuisavelino/task-card-cards/src/configuration/database"
	"github.com/luuisavelino/task-card-cards/src/configuration/logger"
	"github.com/luuisavelino/task-card-cards/src/configuration/rest_err"
	"github.com/luuisavelino/task-card-cards/src/models"
	"go.uber.org/zap"
)

func (c *cardDomainService) FindCards(actionDomain models.ActionDomainInterface) (map[int]models.CardDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findCards model",
		zap.String("journey", "findCards"),
	)

	db := database.ConnectsWithDatabase()

	selectUserRole, err := db.Query(fmt.Sprintf("select roles.role_ from users join roles on roles.id = users.role_id where users.id = %v", actionDomain.GetUserId()))
	if err != nil {
		logger.Error("Error trying to prepare query", err)
		return nil, rest_err.NewForbiddenError("error to find cards")
	}

	var userRole string
	for selectUserRole.Next() {
		if err = selectUserRole.Scan(&userRole); err != nil {
			logger.Error("Error trying to scan", err)
			return nil, rest_err.NewForbiddenError("error to find cards")
		}
	}

	var queryToGetCards string
	switch userRole {
	case "manager":
		queryToGetCards = "select * from cards"
	case "technician":
		queryToGetCards = fmt.Sprintf("select * from cards where user_id = %v", actionDomain.GetUserId())
	}

	selectCards, err := db.Query(queryToGetCards)
	if err != nil {
		logger.Error("Error trying to prepare query", err)
		return nil, rest_err.NewForbiddenError("Error to find cards")
	}

	cards := make(map[int]models.CardDomainInterface)

	for selectCards.Next() {
		var cardId, userId int
		var title, summary, cardStatus, dueDate string
		if err = selectCards.Scan(&cardId, &title, &summary, &dueDate, &cardStatus, &userId); err != nil {
			logger.Error("Error trying to scan", err)
			return nil, rest_err.NewForbiddenError("Error to scan cards")
		}

		cards[cardId] = models.NewCardDomain(
			title,
			summary,
			dueDate,
			cardStatus,
			userId,
		)
	}

	defer db.Close()

	return cards, nil
}

func (c *cardDomainService) FindCardById(cardId int, actionDomain models.ActionDomainInterface) (map[int]models.CardDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findCard model",
		zap.String("journey", "findCard"),
	)

	db := database.ConnectsWithDatabase()

	isManager, err := manager(db, actionDomain.GetUserId())
	if err != nil {
		logger.Error("Error trying to prepare query", err)
		return nil, rest_err.NewForbiddenError("Error to find card")
	}

	var queryToGetCards string
	switch isManager {
	case true:
		queryToGetCards = fmt.Sprintf("select * from cards where id = %v", cardId)
	case false:
		queryToGetCards = fmt.Sprintf("select * from cards where id = %v and user_id = %v", cardId, actionDomain.GetUserId())
	}

	selectCards, err := db.Query(queryToGetCards)
	if err != nil {
		logger.Error("Error trying to prepare query", err)
		return nil, rest_err.NewForbiddenError("Error to find card")
	}

	cards := make(map[int]models.CardDomainInterface)

	for selectCards.Next() {
		var cardId, userId int
		var title, summary, cardStatus, dueDate string
		if err = selectCards.Scan(&cardId, &title, &summary, &dueDate, &cardStatus, &userId); err != nil {
			logger.Error("Error trying to scan", err)
			return nil, rest_err.NewForbiddenError("Error to find card")
		}

		cards[cardId] = models.NewCardDomain(
			title,
			summary,
			dueDate,
			cardStatus,
			userId,
		)
	}


	defer db.Close()

	return cards, nil
}
