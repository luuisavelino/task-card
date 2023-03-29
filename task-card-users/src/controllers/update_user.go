package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/task-card-users/src/configuration/logger"
	"github.com/luuisavelino/task-card-users/src/configuration/rest_success"
	"github.com/luuisavelino/task-card-users/src/configuration/validation"
	"github.com/luuisavelino/task-card-users/src/controllers/model/request"
	"github.com/luuisavelino/task-card-users/src/models"
	"go.uber.org/zap"
)

// @BasePath /api/v1

// UpdateUser godoc
// @Summary Update a user
// @Description Route to update a user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User id"
// @Param user body request.UserRequest true "User info"
// @Success 200 {object} rest_success.BaseRequestReturn
// @Failure 400 {object} rest_success.BaseRequestReturn
// @Router /users/{id} [put]
func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init UpdateUser controller",
		zap.String("journey", "updateUser"),
	)

	userId, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		logger.Error("Error trying to get user id", err)
		resterr := validation.ValidateUserError(err)
		c.JSON(http.StatusBadRequest, resterr)
		return
	}

	var userRequest request.UserRequest

	if err := c.ShouldBind(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err)
		resterr := validation.ValidateUserError(err)
		c.JSON(http.StatusBadRequest, resterr)
		return
	}

	domain := models.NewUserDomain(
		userRequest.Username,
		userRequest.Userpass,
		userRequest.Email,
		userRequest.RoleId,
	)

	if err := uc.service.UpdateUser(userId, domain); err != nil {
		logger.Error("Error when trying to create user", err)
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"UpdateUser controller executed successfully",
		zap.String("journey", "updateUser"))

	c.JSON(http.StatusOK, rest_success.BaseRequestReturn{
		Status: "success", Message: "user updated",
	})
}
