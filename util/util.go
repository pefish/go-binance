package util

import (
	"github.com/pefish/go-binance/futures"
	"time"
)

func WsLoopWrapper(
	do func() (doneC, stopC chan struct{}, err error),
	loopInterval time.Duration,
	errHandler futures.ErrHandler,
) {
	for {
		doneC, _, err := do()
		if err != nil {
			errHandler(err)
			time.Sleep(loopInterval)
			continue
		}

		select {
		case <-doneC:
			time.Sleep(loopInterval)
			continue
		}
	}
}
