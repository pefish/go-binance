package main

import (
	"fmt"
	"github.com/pefish/go-binance/util"
	"log"
)

func main() {
	err := do()
	if err != nil {
		log.Fatal(err)
	}
}

func do() error {
	futureUtil := util.NewFutureUtil()
	symbolInfo, err := futureUtil.SymbolInfo("ONGUSDT")
	if err != nil {
		return err
	}
	fmt.Println(symbolInfo.Filters)
	return nil
}
