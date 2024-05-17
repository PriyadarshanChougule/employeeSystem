package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/priyadarshanchougule/g10CRUD/controllers"
	"github.com/priyadarshanchougule/g10CRUD/middlewares"
)

func EmployeeRoutes(router *gin.Engine) {
	allEmployeeRoutes := router.Group("/employee")
	allEmployeeRoutes.Use(middlewares.LoggerMiddleware)

	allEmployeeRoutes.GET("/getEmployee/:employeeCode", controllers.GetEmployee)
	allEmployeeRoutes.POST("addEmployee", controllers.AddEmployee)
	allEmployeeRoutes.DELETE("deleteEmployee/:employeeCode", controllers.DeleteEmployee)
	allEmployeeRoutes.PUT("updateEmployee/:employeeCode", controllers.UpdateEmployee)
}
