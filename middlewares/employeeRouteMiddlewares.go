package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware(c *gin.Context) {
	fmt.Println("inside middleware")
	c.JSON(http.StatusOK, gin.H{"message": "returned from middleware"})
	// c.Abort()
	// c.AbortWithStatus(401)
	c.Next()
}
