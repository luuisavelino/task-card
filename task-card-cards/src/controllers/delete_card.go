package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/task-card-cards/src/configuration/logger"
	"github.com/luuisavelino/task-card-cards/src/configuration/rest_success"
	"github.com/luuisavelino/task-card-cards/src/configuration/validation"
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
// @Param user_id formData string true "User id of the card"
// @Success 200 {object} globals.BaseRequestReturn
// @Failure 400 {object} globals.BaseRequestReturn
// @Router /cards/{id} [delete]
func (cc *cardControllerInterface) DeleteCard(c *gin.Context) {
	logger.Info("Init DeleteCard controller",
		zap.String("journey", "deleteCard"),
	)

	cardId, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		logger.Error("Error trying to get card id", err)
		resterr := validation.ValidateUserError(err)
		c.JSON(http.StatusBadRequest, resterr)
		return
	}

	userId := 10
	// c.Request.ParseMultipartForm(1000)
	// userId, err := strconv.Atoi(c.Request.PostForm["user_id"][0])
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, rest_success.BaseRequestReturn{
	// 		Status: "error", Message: invalidUserId,
	// 	})
	// 	return
	// }

	if err := cc.service.DeleteCard(cardId, userId); err != nil {
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
