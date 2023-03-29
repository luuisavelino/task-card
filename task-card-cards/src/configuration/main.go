package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/task-card-users/src/controllers"
	"github.com/luuisavelino/task-card-users/src/controllers/routes"
	"github.com/luuisavelino/task-card-users/src/models/service"
)

func main() {
	service := service.NewUserDomainService()
	userController := controllers.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
