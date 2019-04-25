package server

import (
	"github.com/gin-contrib/static"

	"../config"
)

//
func Init() {
	c := config.GetConfig()

	r := NewRouter()
	r.Use(static.Serve("/", static.LocalFile("./www", false)))
	r.Run(c.GetString("app.port"))
}
