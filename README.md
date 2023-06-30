# Install Golang on your platform

# Install gin
go get -u github.com/gin-gonic/gin

# To add feature of getting balance of an eth address we need to install the following dependency
go get github.com/ethereum/go-ethereum

# In case of dependency issues run
go mod tidy

# To run 
# go inside httpd
cd httpd

# Start server using the below command
go run main.go

# In browser open 
http://0.0.0.0:8080/
# to test out various endpoints

# Endpoints
# to test if server is running
/ping

# to check rate of a crypto against a fiat
/rates/:cryptocurrency/:fiat

# to check rate of a crypto against all fiat
/rates/:cryptocurrency

# to check rate of a crypto against all fiat
/rates

# to check 24 hr data of a crypto
/rates/history/:cryptocurrency/:fiat

# to check balance of an eth wallet
/balance/{address}



# Sample url for each endpoints for Testing in local
http://0.0.0.0:8080/ping
http://0.0.0.0:8080/rates/bitcoin/usd
http://0.0.0.0:8080/rates/bitcoin
http://0.0.0.0:8080/rates
http://0.0.0.0:8080/rates/history/BTC/USD
http://0.0.0.0:8080/balance/0x1f9090aaE28b8a3dCeaDf281B0F12828e676c326