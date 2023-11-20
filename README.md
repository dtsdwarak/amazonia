# Amazonia

## What?

Webservice to return Amazon product information - designed to work with book URLs.

___More To Be Updated Soon___

## Running

```bash
$ go run main.go | tee logs
$ curl 'http://localhost:8080/ping?requestURL=https://www.amazon.in/dp/B081HXR95C' | jq
```

## Thanks

1. [Colly Web Scraping Framework](https://go-colly.org/)
2. [Gin Web Framework](https://gin-gonic.com/)