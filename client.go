package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/pdepip/go-binance/binance"
)

func main() {

	client := binance.New(os.Getenv("BINANCE_API"), os.Getenv("BINANCE_SECRET"))
	positions, err := client.GetPositions()

	if err != nil {
		panic(err)
	}

	for _, p := range positions {
		symbol := strings.Join([]string{
			p.Asset,
			"BTC",
		}, "")
		query := binance.SymbolQuery{
			Symbol: symbol,
		}

		res, err := client.GetLastPrice(query)

		if err != nil {
			panic(err)
		}

		fmt.Println(p.Asset, res)
	}

	query := binance.SymbolQuery{
		Symbol: "123456",
	}

	res, err := client.GetLastPrice(query)

	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
