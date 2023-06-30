package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	//"html/template"
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	ID           string  `json:"id"`
	CurrentPrice float32 `json:"current_price"`
}

func checkCoin1(searchValue string) (bool, []string) {
	slice := []string{"bitcoin", "ethereum", "0chain", "achain", "acid", "actinium", "adacash", "adamant", "adapad", "adaswap", "adex", "binamon", "bim", "etherisc", "ethpad", "ethup", "evedo", "falconx", "fanstime", "fantaverse", "dogecoin", "dogeclub", "dogedi"}
	found := false
	for _, item := range slice {
		if item == searchValue {
			found = true
			break
		}
	}

	return found, slice
}

func checkFiat1(searchValue string) (bool, []string) {
	slice := []string{"aed", "ars", "aud", "bch", "bdt", "bhd", "bmd", "bnb", "brl", "btc", "cad", "chf", "clp", "cny", "czk", "dkk", "dot", "eos", "eth", "eur", "gbp", "hkd", "huf", "idr", "ils", "inr", "jpy", "krw", "kwd", "lkr", "ltc", "mmk", "mxn", "myr", "ngn", "nok", "nzd", "php", "pkr", "pln", "rub", "sar", "sek", "sgd", "thb", "try", "twd", "uah", "usd", "vef", "vnd", "xag", "xau", "xdr", "xlm", "xrp", "yfi", "zar", "bits", "link", "sats"}

	found := false
	for _, item := range slice {
		if item == searchValue {
			found = true
			break
		}
	}

	return found, slice
}

func RateGet() gin.HandlerFunc {
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

		fiat := c.Param("fiat")
		foundF, validFiats := checkFiat1(fiat)
		if !foundF {
			var invalidFiat InvalidTokken
			invalidFiat.ErrorMsg = "Invalid fiat,check out the valid fiats"
			invalidFiat.ValidTokken = validFiats
			c.JSON(http.StatusOK, invalidFiat)
			return
		}
		//TODO add validation check and error handling

		// create url using parameter
		url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/markets?vs_currency=%s&ids=%s&order=market_cap_desc&per_page=100&page=1&sparkline=false&locale=en", fiat, cryptocurrency)

		// get response from url
		response, err := http.Get(url)

		// get data from body and attach it to variable data
		var data []ResponseData
		keyval, _ := ioutil.ReadAll(response.Body)
		err = json.Unmarshal(keyval, &data)

		if err != nil {
			c.JSON(http.StatusOK, err)
		} else {
			c.JSON(http.StatusOK, data)
		}

	}
}
