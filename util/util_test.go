package util_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/pefish/go-binance/util"
	i_logger "github.com/pefish/go-interface/i-logger"
	go_test_ "github.com/pefish/go-test"
)

func TestGetTokenInfo(t *testing.T) {
	tokenInfo, err := util.GetTokenInfo(&i_logger.DefaultLogger, "BTC")
	go_test_.Equal(t, nil, err)
	spew.Dump(tokenInfo)
}
