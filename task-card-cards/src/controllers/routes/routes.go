package routes

import (
	"github.com/gin-gonic/gin"

	docs "github.com/luuisavelino/task-card-cards/docs"
	"github.com/luuisavelino/task-card-cards/src/controllers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const apiVersion = "/api/v1"

func InitRoutes(r *gin.RouterGroup, cardController controllers.CardControllerInterface) {
	docs.SwaggerInfo.BasePath = apiVersion

	r.GET(apiVersion+"/cards", cardController.FindCards)
	r.GET(apiVersion+"/cards/:id", cardController.FindCardById)
	r.POST(apiVersion+"/cards", cardController.CreateCard)
	r.PATCH(apiVersion+"/cards/:id", cardController.UpdateCard)
	r.DELETE(apiVersion+"/cards/:id", cardController.DeleteCard)
	r.POST(apiVersion+"/cards/:id", cardController.MoveCard)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
