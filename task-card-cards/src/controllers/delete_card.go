package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/task-card-cards/src/configuration/rest_success"
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
func DeleteCard(c *gin.Context) {
	// id, err := strconv.Atoi(c.Params.ByName("id"))
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, rest_success.BaseRequestReturn{
	// 		Status: "error", Message: invalidId,
	// 	})
	// 	return
	// }

	// c.Request.ParseMultipartForm(1000)
	// userId, err := strconv.Atoi(c.Request.PostForm["user_id"][0])
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, rest_success.BaseRequestReturn{
	// 		Status: "error", Message: invalidUserId,
	// 	})
	// 	return
	// }

	// if err = models.DeleteCard(id, userId); err != nil {
	// 	c.JSON(http.StatusBadRequest, rest_success.BaseRequestReturn{
	// 		Status: "error", Message: err.Error(),
	// 	})
	// 	return
	// }

	c.JSON(http.StatusOK, rest_success.BaseRequestReturn{
		Status: "success", Message: "card deleted",
	})
}
