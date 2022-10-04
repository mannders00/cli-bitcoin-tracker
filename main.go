package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.blockchain.com/v3/exchange/tickers/BTC-USD")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var btc map[string]interface{}
	json.Unmarshal([]byte(body), &btc)

	fmt.Println(btc["price_24h"])
}
