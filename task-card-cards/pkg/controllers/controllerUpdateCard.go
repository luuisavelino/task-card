package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/task-card-cards/pkg/models"
)

func UpdateCard(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, requestReturn{
			"error", invalidId,
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
				c.JSON(http.StatusBadRequest, requestReturn{
					"error", invalidUserId,
				})
				return
			}
			card.UserId = userId
		}
	}

	card.Id = id
	models.UpdateCardInfo(card)
	c.JSON(http.StatusOK, requestReturn{
		"success", "card updated",
	})
}

func MoveCard(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, requestReturn{
			"error", invalidId,
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

	c.JSON(http.StatusOK, requestReturn{
		"success", "card moved to" + card.CardStatus,
	})
}
