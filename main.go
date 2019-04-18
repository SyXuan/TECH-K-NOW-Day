package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Start web service")

	r := gin.Default()

	r.Static("/assets", "./assets")

	r.LoadHTMLGlob("templates/*")

	r.GET("/", index)

	r.POST("/login", login)

	r.Run(":8080")
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html.tmpl", gin.H{
		"Title": "Hello TECH(K)NOW Day",
	})
}

func login(c *gin.Context) {
	username := c.PostForm("username")
	log.Println("username:", username)
	c.String(http.StatusOK, "Hi "+username)
}
