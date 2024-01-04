package main

import (
	"amazonia/controller"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/product", func(c *gin.Context) {
		requestURL := c.DefaultQuery("requestURL", "https://www.amazon.in/dp/9354353010/")
		fmt.Println("Input - " + requestURL)
		product, _ := controller.GetProductData(requestURL)
		c.JSON(http.StatusOK, product)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
