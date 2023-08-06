package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
  r.Static("/static", "./static/css")
	r.LoadHTMLGlob("static/**/*")
	r.GET("/hello", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home", gin.H{
			"Title": "Main website",
			"IP":    c.RemoteIP(),
		})
	})
	r.Run(":8080")
}
