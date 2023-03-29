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

// UpdateCard godoc
// @Summary Update a card
// @Description Update a card
// @Tags cards
// @Accept  json
// @Produce  json
// @Param id path int true "Card id"
// @Param title formData string true "Title of the card"
// @Param summary formData string true "Summary of the card"
// @Param due_date formData string true "Due date of the card"
// @Param card_status formData string true "Status of the card"
// @Param user_id formData string true "User id of the card"
// @Success 200 {object} rest_success.BaseRequestReturn
// @Failure 400 {object} rest_success.BaseRequestReturn
// @Router /cards/{id} [patch]
func (cc *cardControllerInterface) UpdateCard(c *gin.Context) {
	logger.Info("Init UpdateCard controller",
		zap.String("journey", "updateCard"),
	)

	cardId, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		logger.Error("Error trying to get card id", err)
		resterr := validation.ValidateUserError(err)
		c.JSON(http.StatusBadRequest, resterr)
		return
	}

	var cardRequest request.CardRequest

	if err := c.ShouldBind(&cardRequest); err != nil {
		logger.Error("Error trying to validate card info", err)
		resterr := validation.ValidateUserError(err)
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

	if err := cc.service.UpdateCardInfo(cardId, domain); err != nil {
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
