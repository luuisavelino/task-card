package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/task-card-cards/globals"
	"github.com/luuisavelino/task-card-cards/pkg/models"
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
// @Success 200 {object} globals.BaseRequestReturn
// @Failure 400 {object} globals.BaseRequestReturn
// @Router /cards/{id} [patch]
func UpdateCard(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, globals.BaseRequestReturn{
			Status: "error", Message: invalidId,
		})
		return
	}

	var card models.Card

	c.Request.ParseMultipartForm(1000)
	for key, value := range c.Request.PostForm {
		switch key {
		case "title":
			card.Title = value[0]
		case "summary":
			card.Summary = value[0]
		case "due_date":
			card.DueDate = value[0]
		case "card_status":
			card.CardStatus = value[0]
		case "user_id":
			userId, err := strconv.Atoi(c.Request.PostForm["user_id"][0])
			if err != nil {
				c.JSON(http.StatusBadRequest, globals.BaseRequestReturn{
					Status: "error", Message: invalidUserId,
				})
				return
			}
			card.UserId = userId
		}
	}

	card.Id = id
	models.UpdateCardInfo(card)
	c.JSON(http.StatusOK, globals.BaseRequestReturn{
		Status: "success", Message: "card updated",
	})
}

// @BasePath /api/v1

// MoveCard godoc
// @Summary Move a card
// @Description Move a card
// @Tags cards
// @Accept  json
// @Produce  json
// @Param id path int true "Card id"
// @Param card_status formData string true "Status of the card"
// @Success 200 {object} globals.BaseRequestReturn
// @Failure 400 {object} globals.BaseRequestReturn
// @Router /cards/{id} [put]
func MoveCard(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, globals.BaseRequestReturn{
			Status: "error", Message: invalidId,
		})
		return
	}

	var card models.Card
	c.Request.ParseMultipartForm(1000)
	for key, value := range c.Request.PostForm {
		switch key {
		case "card_status":
			card.CardStatus = value[0]
		}
	}

	card.Id = id

	models.UpdateCardStatus(card)
	models.SendNotification(id, "update")

	c.JSON(http.StatusOK, globals.BaseRequestReturn{
		Status: "success", Message: "card moved to" + card.CardStatus,
	})
}
