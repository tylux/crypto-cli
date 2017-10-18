package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	coin = flag.String("c", "", "The Cryptocoin to get a price for")
)

type CoinMarketCapData struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	Rank             string `json:"rank"`
	PriceUsd         string `json:"price_usd"`
	PriceBtc         string `json:"price_btc"`
	Prev24VolumeUsd  string `json:"24h_volume_usd"`
	MarketCapUsd     string `json:"market_cap_usd"`
	AvailableSupply  string `json:"available_supply"`
	TotalSupply      string `json:"total_supply"`
	PercentChange1H  string `json:"percent_change_1h"`
	PercentChange24H string `json:"percent_change_24h"`
	PercentChange7D  string `json:"percent_change_7d"`
	LastUpdated      string `json:"last_updated"`
}

func getJson(url string, target interface{}) error {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	r, err := client.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

func main() {

	flag.Parse()
	if *coin == "" {
		url := fmt.Sprintf("https://api.coinmarketcap.com/v1/ticker/?limit=10")
		priceData := make([]CoinMarketCapData, 0)
		err := getJson(url, &priceData)
		if err != nil {
			log.Printf("Failed to decode json: %v", err)
		} else {
			//log.Printf("Price: %v", priceData[0].PriceUsd)
		}
		//fmt.Printf("%s price is: %s", priceData[0].Name, priceData[0].PriceUsd)
		for i := range priceData {
			fmt.Fprintf(os.Stdout, "%s: $%s | %%Change(24h): %s\n", priceData[i].Name, priceData[i].PriceUsd, priceData[i].PercentChange24H)
		}
	} else {
		url := fmt.Sprintf("https://api.coinmarketcap.com/v1/ticker/%s", *coin)
		priceData := make([]CoinMarketCapData, 0)
		err := getJson(url, &priceData)
		if err != nil {
			log.Printf("Failed to decode json: %v", err)
		} else {
			//log.Printf("Price: %v", priceData[0].PriceUsd)
		}
		//fmt.Printf("%s price is: %s", priceData[0].Name, priceData[0].PriceUsd)
		for i := range priceData {
			fmt.Fprintf(os.Stdout, "%s: $%s | %%Change(24h): %s\n", priceData[i].Name, priceData[i].PriceUsd, priceData[i].PercentChange24H)
		}
	}

}
