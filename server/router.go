package server

import (
	"log"

	"../controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v\n", httpMethod, absolutePath, nuHandlers)
	}
	/*
		router := gin.New()
		router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
			// your custom format
			return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
			)
		}))
	*/
	router.Use(gin.Recovery())

	// router.Use(gin.Logger())
	// router.Use(gin.Recovery())
	// router.Use(middlewares.AuthMiddleware())
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	v1 := router.Group("api/v1")
	{
		userController := new(controllers.UserController)
		pieceController := new(controllers.PieceController)
		assetsController := new(controllers.AssetsController)
		departmentController := new(controllers.DepartmentController)
		employeeController := new(controllers.EmployeeController)

		v1.GET("/users", userController.GetAll)
		v1.GET("/banners", assetsController.GetMainBanners)
		v1.GET("/actions", assetsController.GetMainActions)
		v1.GET("/departments", departmentController.GetAll)
		v1.GET("/news", pieceController.GetAll)
		v1.GET("/employees", employeeController.GetAll)
		v1.GET("/employees/:ID", employeeController.GetById)
		v1.GET("/employeebydept/:ID", employeeController.GetByDepartment)
	}

	return router
}
