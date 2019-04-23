package server

import (
	"../controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	//router.Use(middlewares.AuthMiddleware())

	v1 := router.Group("api/v1")
	{

		userController := new(controllers.UserController)
		departmentController := new(controllers.DepartmentController)

		v1.GET("/users", userController.GetAll)
		v1.GET("/departments", departmentController.GetAll)

	}

	return router
}
