package main

import (
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"github.com/pefish/go-binance/util"
	t_logger "github.com/pefish/go-interface/t-logger"
	go_logger "github.com/pefish/go-logger"
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
	futureUtil := util.NewFutureUtil(go_logger.NewLogger(t_logger.Level_DEBUG))
	symbolInfo, err := futureUtil.SymbolInfo("XCNUSDT")
	if err != nil {
		return err
	}
	spew.Dump(symbolInfo)
	return nil
}
