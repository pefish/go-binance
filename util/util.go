package util

import (
	"context"
	"fmt"
	"strings"

	"github.com/pefish/go-binance/futures"
	go_logger "github.com/pefish/go-logger"
)

func WsLoopWrapper(
	ctx context.Context,
	logger go_logger.InterfaceLogger,
	url string,
	handler func(msg []byte),
) error {
	wsServeChan := make(chan bool, 1)
	wsServeChan <- true
	var doneC chan struct{}
	var stopC chan struct{}
	var err error
	for {
		select {
		case <-wsServeChan:
			logger.InfoF("Connecting %s...", url)
			doneC, stopC, err = futures.WsServe(
				futures.NewWsConfig(fmt.Sprintf("%s/%s", futures.GetWsEndpoint(), url)),
				handler,
				func(err error) {
					if strings.Contains(err.Error(), "connection timed out") {
						logger.InfoF("Connection timed out, reconnect.")
						wsServeChan <- true
					}
				},
			)
			if err != nil {
				return err
			}
			logger.InfoF("Connect done.")
		case <-doneC:
			logger.InfoF("Connection closed, to reconnect...")
			wsServeChan <- true
			continue
		case <-ctx.Done():
			stopC <- struct{}{}
			return nil
		}
	}
}
