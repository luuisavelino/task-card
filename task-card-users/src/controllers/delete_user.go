package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/task-card-users/src/configuration/logger"
	"github.com/luuisavelino/task-card-users/src/configuration/rest_success"
	"github.com/luuisavelino/task-card-users/src/configuration/validation"
	"go.uber.org/zap"
)

// @BasePath /api/v1

// DeleteUser godoc
// @Summary Delete a user
// @Description Route to delete a user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User id"
// @Success 200 {object} globals.BaseRequestReturn
// @Failure 400 {object} globals.BaseRequestReturn
// @Router /users/{id} [delete]
func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init DeleteUser controller",
		zap.String("journey", "deleteUser"),
	)

	userId, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		logger.Error("Error trying to get user id", err)
		resterr := validation.ValidateUserError(err)
		c.JSON(http.StatusBadRequest, resterr)
		return
	}

	if err := uc.service.DeleteUser(userId); err != nil {
		logger.Error("Error when trying to delete user", err)
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"DeleteUser controller executed successfully",
		zap.String("journey", "deleteUser"))

	c.JSON(http.StatusOK, rest_success.BaseRequestReturn{
		Status: "success", Message: "user deleted",
	})
}
