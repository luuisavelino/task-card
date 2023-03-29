package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/luuisavelino/task-card-cards/src/configuration/logger"
	"github.com/luuisavelino/task-card-cards/src/configuration/rest_success"
	"github.com/luuisavelino/task-card-cards/src/configuration/validation"
	"github.com/luuisavelino/task-card-cards/src/controllers/model/request"
	"github.com/luuisavelino/task-card-cards/src/models"
)

// @BasePath /api/v1

// CreateCard godoc
// @Summary Create a card
// @Description Route to update a card
// @Tags cards
// @Accept  json
// @Produce  json
// @Param title formData string true "Title of the card"
// @Param summary formData string true "Summary of the card"
// @Param due_date formData string true "Due date of the card"
// @Param card_status formData string true "Status of the card"
// @Param user_id formData string true "User id of the card"
// @Success 200 {object} globals.BaseRequestReturn
// @Failure 400 {object} globals.BaseRequestReturn
// @Router /cards [post]
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
