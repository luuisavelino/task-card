package service

import (
	"github.com/luuisavelino/task-card-users/src/configuration/database"
	"github.com/luuisavelino/task-card-users/src/configuration/logger"
	"github.com/luuisavelino/task-card-users/src/configuration/rest_err"
	"github.com/luuisavelino/task-card-users/src/models"
	"go.uber.org/zap"
)

func (u *userDomainService) CreateUser(userDomain models.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init createUser model",
		zap.String("journey", "createUser"),
	)

	userDomain.EncryptPassword()

	db := database.ConnectsWithDatabase()

	insertUserIntoDatabase, err := db.Prepare("insert into users(username, userpass, email, role_id) values(?, ?, ?, ?)")
	if err != nil {
		logger.Error("Error trying to prepare query", err)
		return rest_err.NewForbiddenError("error to insert values")
	}

	insertUserIntoDatabase.Exec(userDomain.GetUsername(), userDomain.GetUserpass(), userDomain.GetEmail(), userDomain.GetRoleId())
	defer db.Close()
	return nil
}
