package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/task-card-cards/pkg/models"
)

const (
	invalidId     = "invalid id"
	invalidUserId = "invalid user id"
)

type requestReturn struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Cards(c *gin.Context) {
	c.Request.ParseMultipartForm(1000)
	userId, err := strconv.Atoi(c.Request.PostForm["user_id"][0])
	if err != nil {
		c.JSON(http.StatusBadRequest, requestReturn{
			"error", invalidUserId,
		})
		return
	}

	cards, err := models.GetCards(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, requestReturn{
			"error", err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, cards)
}

func Card(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, requestReturn{
			"error", invalidId,
		})
		return
	}

	card, err := models.GetCard(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, requestReturn{
			"error", err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, card)
}
