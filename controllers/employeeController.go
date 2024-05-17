package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/priyadarshanchougule/g10CRUD/db"
	"github.com/priyadarshanchougule/g10CRUD/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetEmployee(c *gin.Context) {
	employeeCode := c.Param("employeeCode")
	employees := db.GetAllCollections().Employees
	var FoundEmp models.Employee
	employees.FindOne(context.TODO(), bson.M{"employeeCode": employeeCode}).Decode(&FoundEmp)
	c.JSON(http.StatusOK, FoundEmp)
}

func AddEmployee(c *gin.Context) {
	fmt.Println("inside route")
	var newEmp models.Employee
	empCollection := db.GetAllCollections().Employees
	if err := c.BindJSON(&newEmp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := empCollection.InsertOne(context.TODO(), newEmp)
	if err != nil {
		log.Panic(err)
		return
	}

	c.JSON(http.StatusOK, res)

	reflect.TypeOf(res)
}

func DeleteEmployee(c *gin.Context) {
	employeCode := c.Param("employeeCode")
	empCollection := db.GetAllCollections().Employees
	filter := bson.M{"employeeCode": employeCode}
	res, _ := empCollection.DeleteOne(context.TODO(), filter)
	c.JSON(http.StatusOK, res)
}

func UpdateEmployee(c *gin.Context) {
	employeeCode := c.Param("employeeCode")
	// body, _ := io.ReadAll(c.Request.Body)
	// fmt.Println("body-->", string(body))
	empCollection := db.GetAllCollections().Employees
	update := bson.D{
		{"$set", bson.D{
			{"name", "Darshan Chougule"},
		},
		},
	}
	var x map[string]interface{}
	empCollection.FindOneAndUpdate(context.TODO(),
		bson.M{"employeeCode": employeeCode}, update).Decode(&x)

	fmt.Println("decoded", x)
}
