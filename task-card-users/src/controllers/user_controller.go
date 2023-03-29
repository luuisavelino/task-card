package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/task-card-users/src/models/service"
)

type UserControllerInterface interface {
	FindUsers(c *gin.Context)
	FindUserById(c *gin.Context)

	DeleteUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}

func NewUserControllerInterface(serviceInterface service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}
