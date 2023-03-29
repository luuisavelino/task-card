package view

import (
	"github.com/luuisavelino/task-card-users/src/controllers/model/response"
	"github.com/luuisavelino/task-card-users/src/models"
)

func ConvertDomainToResponse(usernameDomains map[int]models.UserDomainInterface) []response.UserResponse {
	var users []response.UserResponse
	for id, userDomain := range usernameDomains {
		users = append(users, response.UserResponse{
			Id:       id,
			Username: userDomain.GetUsername(),
			Email:    userDomain.GetEmail(),
			RoleId:   userDomain.GetRoleId(),
		})
	}

	return users
}
