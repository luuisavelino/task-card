package service

import (
	"github.com/luuisavelino/task-card-users/src/configuration/rest_err"
	"github.com/luuisavelino/task-card-users/src/models"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	FindUsers() (map[int]models.UserDomainInterface, *rest_err.RestErr)
	FindUserById(int) (map[int]models.UserDomainInterface, *rest_err.RestErr)

	CreateUser(models.UserDomainInterface) *rest_err.RestErr
	UpdateUser(int, models.UserDomainInterface) *rest_err.RestErr
	DeleteUser(int) *rest_err.RestErr
}
