This project is created to build a Go based application to fetch cryptocurrencies current rates in CAD using an external API.

Files in the Repo:
    main.go - It contains basic functions in Golang to fetch the data using an external API
    main_test.go - It contains test cases for the functions present in main.go file
    Dockerfile - It contains the docker code to containerize and create image for the application
    go.mod 
    README.md

To use external API we have looked various websites and found the below mentioned APIs which fetch the required data:
	https://api.kucoin.com/api/v1/market/stats?symbol=BTC-USDT
    https://api.coinbase.com/v2/exchange-rates?currency=BTC
    https://api.coincap.io/v2/assets/bitcoin 
    
From the above APIs we choose coincap due to its simplicity.

