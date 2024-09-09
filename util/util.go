package util

import (
	"context"
	"fmt"
	"strings"

	"github.com/pefish/go-binance/futures"
	i_logger "github.com/pefish/go-interface/i-logger"
)

func WsLoopWrapper(
	ctx context.Context,
	logger i_logger.ILogger,
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
			logger.InfoF("Connecting <%s>...", url)
			doneC, stopC, err = futures.WsServe(
				futures.NewWsConfig(fmt.Sprintf("%s/%s", futures.GetWsEndpoint(), url)),
				handler,
				func(err error) {
					if strings.Contains(err.Error(), "connection timed out") {
						logger.InfoF("Connection <%s> timed out, reconnect.", url)
						wsServeChan <- true
					}
				},
			)
			if err != nil {
				return err
			}
			logger.InfoF("Connect <%s> done.", url)
		case <-doneC:
			logger.InfoF("Connection <%s> closed, to reconnect...", url)
			wsServeChan <- true
			continue
		case <-ctx.Done():
			stopC <- struct{}{}
			return nil
		}
	}
}
