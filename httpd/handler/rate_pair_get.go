package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PairResponseData struct {
	USD float64 `json:"USD"`
	EUR float64 `json:"EUR"`
	INR float64 `json:"INR"`
	AED float64 `json:"AED"`
	AUD float64 `json:"AUD"`
	JPY float64 `json:"JPY"`
	CAD float64 `json:"CAD"`
	CNY float64 `json:"CNY"`
}

func RatePair() gin.HandlerFunc {
	return func(c *gin.Context) {

		url := fmt.Sprintf("https://min-api.cryptocompare.com/data/pricemulti?fsyms=BTC,ETH,AETH,TOX,KUJI,AKNC,AUSDC,ALINK,FUMO,HIPPO,SHS,QUANT,HOBO,GPBP,PAAL,LITHO,X,SCAM,FXY,BULLS,KUSA,STRAY,UWU&tsyms=USD,EUR,INR,AED,AUD,JPY,CAD,CNY")

		// get response from url
		response, err := http.Get(url)

		// get data from body and attach it to variable data

		keyval, _ := ioutil.ReadAll(response.Body)

		var currencies map[string]PairResponseData
		err = json.Unmarshal([]byte(keyval), &currencies)
		if err != nil {
			c.JSON(http.StatusOK, err)
			return
		}

		c.JSON(http.StatusOK, currencies)

	}
}
