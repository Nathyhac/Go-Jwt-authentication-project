package routes

import(
  controller "Go-Jwt-authentication/controllers"
  "github.com/gin-gonic/gin"

)

func AuthRoutes(incomingRoutes *gin.Engine)  {
	 incomingRoutes.POST("users/signup", controller.signup())
	 incomingRoutes.POST("users/login", controller.login())
  
}