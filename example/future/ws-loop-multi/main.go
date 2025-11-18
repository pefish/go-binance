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
	return future_util.WsLoopMultiStream(
		context.Background(),
		&i_logger.DefaultLogger,
		[]string{
			"BTCUSDT@kline_5m",
			"ETHUSDT@kline_5m",
		},
		func(message []byte) {
			type DataType struct {
				Stream string               `json:"stream"`
				Data   futures.WsKlineEvent `json:"data"`
			}
			event := new(DataType)
			err := json.Unmarshal(message, event)
			if err != nil {
				fmt.Println(err)
				return
			}
			switch event.Stream {
			case "btcusdt@kline_5m":
				fmt.Println("BTCUSDT", event.Data.Time, event.Data.Kline.StartTime, event.Data.Kline.EndTime, event.Data.Kline.Close, event.Data.Kline.Volume)
			case "ethusdt@kline_5m":
				fmt.Println("ETHUSDT", event.Data.Time, event.Data.Kline.StartTime, event.Data.Kline.EndTime, event.Data.Kline.Close, event.Data.Kline.Volume)
			}
		},
	)
}
