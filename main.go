package main

import (
	"github.com/gin-gonic/gin"
	"github.com/priyadarshanchougule/g10CRUD/db"
	"github.com/priyadarshanchougule/g10CRUD/routes"
)

func main() {
	db.GetAllCollections()

	//create server on 8080 default port
	r := gin.Default()

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	//register routes
	routes.EmployeeRoutes(r)
	routes.LoginRoutes(r)
	//run server and start listening
	r.Run()
}
