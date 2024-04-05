package routes

import (
	"net/http"

	"de.anikate/events-api/middleware"
	"github.com/gin-gonic/gin"
)

func Setup(engine *gin.Engine) {
	engine.GET("/", home)

	engine.GET("/events", getAllEvents)
	engine.GET("/events/:id", getEvent)

	authenticated := engine.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.POST("/events", createEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	engine.POST("/signup", signUp)
	engine.POST("/login", login)

	authenticated.POST("/events/:id/register", Register)
	authenticated.DELETE("/events/:id/register", Unregister)
}

func home(context *gin.Context) {
	context.JSON(
		http.StatusOK,
		gin.H{
			"message":    "Welcome to the Events REST API!",
			"version":    "0.0.1",
			"author":     "Anikate De",
			"author_url": "https://github.com/Anikate-De",
		})
}
