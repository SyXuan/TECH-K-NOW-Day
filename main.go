package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	melody "gopkg.in/olahol/melody.v1"
)

var title = "Hello TECH(K)NOW Day"

func main() {
	log.Println("Start web service")

	r := gin.Default()
	m := melody.New()

	r.Static("/assets", "./assets")

	r.LoadHTMLGlob("templates/*")

	r.GET("/", index)

	r.GET("/login", loginPage)

	r.POST("/login", login)

	r.GET("/chat", chat)

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})

	r.Run(":8080")
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html.tmpl", gin.H{
		"Title": title,
	})
}

func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	log.Println("username:", username)
	if username == "Benjamin" && password == "1234" {
		c.SetCookie("username", username, 60, "/", "localhost", false, true)
		c.String(http.StatusOK, "Hi "+username)
	} else {
		c.String(http.StatusOK, fmt.Sprintf("%s password error", username))
	}
}

func loginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html.tmpl", gin.H{
		"Title": title,
	})
}

func chat(c *gin.Context) {
	username, _ := c.Cookie("username")
	c.HTML(http.StatusOK, "chat.html.tmpl", gin.H{
		"username": username,
	})
}
