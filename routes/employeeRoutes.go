package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/priyadarshanchougule/g10CRUD/controllers"
)

func EmployeeRoutes(router *gin.Engine) {

	allEmployeeRoutes := router.Group("/employee")
	allEmployeeRoutes.GET("/getEmployee/:employeeCode", controllers.GetEmployee)

}
