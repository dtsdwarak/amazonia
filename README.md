# Amazonia

## What?

Webservice to return Amazon.in product information - designed to work with book URLs.

___More To Be Updated Soon___

## Running

```bash
$ make run
$ curl 'http://localhost:8080/product?requestURL=https://www.amazon.in/dp/B081HXR95C' | jq
```

## Thanks

1. [Colly Web Scraping Framework](https://go-colly.org/)
2. [Gin Web Framework](https://gin-gonic.com/)