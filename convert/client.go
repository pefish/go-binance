package convert

// 闪兑

import (
	"encoding/json"
	"fmt"

	go_http "github.com/pefish/go-http"
	i_logger "github.com/pefish/go-interface/i-logger"
	"github.com/pkg/errors"
)

type Client struct {
	host      string
	logger    i_logger.ILogger
	apiKey    string
	secretKey string
}

func New(
	logger i_logger.ILogger,
	apiKey string,
	secretKey string,
) *Client {
	return &Client{
		host:      "https://api.binance.com",
		logger:    logger,
		apiKey:    apiKey,
		secretKey: secretKey,
	}
}

type SupportedPairsResultType struct {
	FromAsset       string  `json:"fromAsset"`
	ToAsset         string  `json:"toAsset"`
	FromAssetMinAmt float64 `json:"fromAssetMinAmount,string"`
	FromAssetMaxAmt float64 `json:"fromAssetMaxAmount,string"`
	ToAssetMinAmt   float64 `json:"toAssetMinAmount,string"`
	ToAssetMaxAmt   float64 `json:"toAssetMaxAmount,string"`
}

// 用户应当 fromAsset 和 toAsset 参数至少填一个
type SupportedPairsParamType struct {
	FromAsset string `json:"fromAsset"`
	ToAsset   string `json:"toAsset"`
}

func (t *Client) SupportedPairs(params *SupportedPairsParamType) ([]*SupportedPairsResultType, error) {
	if params.FromAsset == "" && params.ToAsset == "" {
		return nil, errors.New("fromAsset and toAsset can not both be empty")
	}
	queries := make(map[string]string)
	if params.FromAsset != "" {
		queries["fromAsset"] = params.FromAsset
	}
	if params.ToAsset != "" {
		queries["toAsset"] = params.ToAsset
	}
	var httpResult struct {
		Code int64                       `json:"code"`
		Msg  string                      `json:"msg"`
		Data []*SupportedPairsResultType `json:"data"`
	}
	_, bodyBytes, err := go_http.HttpInstance.Get(
		t.logger,
		&go_http.RequestParams{
			Url:     fmt.Sprintf("%s/sapi/v1/convert/exchangeInfo", t.host),
			Queries: queries,
		},
	)
	if err != nil {
		return nil, err
	}
	if bodyBytes[0] == '[' {
		err = json.Unmarshal(bodyBytes, &httpResult.Data)
		if err != nil {
			return nil, err
		}
		return httpResult.Data, nil
	}
	err = json.Unmarshal(bodyBytes, &httpResult)
	if err != nil {
		return nil, err
	}
	return nil, errors.New(httpResult.Msg)
}

type GetQuoteResultType struct {
	QuoteID        string  `json:"quoteId"`
	Ratio          float64 `json:"ratio,string"`
	InverseRatio   float64 `json:"inverseRatio,string"`
	ValidTimestamp int64   `json:"validTimestamp"`
	ToAmount       float64 `json:"toAmount,string"`
	FromAmount     float64 `json:"fromAmount,string"`
}

// fromAmount 或者 toAmount 只需要提供其中一个
type GetQuoteParamType struct {
	FromAsset  string  `json:"fromAsset"`            // 必填
	ToAsset    string  `json:"toAsset"`              // 必填
	FromAmount float64 `json:"fromAmount,omitempty"` // 成交后将被扣除的金额
	ToAmount   float64 `json:"toAmount,omitempty"`   // 成交后将会获得的金额
	WalletType string  `json:"walletType,omitempty"` // 选择支付钱包，可支持的钱包的选择有SPOT，FUNDING和EARN。组合钱包选择也可支持，如SPOT_FUNDING，FUNDING_EARN，SPOT_FUNDING_EARN或者SPOT_EARN。默认选择为SPOT
}

func (t *Client) GetQuote(params *GetQuoteParamType) (*GetQuoteResultType, error) {
	if params.FromAsset == "" ||
		params.ToAsset == "" {
		return nil, errors.New("fromAsset, toAsset and timestamp are required")
	}
	if params.FromAmount == 0 && params.ToAmount == 0 {
		return nil, errors.New("fromAmount and toAmount can not both be zero")
	}

	requestParams, err := SignRequest(
		&go_http.RequestParams{
			Url:    fmt.Sprintf("%s/sapi/v1/convert/getQuote", t.host),
			Params: params,
		},
		t.apiKey,
		t.secretKey,
	)
	if err != nil {
		return nil, err
	}
	var httpResult struct {
		Code int64               `json:"code"`
		Msg  string              `json:"msg"`
		Data *GetQuoteResultType `json:"data"`
	}
	_, bodyBytes, err := go_http.HttpInstance.PostFormUrlEncoded(
		t.logger,
		requestParams,
	)
	if err != nil {
		return nil, err
	}
	if bodyBytes[0] == '[' {
		err = json.Unmarshal(bodyBytes, &httpResult.Data)
		if err != nil {
			return nil, err
		}
		return httpResult.Data, nil
	}
	err = json.Unmarshal(bodyBytes, &httpResult)
	if err != nil {
		return nil, err
	}
	return nil, errors.New(httpResult.Msg)
}
