package controllers

import (
	"../models"
	"github.com/gin-gonic/gin"
)

type DepartmentController struct{}

var departmentModel = new(models.Department)

func (u DepartmentController) GetAll(c *gin.Context) {
	departments, err := departmentModel.GetAll()

	if err != nil {
		c.JSON(500, gin.H{"message": "Error to retrieve user", "error": err})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": departments})
	return
}
