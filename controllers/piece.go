package controllers

import (
	"../models"
	"github.com/gin-gonic/gin"
)

type PieceController struct{}

var pieceModel = new(models.Piece)

func (u PieceController) GetAll(c *gin.Context) {
	// depart := c.Query("department")
	// log.Println("DEPART:", depart)
	news, err := pieceModel.GetAll()

	if err != nil {
		c.JSON(500, gin.H{"message": "Error to retrieve user", "error": err})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": news})
	return
}

/*
func (u EmployeeController) GetById(c *gin.Context) {
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
*/
