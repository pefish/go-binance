package main

import (
	"context"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"github.com/pefish/go-binance/futures"
)

func main() {
	envMap, _ := godotenv.Read("./.env")
	for k, v := range envMap {
		os.Setenv(k, v)
	}

	err := do()
	if err != nil {
		log.Fatal(err)
	}
}

func do() error {
	binanceFutureClient := futures.NewClient(
		os.Getenv("BINANCE_API_KEY"),
		os.Getenv("BINANCE_API_SECRET"),
	)
	prices, err := binanceFutureClient.NewListPricesService().Pair("XCN1USDT").Do(context.Background())
	if err != nil {
		return err
	}
	spew.Dump(prices)
	return nil
}
