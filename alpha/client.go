package alpha

import (
	"time"

	go_http "github.com/pefish/go-http"
	i_logger "github.com/pefish/go-interface/i-logger"
	"github.com/pkg/errors"
)

type Client struct {
	logger i_logger.ILogger
}

func New(logger i_logger.ILogger) *Client {
	return &Client{
		logger: logger,
	}
}

type TokenInfoType struct {
	TokenId           string  `json:"tokenId"`
	ChainId           string  `json:"chainId"`
	ChainIconUrl      string  `json:"chainIconUrl"`
	ChainName         string  `json:"chainName"`
	ContractAddress   string  `json:"contractAddress"`
	Name              string  `json:"name"`
	Symbol            string  `json:"symbol"`
	IconUrl           string  `json:"iconUrl"`
	Price             float64 `json:"price,string"`
	PercentChange24h  float64 `json:"percentChange24h,string"`
	Volume24h         float64 `json:"volume24h,string"`
	MarketCap         float64 `json:"marketCap,string"`
	Fdv               float64 `json:"fdv,string"`
	Liquidity         float64 `json:"liquidity,string"`
	TotalSupply       float64 `json:"totalSupply,string"`
	CirculatingSupply float64 `json:"circulatingSupply,string"`
	Holders           int     `json:"holders,string"`
	Decimals          int     `json:"decimals"`
	ListingCex        bool    `json:"listingCex"`
	HotTag            bool    `json:"hotTag"`
	CexCoinName       string  `json:"cexCoinName"`
	CanTransfer       bool    `json:"canTransfer"`
	Denomination      int     `json:"denomination"`
	Offline           bool    `json:"offline"`
	TradeDecimal      int     `json:"tradeDecimal"`
	AlphaId           string  `json:"alphaId"`
	Offsell           bool    `json:"offsell"`
	PriceHigh24h      float64 `json:"priceHigh24h,string"`
	PriceLow24h       float64 `json:"priceLow24h,string"`
	Count24h          int     `json:"count24h,string"`
	OnlineTge         bool    `json:"onlineTge"`
	OnlineAirdrop     bool    `json:"onlineAirdrop"`
	Score             int     `json:"score"`
	CexOffDisplay     bool    `json:"cexOffDisplay"`
	StockState        bool    `json:"stockState"`
	ListingTime       int64   `json:"listingTime"`
	MulPoint          int     `json:"mulPoint"`
	BnExclusiveState  bool    `json:"bnExclusiveState"`
}

func (t *Client) ListTokens() ([]*TokenInfoType, error) {
	var httpResult struct {
		Code    string           `json:"code"`
		Msg     string           `json:"message"`
		Data    []*TokenInfoType `json:"data"`
		Success bool             `json:"success"`
	}
	_, _, err := go_http.NewHttpRequester(
		go_http.WithLogger(t.logger),
		go_http.WithTimeout(5*time.Second),
	).GetForStruct(
		&go_http.RequestParams{
			Url: "https://www.binance.com/bapi/defi/v1/public/wallet-direct/buw/wallet/cex/alpha/all/token/list",
		},
		&httpResult,
	)
	if err != nil {
		return nil, err
	}
	if httpResult.Code != "000000" {
		return nil, errors.New(httpResult.Msg)
	}
	return httpResult.Data, nil
}
