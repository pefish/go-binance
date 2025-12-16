package alpha

import (
	go_format_type "github.com/pefish/go-format/type"
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
	TokenId           string                       `json:"tokenId"`
	ChainId           string                       `json:"chainId"`
	ChainIconUrl      string                       `json:"chainIconUrl"`
	ChainName         string                       `json:"chainName"`
	ContractAddress   string                       `json:"contractAddress"`
	Name              string                       `json:"name"`
	Symbol            string                       `json:"symbol"`
	IconUrl           string                       `json:"iconUrl"`
	Price             go_format_type.Float64String `json:"price"`
	PercentChange24h  go_format_type.Float64String `json:"percentChange24h"`
	Volume24h         go_format_type.Float64String `json:"volume24h"`
	MarketCap         go_format_type.Float64String `json:"marketCap"`
	Fdv               go_format_type.Float64String `json:"fdv"`
	Liquidity         go_format_type.Float64String `json:"liquidity"`
	TotalSupply       go_format_type.Float64String `json:"totalSupply"`
	CirculatingSupply go_format_type.Float64String `json:"circulatingSupply"`
	Holders           go_format_type.Int64String   `json:"holders"`
	Decimals          int                          `json:"decimals"`
	ListingCex        bool                         `json:"listingCex"`
	HotTag            bool                         `json:"hotTag"`
	CexCoinName       string                       `json:"cexCoinName"`
	CanTransfer       bool                         `json:"canTransfer"`
	Denomination      int                          `json:"denomination"`
	Offline           bool                         `json:"offline"`
	TradeDecimal      int                          `json:"tradeDecimal"`
	AlphaId           string                       `json:"alphaId"`
	Offsell           bool                         `json:"offsell"`
	PriceHigh24h      go_format_type.Float64String `json:"priceHigh24h"`
	PriceLow24h       go_format_type.Float64String `json:"priceLow24h"`
	Count24h          go_format_type.Int64String   `json:"count24h"`
	OnlineTge         bool                         `json:"onlineTge"`
	OnlineAirdrop     bool                         `json:"onlineAirdrop"`
	Score             int                          `json:"score"`
	CexOffDisplay     bool                         `json:"cexOffDisplay"`
	StockState        bool                         `json:"stockState"`
	ListingTime       int64                        `json:"listingTime"`
	MulPoint          int                          `json:"mulPoint"`
	BnExclusiveState  bool                         `json:"bnExclusiveState"`
}

func (t *Client) ListTokens() ([]*TokenInfoType, error) {
	var httpResult struct {
		Code    string           `json:"code"`
		Msg     string           `json:"message"`
		Data    []*TokenInfoType `json:"data"`
		Success bool             `json:"success"`
	}
	_, _, err := go_http.HttpInstance.GetForStruct(
		t.logger,
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
