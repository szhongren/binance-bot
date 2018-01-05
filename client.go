package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@localhost/shaoz?sslmode=disable",
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASS"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Exec("CREATE TABLE test(id integer);")
	fmt.Println(rows, err)
	defer db.Close()

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
