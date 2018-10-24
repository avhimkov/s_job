package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	// Set Gin to production mode
	//gin.SetMode(gin.ReleaseMode)

	g := gin.Default()

	g.Static("/web", "./web")
	g.Static("/node_modules", "./node_modules")
	g.LoadHTMLGlob("templates/*.html")

	// g := gin.New()

	// Logging middleware
	g.Use(gin.Logger())
	// Recovery middleware
	g.Use(gin.Recovery())
	// g.Use(favicon.New("./assets/favicon.ico"))
	// g.Use(static.Serve("/web", static.LocalFile("/web", false)))
	// v1 := router.Group("api/v1")
	// {
	// 	v1.GET("/instructions", GetInstructions)
	// }

	return g
}

func main() {
	fmt.Println("add")
}
