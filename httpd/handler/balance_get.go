package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

func GetBalance() gin.HandlerFunc {
	return func(c *gin.Context) {

		address := c.Param("address")
		//coin validation check

		client, err := ethclient.Dial("https://cloudflare-eth.com")
		if err != nil {
			log.Fatal(err)
		}

		account := common.HexToAddress(address)
		bal, err := client.BalanceAt(context.Background(), account, nil)
		if err != nil {
			log.Fatal(err)
		}
		balanceStr := fmt.Sprint(bal)

		balRes := fmt.Sprintf("Balance: %s.%s Ether", balanceStr[:len(balanceStr)-18], balanceStr[len(balanceStr)-18:])
		if balanceStr == "11739100857380173743863" {
			c.JSON(http.StatusOK, "Invalid Ether Address")
			return
		} else {
			c.JSON(http.StatusOK, balRes)
		}

	}
}
