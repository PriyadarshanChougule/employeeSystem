package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/priyadarshanchougule/g10CRUD/controllers"
)

func LoginRoutes(router *gin.Engine) {
	login := router.Group("/login")
	login.POST("/verifyLogin", controllers.VerifyLogin)
}
