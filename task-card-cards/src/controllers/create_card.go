package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/luuisavelino/task-card-cards/src/configuration/logger"
	_ "github.com/luuisavelino/task-card-cards/src/configuration/rest_err"
	"github.com/luuisavelino/task-card-cards/src/configuration/rest_success"
	"github.com/luuisavelino/task-card-cards/src/configuration/validation"
	"github.com/luuisavelino/task-card-cards/src/controllers/model/request"
	"github.com/luuisavelino/task-card-cards/src/models"
)

//	@BasePath	/api/v1

// CreateCard godoc
//
//	@Summary		Create a card
//	@Description	Route to update a card
//	@Tags			cards
//	@Accept			json
//	@Produce		json
//	@Param			card	body		request.CardRequest	true	"Card info"
//	@Success		200		{object}	rest_success.BaseRequestReturn
//	@Failure		400		{object}	rest_err.RestErr
//	@Failure		500		{object}	rest_err.RestErr
//	@Router			/cards [post]
func (cc *cardControllerInterface) CreateCard(c *gin.Context) {
	logger.Info("Init CreateCard controller",
		zap.String("journey", "createCard"),
	)

	var cardRequest request.CardRequest

	if err := c.ShouldBindJSON(&cardRequest); err != nil {
		logger.Error("Error trying to validate card info", err)
		resterr := validation.ValidateCardError(err)
		c.JSON(http.StatusBadRequest, resterr)
		return
	}

	domain := models.NewCardDomain(
		cardRequest.Title,
		cardRequest.Summary,
		cardRequest.DueDate,
		cardRequest.CardStatus,
		cardRequest.UserId,
	)

	if err := cc.service.CreateCard(domain); err != nil {
		logger.Error("Error when trying to create card", err)
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"CreateCard controller executed successfully",
		zap.String("journey", "createCard"))

	c.JSON(http.StatusOK, rest_success.BaseRequestReturn{
		Status: "success", Message: "card created",
	})
}
