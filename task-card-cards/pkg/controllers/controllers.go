package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/task-card-cards/pkg/models"
)

const (
	invalidId = "invalid id"
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

func DeleteCard(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, requestReturn{
			"error", invalidId,
		})
		return
	}

	c.Request.ParseMultipartForm(1000)
	userId, err := strconv.Atoi(c.Request.PostForm["user_id"][0])
	if err != nil {
		c.JSON(http.StatusBadRequest, requestReturn{
			"error", invalidUserId,
		})
		return
	}

	if err = models.DeleteCard(id, userId); err != nil {
		c.JSON(http.StatusBadRequest, requestReturn{
			"error", err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, requestReturn{
		"success", "card deleted",
	})
}

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
		}
	}

	card.Id = id
	models.UpdateCard(card)

	c.JSON(http.StatusOK, requestReturn{
		"success", "card updated",
	})
}
