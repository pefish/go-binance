package util

import (
	"context"
	"github.com/pefish/go-binance/futures"
	"time"
)

func WsLoopWrapper(
	ctx context.Context,
	do func() (doneC, stopC chan struct{}, err error),
	loopInterval time.Duration,
	errHandler futures.ErrHandler,
) {
	for {
		doneC, stopC, err := do()
		if err != nil {
			errHandler(err)
			time.Sleep(loopInterval)
			continue
		}

		select {
		case <-doneC:
			time.Sleep(loopInterval)
			continue
		case <-ctx.Done():
			stopC <- struct{}{}
			return
		}
	}
}
