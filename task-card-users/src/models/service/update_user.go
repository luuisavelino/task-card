package service

import (
	"github.com/luuisavelino/task-card-users/src/configuration/database"
	"github.com/luuisavelino/task-card-users/src/configuration/logger"
	"github.com/luuisavelino/task-card-users/src/configuration/rest_err"
	"github.com/luuisavelino/task-card-users/src/models"
	"go.uber.org/zap"
)

func (u *userDomainService) UpdateUser(userId int, userDomain models.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init updateUser model",
		zap.String("journey", "updateUser"),
	)

	db := database.ConnectsWithDatabase()

	updateUser, err := db.Prepare("Update users set username=?, userpass=?, email=?, role_id=? where id=?")
	if err != nil {
		logger.Error("Error trying to prepare query", err)
		return rest_err.NewForbiddenError("error to get values")
	}

	updateUser.Exec(userDomain.GetUsername(), userDomain.GetUserpass(), userDomain.GetEmail(), userDomain.GetRoleId(), userId)
	defer db.Close()

	return nil
}
