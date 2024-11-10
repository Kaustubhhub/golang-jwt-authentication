package routes

import (
	controller "jwt_gin-gonic/controller"
	"jwt_gin-gonic/middleware"

	"github.com/gin-gonic/gin"
)

func userRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:userId", controller.GetUser())
}
