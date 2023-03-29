package routes

import (
	"github.com/gin-gonic/gin"

	docs "github.com/luuisavelino/task-card-users/docs"
	"github.com/luuisavelino/task-card-users/src/controllers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const apiVersion = "/api/v1"

func InitRoutes(r *gin.RouterGroup, userController controllers.UserControllerInterface) {
	docs.SwaggerInfo.BasePath = apiVersion

	r.GET(apiVersion+"/users", userController.FindUsers)
	r.GET(apiVersion+"/users/:id", userController.FindUserById)
	r.POST(apiVersion+"/users", userController.CreateUser)
	r.PUT(apiVersion+"/users/:id", userController.UpdateUser)
	r.DELETE(apiVersion+"/users/:id", userController.DeleteUser)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
