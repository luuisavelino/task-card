package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/task-card-cards/src/configuration/logger"
	"github.com/luuisavelino/task-card-cards/src/configuration/rest_success"
	"github.com/luuisavelino/task-card-cards/src/configuration/validation"
	"github.com/luuisavelino/task-card-cards/src/controllers/model/request"
	"github.com/luuisavelino/task-card-cards/src/models"
	"go.uber.org/zap"
)

// @BasePath /api/v1

// MoveCard godoc
// @Summary Move a card
// @Description Move a card
// @Tags cards
// @Accept  json
// @Produce  json
// @Param id path int true "Card id"
// @Param card_status formData string true "Status of the card"
// @Success 200 {object} rest_success.BaseRequestReturn
// @Failure 400 {object} rest_success.BaseRequestReturn
// @Router /cards/{id} [put]
func (cc *cardControllerInterface) MoveCard(c *gin.Context) {
	logger.Info("Init MoveCard controller",
		zap.String("journey", "moveCard"),
	)

	cardId, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		logger.Error("Error trying to move card id", err)
		resterr := validation.ValidateCardError(err)
		c.JSON(http.StatusBadRequest, resterr)
		return
	}

	var actionRequest request.ActionRequest

	if err := c.ShouldBindJSON(&actionRequest); err != nil {
		logger.Error("Error trying to validate card info", err)
		resterr := validation.ValidateCardError(err)
		c.JSON(http.StatusBadRequest, resterr)
		return
	}

	actionDomain := models.NewActionDomain(
		actionRequest.UserId,
	)

	if err := cc.service.MoveCard(cardId, actionDomain); err != nil {
		logger.Error("Error when trying to move card", err)
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"MoveCard controller executed successfully",
		zap.String("journey", "moveCard"))

	// Send Notification

	c.JSON(http.StatusOK, rest_success.BaseRequestReturn{
		Status: "success", Message: "card updated",
	})
}
