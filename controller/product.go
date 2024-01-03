package controller

import (
	"amazonia/model"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
	"github.com/pkg/errors"
)

func GetProductData(requestURL string) (model.Product, error) {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only Amazon domains
		colly.AllowedDomains("amazon.in", "www.amazon.in", "amazon.com", "www.amazon.com"),

		// Cache responses to prevent multiple download of pages
		// even if the collector is restarted
		colly.CacheDir("./amzn_cache"),
	)

	c.OnRequest(func(r *colly.Request) {
		fixedURL, err := makeIndiaDomain(requestURL)
		if err != nil {
			return
		}
		r.Ctx.Put("url", fixedURL)
		fmt.Println("Fixed link of the page:", fixedURL)
	})

	var product model.Product

	// Find and extract JavaScript content
	c.OnHTML("script", func(e *colly.HTMLElement) {
		// Extract the JavaScript code from the script tag
		jsCode := e.Text
		subPageType := "randomPageType"

		// Check if the JavaScript code contains window.$Nav and config.subPageType
		if strings.Contains(jsCode, "window.$Nav") && strings.Contains(jsCode, "'config.subPageType'") {

			fmt.Println("Fetch subpageType")

			// Use regular expression to extract the value of 'config.subPageType'
			re := regexp.MustCompile(`'config\.subPageType','([^']*)'`)
			matches := re.FindStringSubmatch(jsCode)

			if len(matches) == 2 {
				// The value is captured in the second element of the matches slice
				subPageType = matches[1]
				fmt.Println("SubPageType:", subPageType)
			}
		}

	})

	c.OnHTML("html", func(h *colly.HTMLElement) {
		product = model.Product{
			Title:       h.ChildText("#productTitle"),
			Author:      h.ChildText(".author > a:nth-child(1)"),
			Description: h.ChildText("#bookDescription_feature_div > div:nth-child(1) > div:nth-child(1) > span:nth-child(1)"),
			Rating:      h.ChildText("#averageCustomerReviews > span:nth-child(1) > span:nth-child(1) > span:nth-child(1) > a:nth-child(1) > i:nth-child(2) > span:nth-child(1)"),
			Publisher:   h.ChildText("ul.a-vertical:nth-child(1) > li:nth-child(1) > span:nth-child(1) > span:nth-child(2)"),
			Language:    h.ChildText("ul.a-vertical:nth-child(1) > li:nth-child(2) > span:nth-child(1) > span:nth-child(2)"),
			ISBN10:      h.ChildText("ul.a-vertical:nth-child(1) > li:nth-child(4) > span:nth-child(1) > span:nth-child(2)"),
			ISBN13:      h.ChildText("ul.a-vertical:nth-child(1) > li:nth-child(5) > span:nth-child(1) > span:nth-child(2)"),
			URL:         h.Request.Ctx.Get("url"),
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

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.Ctx.Get("url"))
	})

	c.Visit(requestURL)

	return product, nil
}

func makeIndiaDomain(requestURL string) (string, error) {
	parsedURL, err := url.Parse(requestURL)
	if err != nil {
		return "", errors.Wrapf(err, "Error parsing the provided URL to fix India domain")
	}

	if parsedURL != nil && parsedURL.Host != "amazon.in" {
		parsedURL.Host = "amazon.in"
		return parsedURL.String(), nil
	}

	return requestURL, nil
}
