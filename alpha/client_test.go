package alpha_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/pefish/go-binance/alpha"
	i_logger "github.com/pefish/go-interface/i-logger"
	go_test_ "github.com/pefish/go-test"
)

var client *alpha.Client

func init() {
	client = alpha.New(&i_logger.DefaultLogger)
}

func TestClient_ListTokens(t *testing.T) {
	tokens, err := client.ListTokens()
	go_test_.Equal(t, nil, err)
	spew.Dump(tokens)
}
