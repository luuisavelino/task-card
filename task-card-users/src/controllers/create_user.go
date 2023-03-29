package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/task-card-users/src/configuration/logger"
	"github.com/luuisavelino/task-card-users/src/configuration/rest_success"
	"github.com/luuisavelino/task-card-users/src/configuration/validation"
	"github.com/luuisavelino/task-card-users/src/controllers/model/request"
	"github.com/luuisavelino/task-card-users/src/models"
	"go.uber.org/zap"
)

var (
	UserDomainInterface models.UserDomainInterface
)

// @BasePath /api/v1

// CreateUser godoc
// @Summary 	Create new user
// @Description Route to create a new user
// @Tags 		users
// @Accept 		json
// @Produce 	json
// @Param 		user body request.UserRequest true "User info"
// @Success 200 {object} globals.BaseRequestReturn
// @Failure 400 {object} globals.BaseRequestReturn
// @Failure 500 {object} globals.BaseRequestReturn
// @Router /users [post]
func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)

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

	if err := uc.service.CreateUser(domain); err != nil {
		logger.Error("Error when trying to create user", err)
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"CreateUser controller executed successfully",
		zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, rest_success.BaseRequestReturn{
		Status: "success", Message: "user created",
	})
}
