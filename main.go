package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	baseURL     = "https://api.coincap.io/v2/assets"
	apiKey      = "YOUR_COINCAP_API_KEY"
	exchangeURL = "https://api.exchangerate-api.com/v4/latest/USD"
)

type CoinCapResponse struct {
	Data struct {
		PriceUsd string `json:"priceUsd"`
	} `json:"data"`
}

type ExchangeRateResponse struct {
	Rates struct {
		CAD float64 `json:"CAD"`
	} `json:"rates"`
}

func getPrice(symbol string) (string, error) {
	url := fmt.Sprintf("%s/%s", baseURL, symbol)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data CoinCapResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	if data.Data.PriceUsd != "" {
		exchangeRate, err := getExchangeRate()
		if err != nil {
			return "", err
		}

		priceUSD := data.Data.PriceUsd
		priceCAD := fmt.Sprintf("%.2f", exchangeRate*toNumeric(priceUSD))
		return priceCAD, nil
	}

	return "", fmt.Errorf("price for %s not found", symbol)
}

func getExchangeRate() (float64, error) {
	resp, err := http.Get(exchangeURL)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var data ExchangeRateResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return 0, err
	}

	return data.Rates.CAD, nil
}

func toNumeric(s string) float64 {
	var result float64
	_, err := fmt.Sscanf(s, "%f", &result)
	if err != nil {
		return 0
	}
	return result
}

func bitcoinPrice(w http.ResponseWriter, r *http.Request) {
	price, err := getPrice("bitcoin")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Bitcoin price: %s CAD\n", price)
}

func ethereumPrice(w http.ResponseWriter, r *http.Request) {
	price, err := getPrice("ethereum")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Ethereum price: %s CAD\n", price)
}

func tetherPrice(w http.ResponseWriter, r *http.Request) {
	price, err := getPrice("tether")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Tether price: %s CAD\n", price)
}

func main() {
	http.HandleFunc("/bitcoin-price", bitcoinPrice)
	http.HandleFunc("/ethereum-price", ethereumPrice)
	http.HandleFunc("/tether-price", tetherPrice)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
