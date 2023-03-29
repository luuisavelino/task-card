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

// @BasePath /api/v1

// DeleteCard godoc
// @Summary Delete a card
// @Description Delete a card
// @Tags cards
// @Accept  json
// @Produce  json
// @Param id path int true "Card id"
// @Param action body request.ActionRequest true "Action info"
// @Success 200 {object} rest_success.BaseRequestReturn
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /cards/{id} [delete]
func (cc *cardControllerInterface) DeleteCard(c *gin.Context) {
	logger.Info("Init DeleteCard controller",
		zap.String("journey", "deleteCard"),
	)

	cardId, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		logger.Error("Error trying to get card id", err)
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

	if err := cc.service.DeleteCard(cardId, actionDomain); err != nil {
		logger.Error("Error when trying to delete card", err)
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"DeleteCard controller executed successfully",
		zap.String("journey", "deleteCard"))

	c.JSON(http.StatusOK, rest_success.BaseRequestReturn{
		Status: "success", Message: "card deleted",
	})
}
