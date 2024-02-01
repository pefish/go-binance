package util

import (
	"context"
	"github.com/pefish/go-binance/futures"
	"github.com/pkg/errors"
	"time"
)

type FutureUtil struct {
}

func NewFutureUtil() *FutureUtil {
	return &FutureUtil{}
}

func (f *FutureUtil) SymbolInfo(symbol string) (*futures.Symbol, error) {
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
