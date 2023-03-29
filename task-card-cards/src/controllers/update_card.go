package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/task-card-cards/src/configuration/logger"
	_ "github.com/luuisavelino/task-card-cards/src/configuration/rest_err"
	"github.com/luuisavelino/task-card-cards/src/configuration/rest_success"
	"github.com/luuisavelino/task-card-cards/src/configuration/validation"
	"github.com/luuisavelino/task-card-cards/src/controllers/model/request"
	"github.com/luuisavelino/task-card-cards/src/models"
	"go.uber.org/zap"
)

//	@BasePath	/api/v1

// UpdateCard godoc
//	@Summary		Update a card
//	@Description	Update a card
//	@Tags			cards
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Card id"
//	@Param			card	body		request.CardRequest	true	"Card info"
//	@Success		200		{object}	rest_success.BaseRequestReturn
//	@Failure		400		{object}	rest_err.RestErr
//	@Failure		500		{object}	rest_err.RestErr
//	@Router			/cards/{id} [put]
func (cc *cardControllerInterface) UpdateCard(c *gin.Context) {
	logger.Info("Init UpdateCard controller",
		zap.String("journey", "updateCard"),
	)

	cardId, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		logger.Error("Error trying to get card id", err)
		resterr := validation.ValidateCardError(err)
		c.JSON(http.StatusBadRequest, resterr)
		return
	}

	var cardRequest request.CardRequest

	if err := c.ShouldBindJSON(&cardRequest); err != nil {
		logger.Error("Error trying to validate card info", err)
		resterr := validation.ValidateCardError(err)
		c.JSON(http.StatusBadRequest, resterr)
		return
	}

	cardDomain := models.NewCardDomain(
		cardRequest.Title,
		cardRequest.Summary,
		cardRequest.DueDate,
		cardRequest.CardStatus,
		cardRequest.UserId,
	)

	if err := cc.service.UpdateCardInfo(cardId, cardDomain); err != nil {
		logger.Error("Error when trying to update card", err)
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"UpdateCard controller executed successfully",
		zap.String("journey", "updateCard"))

	c.JSON(http.StatusOK, rest_success.BaseRequestReturn{
		Status: "success", Message: "card updated",
	})
}
