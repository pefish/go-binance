package alpha_test

import (
	"os"
	"path"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"github.com/pefish/go-binance/alpha"
	i_logger "github.com/pefish/go-interface/i-logger"
	go_test_ "github.com/pefish/go-test"
)

var client *alpha.Client

func init() {
	client = alpha.New(&i_logger.DefaultLogger)
	projectRoot, _ := go_test_.ProjectRoot()
	envMap, err := godotenv.Read(path.Join(projectRoot, ".env"))
	if err != nil {
		panic(err)
	}
	for k, v := range envMap {
		os.Setenv(k, v)
	}
}

func TestClient_ListTokens(t *testing.T) {
	tokens, err := client.ListTokens()
	go_test_.Equal(t, nil, err)
	spew.Dump(tokens)
}
