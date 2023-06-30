package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CurrencyData struct {
	Time             int64   `json:"time"`
	High             float64 `json:"high"`
	Low              float64 `json:"low"`
	Open             float64 `json:"open"`
	VolumeFrom       float64 `json:"volumefrom"`
	VolumeTo         float64 `json:"volumeto"`
	Close            float64 `json:"close"`
	ConversionType   string  `json:"conversionType"`
	ConversionSymbol string  `json:"conversionSymbol"`
}

type CryptoResponse struct {
	Data CryptoData `json:"Data"`
}

type CryptoData struct {
	Data []CurrencyData `json:"Data"`
}

func checkCoin(searchValue string) (bool, []string) {
	slice := []string{"BTC", "ETH", "AETH", "TOX", "KUJI", "AKNC", "AUSDC", "ALINK", "FUMO", "HIPPO", "SHS", "QUANT", "HOBO", "GPBP", "PAAL", "LITHO", "X", "SCAM", "FXY", "BULLS", "KUSA", "STRAY", "UWU"}
	found := false
	for _, item := range slice {
		if item == searchValue {
			found = true
			break
		}
	}

	return found, slice
}

func checkFiat(searchValue string) (bool, []string) {
	slice := []string{"USD", "EUR", "INR", "AED", "AUD", "JPY", "CAD", "CNY"}

	found := false
	for _, item := range slice {
		if item == searchValue {
			found = true
			break
		}
	}

	return found, slice
}

type InvalidTokken struct {
	ErrorMsg    string   `json:"error_msg"`
	ValidTokken []string `json:"valid_coins"`
}

func RateHourly() gin.HandlerFunc {
	return func(c *gin.Context) {

		cryptocurrency := c.Param("cryptocurrency")
		//coin validation check
		foundC, validCoins := checkCoin(cryptocurrency)
		if !foundC {
			var invalidCrypto InvalidTokken
			invalidCrypto.ErrorMsg = "Invalid crypto coin,check out the valid coins"
			invalidCrypto.ValidTokken = validCoins
			c.JSON(http.StatusOK, invalidCrypto)
			return
		}

		fiat := c.Param("fiat")
		foundF, validFiats := checkFiat(fiat)
		if !foundF {
			var invalidFiat InvalidTokken
			invalidFiat.ErrorMsg = "Invalid fiat,check out the valid fiats"
			invalidFiat.ValidTokken = validFiats
			c.JSON(http.StatusOK, invalidFiat)
			return
		}

		url := fmt.Sprintf("https://min-api.cryptocompare.com/data/v2/histohour?fsym=%s&tsym=%s&limit=24", cryptocurrency, fiat)

		// get response from url
		response, err := http.Get(url)

		// get data from body and attach it to variable data

		keyval, _ := ioutil.ReadAll(response.Body)

		var cryptoResponse CryptoResponse
		err = json.Unmarshal([]byte(keyval), &cryptoResponse)
		if err != nil {
			c.JSON(http.StatusBadGateway, err)
			return
		}

		c.JSON(http.StatusOK, cryptoResponse)

	}
}
