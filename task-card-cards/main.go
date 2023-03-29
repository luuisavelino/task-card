package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/task-card-cards/src/controllers"
	"github.com/luuisavelino/task-card-cards/src/controllers/routes"
	"github.com/luuisavelino/task-card-cards/src/models/service"
)

func main() {
	service := service.NewCardDomainService()
	userController := controllers.NewCardControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
