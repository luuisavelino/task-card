package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/task-card-cards/globals"
	"github.com/luuisavelino/task-card-cards/pkg/models/request"
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
func CreateCard(c *gin.Context) {
	var cardRequest request.CardRequest

	if err := c.ShouldBind(&cardRequest); err != nil {
		resterr := globals.NewBadRequestError(
			fmt.Sprint("there are some incrorrect filds, error=" + err.Error()),
		)
		c.JSON(http.StatusBadRequest, resterr)
		return
	}

	fmt.Println(cardRequest)

	// if err := helpers.ValidateCard(cardRequest); err != nil {
	// 	c.JSON(http.StatusBadRequest, globals.BaseRequestReturn{
	// 		Status: "error", Message: err.Error(),
	// 	})
	// 	return
	// }

	// if err := models.CreateNewCard(cardRequest); err != nil {
	// 	c.JSON(http.StatusBadRequest, globals.BaseRequestReturn{
	// 		Status: "error", Message: err.Error(),
	// 	})
	// 	return
	// }

	c.JSON(http.StatusOK, globals.BaseRequestReturn{
		Status: "success", Message: "card created",
	})
}
