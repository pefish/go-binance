package main

import (
	"context"
	"log"
	"os"

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
	//startTime := 1704246284744
	binanceFutureClient := futures.NewClient(
		os.Getenv("API_KEY"),
		os.Getenv("API_SECRET"),
	)

	symbol := "FLMUSDT"

	_, err := binanceFutureClient.NewChangeLeverageService().Symbol(symbol).Leverage(5).Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}
