package view

import (
	"github.com/luuisavelino/task-card-cards/src/controllers/model/response"
	"github.com/luuisavelino/task-card-cards/src/models"
)

func ConvertDomainToResponse(usernameDomains map[int]models.CardDomainInterface) []response.CardResponse {
	var cards []response.CardResponse
	for cardId, userDomain := range usernameDomains {
		cards = append(cards, response.CardResponse{
			Id:       cardId,
			Title:    userDomain.GetTitle(),
			Summary:  userDomain.GetSummary(),
			DueDate:  userDomain.GetDueDate(),
			CardStatus: userDomain.GetCardStatus(),
			UserId:   userDomain.GetUserId(),
		})
	}

	return cards
}
