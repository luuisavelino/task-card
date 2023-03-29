package service

import (
	"github.com/luuisavelino/task-card-users/src/configuration/database"
	"github.com/luuisavelino/task-card-users/src/configuration/logger"
	"github.com/luuisavelino/task-card-users/src/configuration/rest_err"
	"github.com/luuisavelino/task-card-users/src/models"
	"go.uber.org/zap"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Userpass string `json:"userpass"`
	Email    string `json:"email"`
	RoleId   int    `json:"role_id"`
}

func (u *userDomainService) FindUsers() (map[int]models.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUser model",
		zap.String("journey", "findUser"),
	)

	db := database.ConnectsWithDatabase()

	selectAllUsers, err := db.Query("select * from users order by id asc")
	if err != nil {
		return nil, rest_err.NewForbiddenError("error to get values")
	}

	users := make(map[int]models.UserDomainInterface)
	for selectAllUsers.Next() {
		var id, roleId int
		var username, userpass, email string

		err = selectAllUsers.Scan(&id, &username, &userpass, &email, &roleId)
		if err != nil {
			return nil, rest_err.NewForbiddenError("error to get values")
		}

		users[id] = models.NewUserDomain(
			username,
			userpass,
			email,
			roleId,
		)
	}

	defer db.Close()
	return users, nil
}

func (u *userDomainService) FindUserById(userId int) (map[int]models.UserDomainInterface, *rest_err.RestErr) {
	db := database.ConnectsWithDatabase()

	selectUser, err := db.Query("select * from users where id=?", userId)
	if err != nil {
		logger.Error("Error trying to prepare query", err)
		return nil, rest_err.NewForbiddenError("error to get values")
	}

	user := make(map[int]models.UserDomainInterface)
	for selectUser.Next() {
		var id, roleId int
		var username, userpass, email string

		err = selectUser.Scan(&id, &username, &userpass, &email, &roleId)
		if err != nil {
			logger.Error("Error trying to scan values", err)
			return nil, rest_err.NewForbiddenError("error to get values")
		}

		user[id] = models.NewUserDomain(
			username,
			userpass,
			email,
			roleId,
		)
	}

	defer db.Close()
	return user, nil
}
