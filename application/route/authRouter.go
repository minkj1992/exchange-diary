package route

import (
	"github.com/ExchangeDiary/exchange-diary/application/controller"
	"github.com/gin-gonic/gin"
)

// AuthRoutes ...
func AuthRoutes(router *gin.RouterGroup, controller controller.AuthController) {
	redirectLogin := router.Group("/login")
	{
		redirectLogin.GET("/:auth_type", controller.Redirect())
	}
	auth := router.Group("/authentication")
	{
		auth.GET("/login/:auth_type", controller.Login())
		auth.GET("/authenticated", controller.Authenticate())
		auth.POST("/mock", controller.MockRegister())
	}
}
