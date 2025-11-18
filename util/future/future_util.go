package future_util

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/pefish/go-binance/futures"
	i_logger "github.com/pefish/go-interface/i-logger"
	"github.com/pkg/errors"
)

func SymbolInfo(symbol string) (*futures.Symbol, error) {
	binanceFutureClient := futures.NewClient("", "")
	newCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	exchangeInfo, err := binanceFutureClient.NewExchangeInfoService().Do(newCtx)
	if err != nil {
		return nil, err
	}
	for _, e := range exchangeInfo.Symbols {
		if e.Symbol == symbol {
			return &e, nil
		}
	}

	return nil, errors.New("Symbol not found.")
}

func WsLoopSingleStream(
	ctx context.Context,
	logger i_logger.ILogger,
	pair string,
	dataType string,
	handler func(msg []byte),
) error {
	streamName := fmt.Sprintf("%s@%s", strings.ToLower(pair), dataType)
	url := fmt.Sprintf("%s/%s", futures.GetWsEndpoint(), streamName)
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
				futures.NewWsConfig(url),
				handler,
				func(err error) {
					if strings.Contains(err.Error(), "connection timed out") {
						logger.InfoF("Connection <%s> timed out, reconnect.", url)
						wsServeChan <- true
					} else {
						logger.ErrorF("Connection <%s> error: %v", url, err)
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
			doneC = nil // 阻止这个分支被多次执行
			continue
		case <-ctx.Done():
			stopC <- struct{}{}
			return nil
		}
	}
}

func WsLoopMultiStream(
	ctx context.Context,
	logger i_logger.ILogger,
	streamNames []string, // e.g. btcusdt@kline_5m
	handler func(msg []byte),
) error {
	realStreamNames := make([]string, 0)
	for _, streamName := range streamNames {
		realStreamNames = append(realStreamNames, strings.ToLower(streamName))
	}
	url := fmt.Sprintf("%s/%s", futures.GetCombinedEndpoint(), strings.Join(realStreamNames, "/"))
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
				futures.NewWsConfig(url),
				handler,
				func(err error) {
					if strings.Contains(err.Error(), "connection timed out") {
						logger.InfoF("Connection <%s> timed out, reconnect.", url)
						wsServeChan <- true
					} else {
						logger.ErrorF("Connection <%s> error: %v", url, err)
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
			doneC = nil // 阻止这个分支被多次执行
			continue
		case <-ctx.Done():
			stopC <- struct{}{}
			return nil
		}
	}
}
