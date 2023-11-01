package main

import (
	"context"
	"fmt"
	"github.com/pefish/go-binance/futures"
	"github.com/pefish/go-binance/util"
	"time"
)

func main() {

	util.WsLoopWrapper(
		context.Background(),
		func() (doneC, stopC chan struct{}, err error) {
			return futures.WsKlineServe("BTCUSDT", "1m", func(event *futures.WsKlineEvent) {
				fmt.Println(event.Time, event.Kline.Close)
				return
			}, func(err error) {
				if err != nil {
					fmt.Println(err)
				}
			})
		},
		5*time.Second,
		func(err error) {
			if err != nil {
				fmt.Println(err)
			}
		},
	)

	select {}

}
