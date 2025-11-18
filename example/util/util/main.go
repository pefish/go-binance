package main

import (
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"github.com/pefish/go-binance/util"
	i_logger "github.com/pefish/go-interface/i-logger"
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
	info, err := util.GetCirculatingSupplyInfo(&i_logger.DefaultLogger, "XCN")
	if err != nil {
		return err
	}
	spew.Dump(info)
	return nil
}
