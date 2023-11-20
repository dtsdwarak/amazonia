package main

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type Product struct {
	Title       string
	Description string
	Author      string
	Rating      string
	Images      []string
	URL         string
	Publisher   string
	Language    string
	ISBN10      string
	ISBN13      string
}

func getProductData(requestURL string) (Product, error) {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: amazon.in, www.amazon.in
		colly.AllowedDomains("amazon.in", "www.amazon.in"),

		// Cache responses to prevent multiple download of pages
		// even if the collector is restarted
		colly.CacheDir("./amzn_cache"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Link of the page:", requestURL)
	})

	var product Product

	c.OnHTML("html", func(h *colly.HTMLElement) {
		product = Product{
			Title:       h.ChildText("#productTitle"),
			Author:      h.ChildText(".author > a:nth-child(1)"),
			Description: h.ChildText("#bookDescription_feature_div > div:nth-child(1) > div:nth-child(1) > span:nth-child(1)"),
			Rating:      h.ChildText("#averageCustomerReviews > span:nth-child(1) > span:nth-child(1) > span:nth-child(1) > a:nth-child(1) > i:nth-child(2) > span:nth-child(1)"),
			Publisher:   h.ChildText("ul.a-vertical:nth-child(1) > li:nth-child(1) > span:nth-child(1) > span:nth-child(2)"),
			Language:    h.ChildText("ul.a-vertical:nth-child(1) > li:nth-child(2) > span:nth-child(1) > span:nth-child(2)"),
			ISBN10:      h.ChildText("ul.a-vertical:nth-child(1) > li:nth-child(4) > span:nth-child(1) > span:nth-child(2)"),
			ISBN13:      h.ChildText("ul.a-vertical:nth-child(1) > li:nth-child(5) > span:nth-child(1) > span:nth-child(2)"),
			URL:         requestURL,
		}
	})

	c.OnHTML("#imgTagWrapperId", func(e *colly.HTMLElement) {
		innerHtml, err := e.DOM.Html()
		if err != nil {
			fmt.Println("Error")
		}
		r := regexp.MustCompile(`https?://[^\s]+?\.jpg`)
		product.Images = r.FindAllString(innerHtml, -1)

	})

	c.Visit(requestURL)

	return product, nil
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		requestURL := c.DefaultQuery("requestURL", "https://www.amazon.in/dp/9354353010/")
		fmt.Println("input - " + requestURL)
		product, _ := getProductData(requestURL)
		// lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
		c.JSON(http.StatusOK, product)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
