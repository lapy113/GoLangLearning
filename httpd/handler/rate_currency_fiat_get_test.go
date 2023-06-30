package handler

import (
	"fmt"
	"testing"
)

func TestCheckCoin1(t *testing.T) {
	slice := []string{"bitcoin", "ethereum", "0chain", "achain", "acid", "actinium", "adacash", "adamant", "adapad", "adaswap", "adex", "binamon", "bim", "etherisc", "ethpad", "ethup", "evedo", "falconx", "fanstime", "fantaverse", "dogecoin", "dogeclub", "dogedi"}
	for i := 0; i < len(slice); i++ {
		str := slice[i]
		value, _ := checkCoin1(str)
		if value == false {
			t.Error("Error in CheckCoin1")
		}

	}

}

func TestCheckFiat1(t *testing.T) {
	slice := []string{"aed", "ars", "aud", "bch", "bdt", "bhd", "bmd", "bnb", "brl", "btc", "cad", "chf", "clp", "cny", "czk", "dkk", "dot", "eos", "eth", "eur", "gbp", "hkd", "huf", "idr", "ils", "inr", "jpy", "krw", "kwd", "lkr", "ltc", "mmk", "mxn", "myr", "ngn", "nok", "nzd", "php", "pkr", "pln", "rub", "sar", "sek", "sgd", "thb", "try", "twd", "uah", "usd", "vef", "vnd", "xag", "xau", "xdr", "xlm", "xrp", "yfi", "zar", "bits", "link", "sats"}
	for i := 0; i < len(slice); i++ {
		str := slice[i]
		value, _ := checkFiat1(str)
		if value == false {
			t.Error("Error in CheckCoin1")
		}

	}

}

func TestRateGet(t *testing.T) {

	val := RateGet()
	fmt.Println(val)
	// if types(val) == gin.HandlerFunc {

	// }
}
