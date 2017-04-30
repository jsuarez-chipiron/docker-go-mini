package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func handleUncontrolledError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func main() {
	port := "8080"

	router := gin.New()
	router.Use(gin.Logger())
	// router.LoadHTMLGlob("templates/*.tmpl.html")
	// router.Static("/static", "static")

	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.tmpl.html", nil)
	// })

	router.GET("/test", func(c *gin.Context) {
		// c.String(http.StatusOK, telnet())
		c.String(http.StatusOK, "It works")
	})

	// router.GET("/repeat", repeatFunc)
	// router.GET("/db", dbFunc)

	router.Run(":" + port)
}
