package handler

import (
	"fmt"
	"testing"
)

func TestCheckCoin(t *testing.T) {
	slice := []string{"BTC", "ETH", "AETH", "TOX", "KUJI", "AKNC", "AUSDC", "ALINK", "FUMO", "HIPPO", "SHS", "QUANT", "HOBO", "GPBP", "PAAL", "LITHO", "X", "SCAM", "FXY", "BULLS", "KUSA", "STRAY", "UWU"}
	for i := 0; i < len(slice); i++ {
		str := slice[i]
		value, _ := checkCoin(str)
		if value == false {
			t.Error("Error in CheckCoin1")
		}

	}

}

func TestCheckFiat(t *testing.T) {
	slice := []string{"USD", "EUR", "INR", "AED", "AUD", "JPY", "CAD", "CNY"}
	for i := 0; i < len(slice); i++ {
		str := slice[i]
		value, _ := checkFiat(str)
		if value == false {
			t.Error("Error in CheckCoin1")
		}

	}

}

func TestRateHourly(t *testing.T) {

	val := RateHourly()
	fmt.Println(val)
	// if types(val) == gin.HandlerFunc {

	// }
}
