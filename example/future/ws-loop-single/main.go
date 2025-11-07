package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/pefish/go-binance/futures"
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
	return futureUtil.WsLoopSingleStream(
		context.Background(),
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
