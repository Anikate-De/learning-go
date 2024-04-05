package main

import (
	"de.anikate/events-api/db"
	"de.anikate/events-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	engine := gin.Default()

	routes.Setup(engine)

	engine.Run(":8080")
}
