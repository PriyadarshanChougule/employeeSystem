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

	//register routes
	routes.EmployeeRoutes(r)

	//run server and start listening
	r.Run()
}
