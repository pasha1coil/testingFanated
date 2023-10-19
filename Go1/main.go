package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type Data struct {
	Id     string  `json:"id"`
	Symbol string  `json:"symbol"`
	Name   string  `json:"name"`
	Price  float64 `json:"current_price"`
}

type DataMu struct {
	sync.RWMutex
	mapa map[string]Data
}

func main() {
	symbol := "btc"
	dtmu := DataMu{
		mapa: make(map[string]Data),
	}

	go func() {
		for {
			update(&dtmu)
			time.Sleep(10 * time.Minute)
		}
	}()
	time.Sleep(5 * time.Second)
	for {
		printRandom(&dtmu)
		time.Sleep(5 * time.Second)
		printSelected(&dtmu, symbol)
		time.Sleep(3 * time.Second)
	}
}

func update(dtmu *DataMu) {
	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1")
	if err != nil {
		log.Fatalf("Ошибка при выполнении HTTP-запроса:%s", err)
	}
	defer resp.Body.Close()

	var coinData []Data
	err = json.NewDecoder(resp.Body).Decode(&coinData)
	if err != nil {
		log.Fatal("Ошибка при обработке resp.Body:", err)
	}

	dtmu.Lock()
	for _, coin := range coinData {
		dtmu.mapa[coin.Symbol] = coin
	}
	dtmu.Unlock()

	log.Println("Данные о курсах обновлены")
}

func printRandom(dtmu *DataMu) {
	dtmu.RLock()
	symbols := make([]string, 0, len(dtmu.mapa))
	for symbol := range dtmu.mapa {
		symbols = append(symbols, symbol)
	}
	randomSymbol := symbols[rand.Intn(len(dtmu.mapa))]
	currency := dtmu.mapa[randomSymbol]
	dtmu.RUnlock()

	fmt.Printf("Symbol: %s\n", currency.Symbol)
	fmt.Printf("Name: %s\n", currency.Name)
	fmt.Printf("Price: %.2f $\n", currency.Price)
	fmt.Println()
}

func printSelected(dtmu *DataMu, symbol string) {
	dtmu.RLock()
	currency := dtmu.mapa[symbol]
	dtmu.RUnlock()
	fmt.Printf("Symbol: %s\n", currency.Symbol)
	fmt.Printf("Name: %s\n", currency.Name)
	fmt.Printf("Price: %.2f $\n", currency.Price)
	fmt.Println()
}
