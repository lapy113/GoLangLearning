package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CurrencyResponseData struct {
	ID         string `json:"id"`
	MarketData struct {
		CurrentPrice map[string]float64 `json:"current_price"`
	} `json:"market_data"`
}

func RateCurrencyGet() gin.HandlerFunc {
	return func(c *gin.Context) {

		//get params from url
		cryptocurrency := c.Param("cryptocurrency")

		foundC, validCoins := checkCoin1(cryptocurrency)
		if !foundC {
			var invalidCrypto InvalidTokken
			invalidCrypto.ErrorMsg = "Invalid crypto coin,check out the valid coins"
			invalidCrypto.ValidTokken = validCoins
			c.JSON(http.StatusOK, invalidCrypto)
			return
		}

		currentTime := time.Now()
		formattedDate := currentTime.Format("02-01-2006")

		//TODO add validation check and error handling for crypto name

		url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s/history?date=%s&localization=false", cryptocurrency, formattedDate)
		response, err := http.Get(url)
		var data CurrencyResponseData
		keyval, _ := ioutil.ReadAll(response.Body)

		err = json.Unmarshal(keyval, &data)
		if err != nil {
			c.JSON(http.StatusOK, err)
		} else {
			c.JSON(http.StatusOK, data)

		}

	}
}
