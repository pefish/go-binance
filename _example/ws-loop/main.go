package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/pefish/go-binance/futures"
	"github.com/pefish/go-binance/util"
	go_logger "github.com/pefish/go-logger"
)

func main() {
	err := do()
	if err != nil {
		log.Fatal(err)
	}
}

func do() error {
	return util.WsLoopWrapper(
		context.Background(),
		go_logger.Logger,
		fmt.Sprintf("%s@kline_%s", strings.ToLower("BTCUSDT"), "5m"),
		func(message []byte) {
			event := new(futures.WsKlineEvent)
			err := json.Unmarshal(message, event)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(event.Time, event.Kline.StartTime, event.Kline.EndTime, event.Kline.Close, event.Kline.Volume)
		},
	)
}
