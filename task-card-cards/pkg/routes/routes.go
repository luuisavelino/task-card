package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/task-card-cards/pkg/controllers"

	docs "github.com/luuisavelino/task-card-cards/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const apiVersion = "/api/v1"

func InitRoutes(r *gin.RouterGroup) {
	docs.SwaggerInfo.BasePath = apiVersion

	r.GET(apiVersion+"/cards", controllers.Cards)
	r.GET(apiVersion+"/cards/:id", controllers.Card)
	r.POST(apiVersion+"/cards", controllers.CreateCard)
	r.PATCH(apiVersion+"/cards/:id", controllers.UpdateCard)
	r.DELETE(apiVersion+"/cards/:id", controllers.DeleteCard)
	r.POST(apiVersion+"/cards/:id", controllers.MoveCard)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
