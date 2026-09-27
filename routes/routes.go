package routes

import (
	"Go-gin-rest-api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	authenticate := server.Group("/")
	authenticate.Use(middleware.Authenticate)
	authenticate.POST("/events", createEvent)
	authenticate.PUT("/events/:id", updateEventById)
	authenticate.DELETE("/events/:id", deleteEvent)
	authenticate.POST("/events/:id/register", registerForEvent)
	authenticate.DELETE("/events/:id/register", cancelRegistration)

	server.GET("/events/:id", getEventById)
	server.GET("/events", getEvents)
	server.POST("/signUp", signUp)
	server.POST("/login", login)
}
