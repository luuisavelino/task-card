package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/task-card-cards/pkg/models"
)

func CreateCard(c *gin.Context) {
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
			userId, err := strconv.Atoi(value[0])
			if err != nil {
				c.JSON(http.StatusBadRequest, requestReturn{
					"error", invalidUserId,
				})
				return
			}
			card.UserId = userId
		}
	}

	if err := models.CreateNewCard(card); err != nil {
		c.JSON(http.StatusBadRequest, requestReturn{
			"error", err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, requestReturn{
		"success", "card created",
	})
}
