package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/task-card-cards/pkg/controllers"
)

const apiVersion = "/api/v1"

func HandlerRequest() {
	router := gin.Default()

	// Card management
	router.GET(apiVersion+"/cards", controllers.Cards)
	router.GET(apiVersion+"/cards/:id", controllers.Card)
	router.POST(apiVersion+"/cards", controllers.CreateCard)
	router.PATCH(apiVersion+"/cards/:id", controllers.UpdateCard)
	router.DELETE(apiVersion+"/cards/:id", controllers.DeleteCard)
	router.POST(apiVersion+"/cards/:id", controllers.MoveCard)

	router.Run(":8080")
}
