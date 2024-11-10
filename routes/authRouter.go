package routes

import (
	controller "jwt_gin-gonic/controller"

	"github.com/gin-gonic/gin"
)

func authRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/signup", controller.Signup())
	incomingRoutes.POST("/login", controller.Login())
}
