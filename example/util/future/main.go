package main

import (
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	future_util "github.com/pefish/go-binance/util/future"
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
	symbolInfo, err := future_util.PairInfo("XCNUSDT")
	if err != nil {
		return err
	}
	spew.Dump(symbolInfo)
	return nil
}
