package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sort"

	_ "github.com/lib/pq"
	coinApi "github.com/miguelmota/go-coinmarketcap"
)

type byMarketCap []coinApi.Coin

func (a byMarketCap) Len() int           { return len(a) }
func (a byMarketCap) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byMarketCap) Less(i, j int) bool { return a[i].MarketCapUsd < a[j].MarketCapUsd }

func getPGConnection() (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@localhost/shaoz?sslmode=disable",
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASS"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

func getCoinsByMarketCap() ([]coinApi.Coin, error) {
	coinMarketData, err := coinApi.GetAllCoinData(50)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var sortedCoinMarketData []coinApi.Coin
	for _, v := range coinMarketData {
		sortedCoinMarketData = append(sortedCoinMarketData, v)
	}
	sort.Sort(sort.Reverse(byMarketCap(sortedCoinMarketData)))
	return sortedCoinMarketData, nil
}

func main() {

	db, err := getPGConnection()
	defer db.Close()

	sortedCoins, err := getCoinsByMarketCap()
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range sortedCoins {
		fmt.Println(v.ID, v.PercentChange24h)
	}
	// client := binance.NewClient(os.Getenv("BINANCE_API"), os.Getenv("BINANCE_SECRET"))

	// res, err := client.NewGetAccountService().Do(context.Background())
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// for i, v := range res.Balances {
	// 	fmt.Println(i, v)
	// }

	// wsDepthHandler := func(event *binance.WsDepthEvent) {
	// 	fmt.Println(event)
	// }
	// done, err := binance.WsDepthServe("LTCBTC", wsDepthHandler)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// <-done

	// 	depthRes, err := client.NewDepthService().Symbol("XLMBTC").
	// 		Do(context.Background())
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	for _, v := range depthRes.Bids {
	// 		fmt.Println(v)
	// 		// largest to smallest limit buy
	// 	}
	// 	fmt.Println("--------------------------------------------------------------------------------")
	// 	for _, v := range depthRes.Asks {
	// 		fmt.Println(v)
	// 		// smallest to largest limit sell
	// 	}
}
