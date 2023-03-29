package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/task-card-users/src/configuration/logger"
	"github.com/luuisavelino/task-card-users/src/configuration/validation"
	"github.com/luuisavelino/task-card-users/src/view"
	"go.uber.org/zap"
)

// @BasePath /api/v1

// FindUsers godoc
// @Summary Get all users
// @Description Route to get all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} []response.UserResponse
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /users [get]
func (uc *userControllerInterface) FindUsers(c *gin.Context) {
	logger.Info("Init FindUsers controller",
		zap.String("journey", "findUsers"),
	)

	domain, err := uc.service.FindUsers()
	if err != nil {
		logger.Error("Error when trying to find users", err)
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"FindUsers controller executed successfully",
		zap.String("journey", "findUsers"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domain))
}

// @BasePath /api/v1

// FindUserById godoc
// @Summary Get a user
// @Description Route to get a user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User id"
// @Success 200 {object} response.UserResponse
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /users/{id} [get]
func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	logger.Info("Init FindUserById controller",
		zap.String("journey", "findUsersById"),
	)

	userId, errToConvert := strconv.Atoi(c.Params.ByName("id"))
	if errToConvert != nil {
		logger.Error("Error trying to get user id", errToConvert)
		resterr := validation.ValidateUserError(errToConvert)
		c.JSON(http.StatusBadRequest, resterr)
		return
	}

	domain, err := uc.service.FindUserById(userId)
	if err != nil {
		logger.Error("Error when trying to find user", err)
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"FindUserById controller executed successfully",
		zap.String("journey", "findUsersById"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domain))
}
