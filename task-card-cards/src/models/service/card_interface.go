package service

import (
	"github.com/luuisavelino/task-card-cards/src/configuration/rest_err"
	"github.com/luuisavelino/task-card-cards/src/models"
)

func NewCardDomainService() CardDomainService {
	return &cardDomainService{}
}

type cardDomainService struct {
}

type CardDomainService interface {
	FindCards() (map[int]models.CardDomainInterface, *rest_err.RestErr)
	FindCardById(int) (map[int]models.CardDomainInterface, *rest_err.RestErr)

	UpdateCardInfo(int, models.CardDomainInterface) *rest_err.RestErr
	MoveCard(int, models.CardDomainInterface) *rest_err.RestErr

	CreateCard(models.CardDomainInterface) *rest_err.RestErr
	DeleteCard(int) *rest_err.RestErr
}
