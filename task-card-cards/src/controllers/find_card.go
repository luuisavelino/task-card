package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/task-card-cards/src/configuration/logger"
	"github.com/luuisavelino/task-card-cards/src/view"
	"go.uber.org/zap"
)

const (
	invalidId     = "invalid id"
	invalidUserId = "invalid user id"
)

// @BasePath /api/v1

// Cards godoc
// @Summary Get all cards
// @Description Get all cards
// @Tags cards
// @Accept  json
// @Produce  json
// @Param user_id formData string true "User id of the card"
// @Success 200 {object} []models.Card
// @Failure 400 {object} rest_success.BaseRequestReturn
// @Router /cards [get]
func (cc *cardControllerInterface) FindCards(c *gin.Context) {
	logger.Info("Init FindCards controller",
		zap.String("journey", "findCards"),
	)

	userId := 10
	// c.Request.ParseMultipartForm(1000)
	// userId, err := strconv.Atoi(c.Request.PostForm["user_id"][0])
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, rest_success.BaseRequestReturn{
	// 		Status: "error", Message: invalidUserId,
	// 	})
	// 	return
	// }

	domain, err := cc.service.FindCards(userId)
	if err != nil {
		logger.Error("Error when trying to find users", err)
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"FindCards controller executed successfully",
		zap.String("journey", "findCards"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domain))
}

// @BasePath /api/v1

// Card godoc
// @Summary Get a card
// @Description Get a card
// @Tags cards
// @Accept  json
// @Produce  json
// @Param id path int true "Card id"
// @Success 200 {object} models.Card
// @Failure 400 {object} rest_success.BaseRequestReturn
// @Router /cards/{id} [get]
func (cc *cardControllerInterface) FindCardById(c *gin.Context) {
	logger.Info("Init FindCardById controller",
		zap.String("journey", "findCardById"),
	)

	cardId := 10
	// cardId, err := strconv.Atoi(c.Params.ByName("id"))
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, rest_success.BaseRequestReturn{
	// 		Status: "error", Message: invalidId,
	// 	})
	// 	return
	// }

	userId := 10
	// c.Request.ParseMultipartForm(1000)
	// userId, err := strconv.Atoi(c.Request.PostForm["user_id"][0])
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, rest_success.BaseRequestReturn{
	// 		Status: "error", Message: invalidUserId,
	// 	})
	// 	return
	// }

	domain, err := cc.service.FindCardById(cardId, userId)
	if err != nil {
		logger.Error("Error when trying to find user", err)
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"FindCardById controller executed successfully",
		zap.String("journey", "findCardById"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domain))
}
