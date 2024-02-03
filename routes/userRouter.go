package routes

import (
	controller "Go-Jwt-authentication/controllers"

	"github.com/gin-gonic/gin"
	"golang-jwt-project/middleware"
)

  func UserRoutes(incomingRoutes *gin.Engine)  {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUser())
	incomingRoutes.GET("users/user_id", controller.GetUser())
	
  }