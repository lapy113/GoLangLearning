package main

import (
	"assesment/httpd/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", handler.PingGet())
	r.GET("/rates/:cryptocurrency/:fiat", handler.RateGet())
	r.GET("/rates/:cryptocurrency", handler.RateCurrencyGet())
	r.GET("/rates", handler.RatePair())
	r.GET("/rates/history/:cryptocurrency/:fiat", handler.RateHourly())
	r.GET("/balance/:address", handler.GetBalance())
	r.Run() // listen and serve on 0.0.0.0:8080
}
