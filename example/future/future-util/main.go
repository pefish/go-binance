package main

import (
	"fmt"
	"log"

	"github.com/pefish/go-binance/util"
	t_logger "github.com/pefish/go-interface/t-logger"
	go_logger "github.com/pefish/go-logger"
)

func main() {
	err := do()
	if err != nil {
		log.Fatal(err)
	}
}

func do() error {
	futureUtil := util.NewFutureUtil(go_logger.NewLogger(t_logger.Level_DEBUG))
	symbolInfo, err := futureUtil.SymbolInfo("ONGUSDT")
	if err != nil {
		return err
	}
	fmt.Println(symbolInfo.Filters)
	return nil
}
