package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pefish/go-binance/spot"
)

func main() {
	envMap, _ := godotenv.Read("./.env")
	for k, v := range envMap {
		os.Setenv(k, v)
	}

	err := do()
	if err != nil {
		log.Fatalf("%+v", err)
	}
}

func do() error {
	spotClient := spot.NewClient(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
	res, err := spotClient.NewGetAllCoinsInfoService().Do(context.Background())
	if err != nil {
		return err
	}
	for _, r := range res {
		fmt.Printf("<coin: %s> <networks: %+v>", r.Coin, r.NetworkList)
	}
	return nil
}
