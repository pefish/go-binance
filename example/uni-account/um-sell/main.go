package main

import (
	"context"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	uni_account "github.com/pefish/go-binance/uni-account"
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
	uniAccountClient := uni_account.NewClient(
		os.Getenv("API_KEY"),
		os.Getenv("API_SECRET"),
	)

	symbol := "FLMUSDT"

	res, err := uniAccountClient.
		NewCreateUMOrderService().
		Symbol(symbol).
		Side(uni_account.SideTypeSell).
		Type(uni_account.OrderTypeMarket).
		Quantity("340").
		Do(context.Background())
	if err != nil {
		return err
	}

	order, err := uniAccountClient.NewGetUMOrderService().Symbol(symbol).OrderID(res.OrderID).Do(context.Background())
	if err != nil {
		return err
	}
	spew.Dump(order)

	return nil
}
