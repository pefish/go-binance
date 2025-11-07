package main

import (
	"context"
	"fmt"
	"log"

	"github.com/pefish/go-binance/futures"
)

func main() {
	err := do()
	if err != nil {
		log.Fatal(err)
	}
}

func do() error {
	//startTime := 1704246284744
	binanceFutureClient := futures.NewClient(
		"",
		"",
	)

	symbol := "BTCUSDT"

	res, err := binanceFutureClient.
		NewCreateOrderService().
		Symbol(symbol).
		Side(futures.SideTypeBuy).
		Type(futures.OrderTypeMarket).
		Quantity("0.01").
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
