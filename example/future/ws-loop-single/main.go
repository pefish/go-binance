package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/pefish/go-binance/futures"
	future_util "github.com/pefish/go-binance/util/future"
	i_logger "github.com/pefish/go-interface/i-logger"
)

func main() {
	err := do()
	if err != nil {
		log.Fatal(err)
	}
}

func do() error {
	return future_util.WsLoopSingleStream(
		context.Background(),
		&i_logger.DefaultLogger,
		"BTCUSDT",
		"kline_5m",
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
