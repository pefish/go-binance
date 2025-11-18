package main

import (
	"context"
	"fmt"
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

	res, err := binanceFutureClient.
		NewCreateOrderService().
		Symbol(symbol).
		Side(futures.SideTypeBuy).
		Type(futures.OrderTypeMarket).
		Quantity("2000").
		Do(context.Background())
	if err != nil {
		return err
	}

	order, err := binanceFutureClient.NewGetOrderService().Symbol(symbol).OrderID(res.OrderID).Do(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf(
		`
		AvgPrice: %s
		OrderID: %d
		Price: %s
		Symbol: %s
		PositionSide: %s
		CumQuote: %s
		ExecutedQuantity: %s
		`,
		order.AvgPrice,
		order.OrderID,
		order.Side,
		order.Symbol,
		order.PositionSide,
		order.CumQuote,
		order.ExecutedQuantity,
	)

	return nil
}
