package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEmployee(c *gin.Context) {

	employeeCode := c.Param("employeeCode")

	c.JSON(http.StatusOK, gin.H{"employeeCode": employeeCode})

}
