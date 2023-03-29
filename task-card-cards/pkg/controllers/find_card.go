package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/task-card-cards/globals"
	"github.com/luuisavelino/task-card-cards/pkg/models"
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
// @Failure 400 {object} globals.BaseRequestReturn
// @Router /cards [get]
func Cards(c *gin.Context) {
	c.Request.ParseMultipartForm(1000)
	userId, err := strconv.Atoi(c.Request.PostForm["user_id"][0])
	if err != nil {
		c.JSON(http.StatusBadRequest, globals.BaseRequestReturn{
			Status: "error", Message: invalidUserId,
		})
		return
	}

	cards, err := models.GetCards(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, globals.BaseRequestReturn{
			Status: "error", Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, cards)
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
// @Failure 400 {object} globals.BaseRequestReturn
// @Router /cards/{id} [get]
func Card(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, globals.BaseRequestReturn{
			Status: "error", Message: invalidId,
		})
		return
	}

	card, err := models.GetCard(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, globals.BaseRequestReturn{
			Status: "error", Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, card)
}
