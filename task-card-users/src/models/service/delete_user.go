package service

import (
	"fmt"

	"github.com/luuisavelino/task-card-users/src/configuration/database"
	"github.com/luuisavelino/task-card-users/src/configuration/logger"
	"github.com/luuisavelino/task-card-users/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (u *userDomainService) DeleteUser(id int) *rest_err.RestErr {
	logger.Info("Init deleteUser model",
		zap.String("journey", "deleteUser"),
	)

	db := database.ConnectsWithDatabase()

	deleteUserById, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		logger.Error("Error trying to prepare query", err)
		return rest_err.NewForbiddenError("error to delete values")
	}

	fmt.Println(id)

	deleteUserById.Exec(id)
	defer db.Close()
	return nil
}
