package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/task-card-users/pkg/controllers"
)

const apiVersion = "/api/v1"

func HandlerRequest() {
	router := gin.Default()

	// User management
	router.GET(apiVersion+"/users", controllers.Users)
	router.GET(apiVersion+"/users/:id", controllers.User)
	router.POST(apiVersion+"/users", controllers.CreateUser)
	router.PATCH(apiVersion+"/users/:id", controllers.UpdateUser)
	router.DELETE(apiVersion+"/users/:id", controllers.DeleteUser)

	// Card management
	router.GET(apiVersion+"/users/:id/cards", nil)
	router.GET(apiVersion+"/users/:id/cards/:card", nil)
	router.POST(apiVersion+"/users/:id/cards", nil)
	router.PATCH(apiVersion+"/users/:id/cards/:card", nil)
	router.DELETE(apiVersion+"/users/:id/cards/:card", nil)

	router.Run(":8080")
}
