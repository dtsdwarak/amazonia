package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "you")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome " + firstname + " " + lastname + " to Amazonia - Amazon product exploration API",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
