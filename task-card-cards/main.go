package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/luuisavelino/task-card-cards/src/controllers/routes"
)

func main() {
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
