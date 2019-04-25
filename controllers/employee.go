package controllers

import (
	"log"

	"../models"
	"github.com/gin-gonic/gin"
)

type EmployeeController struct{}

var employeeModel = new(models.Employee)

func (u EmployeeController) GetAll(c *gin.Context) {
	depart := c.Query("department")
	log.Println("DEPART:", depart)
	employees, err := employeeModel.GetAll()

	if err != nil {
		c.JSON(500, gin.H{"message": "Error to retrieve user", "error": err})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": employees})
	return
}

func (u EmployeeController) GetByDepartment(c *gin.Context) {
	departId := c.Param("ID")
	log.Println(departId)
	employee, err := employeeModel.GetByDepartment(departId)

	if err != nil {
		c.JSON(500, gin.H{"message": "Error to retrieve user", "error": err})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": employee})
	return
}

func (u EmployeeController) GetById(c *gin.Context) {
	Id := c.Param("ID")
	log.Println(Id)
	employee, err := employeeModel.Get(Id)

	if err != nil {
		c.JSON(500, gin.H{"message": "Error to retrieve user", "error": err})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": employee})
	return
}
