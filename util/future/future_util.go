package future_util

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pefish/go-binance/futures"
	i_logger "github.com/pefish/go-interface/i-logger"
	"github.com/pkg/errors"
)

type PairInfoType struct {
	futures.Symbol
	MinQuantity    float64 // 最小下单数量
	MaxQuantity    float64
	OrderPrecision int     // 下单数量精度
	MinUAmount     float64 // 最小下单金额（USDT）
}

func PairInfo(pair string) (*PairInfoType, error) {
	binanceFutureClient := futures.NewClient("", "")
	newCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	exchangeInfo, err := binanceFutureClient.NewExchangeInfoService().Do(newCtx)
	if err != nil {
		return nil, err
	}
	for _, e := range exchangeInfo.Symbols {
		if e.Symbol == pair {
			pairInfo := &PairInfoType{
				Symbol: e,
			}
			for _, map_ := range e.Filters {
				filterType := futures.SymbolFilterType(map_["filterType"].(string))
				switch filterType {
				case futures.SymbolFilterTypeLotSize:
					minQty, err := strconv.ParseFloat(map_["minQty"].(string), 64)
					if err != nil {
						return nil, err
					}
					maxQty, err := strconv.ParseFloat(map_["maxQty"].(string), 64)
					if err != nil {
						return nil, err
					}
					stepSize, err := strconv.ParseFloat(map_["stepSize"].(string), 64)
					if err != nil {
						return nil, err
					}
					orderPrecision := 0
					for step := stepSize; step < 1; step *= 10 {
						orderPrecision++
					}
					pairInfo.MinQuantity = minQty
					pairInfo.MaxQuantity = maxQty
					pairInfo.OrderPrecision = orderPrecision
				case futures.SymbolFilterTypeMinNotional:
					minNotional, err := strconv.ParseFloat(map_["notional"].(string), 64)
					if err != nil {
						return nil, err
					}
					pairInfo.MinUAmount = minNotional
				}
			}
			return pairInfo, nil
		}
	}

	return nil, errors.New("Pair not found.")
}

func Price(pair string) (float64, error) {
	binanceFutureClient := futures.NewClient("", "")
	newCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	prices, err := binanceFutureClient.NewListPricesService().Pair(pair).Do(newCtx)
	if err != nil {
		return 0, err
	}
	if len(prices) != 1 {
		return 0, errors.Errorf("pairName %s prices length not 1", pair)
	}

	price, err := strconv.ParseFloat(prices[0].Price, 64)
	if err != nil {
		return 0, err
	}
	return price, nil
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
