package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/task-card-cards/src/models/service"
)

type cardControllerInterface struct {
	service  service.CardDomainService
}

type CardControllerInterface interface {
	FindCards(c *gin.Context)
	FindCardById(c *gin.Context)

	UpdateCard(c *gin.Context)
	MoveCard(c *gin.Context)

	DeleteCard(c *gin.Context)
	CreateCard(c *gin.Context)
}

func NewCardControllerInterface(serviceInterface service.CardDomainService) CardControllerInterface {
	return &cardControllerInterface{
		service: serviceInterface,
	}
}
