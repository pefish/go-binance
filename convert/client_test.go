package convert

import (
	"os"
	"path"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	i_logger "github.com/pefish/go-interface/i-logger"
	go_test_ "github.com/pefish/go-test"
)

var client *Client

func init() {
	projectRoot, _ := go_test_.ProjectRoot()
	envMap, err := godotenv.Read(path.Join(projectRoot, ".env"))
	if err != nil {
		panic(err)
	}
	for k, v := range envMap {
		os.Setenv(k, v)
	}

	client = New(&i_logger.DefaultLogger, os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
}

func TestClient_SupportedPairs(t *testing.T) {
	pairs, err := client.SupportedPairs(&SupportedPairsParamType{
		// FromAsset: "BTC",
		ToAsset: "MERL",
	})
	go_test_.Equal(t, nil, err)
	spew.Dump(pairs)
}

func TestClient_GetQuote(t *testing.T) {
	quote, err := client.GetQuote(&GetQuoteParamType{
		FromAsset:  "USDT",
		ToAsset:    "币安人生",
		FromAmount: 0.1,
	})
	go_test_.Equal(t, nil, err)
	spew.Dump(quote)
}

// echo -n "fromAmount=0.1&fromAsset=BTC&timestamp=1765774606380&toAsset=USDT" | openssl dgst -sha256 -hmac "s1DDEA3kPy0SiefYLrYU9uIXAOuRTmWg1lC6hkpZSDUIsdzVxGjVfMCl9pT8Dpto"
