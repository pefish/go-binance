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
	uniAccountClient := uni_account.NewClient(
		os.Getenv("API_KEY"),
		os.Getenv("API_SECRET"),
	)

	symbol := "FLMUSDT"

	r, err := uniAccountClient.NewChangeUMLeverageService().Symbol(symbol).Leverage(5).Do(context.Background())
	if err != nil {
		return err
	}
	spew.Dump(r)

	return nil
}
