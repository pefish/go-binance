package util

import (
	"fmt"
	"time"

	go_http "github.com/pefish/go-http"
	i_logger "github.com/pefish/go-interface/i-logger"
)

type CirculatingSupplyInfoType struct {
	UnlockedAmount  float64 `json:"unlockedAmount"`  // 流通数量
	LockedAmount    float64 `json:"lockedAmount"`    // 锁定数量
	UntrackedAmount float64 `json:"untrackedAmount"` // 未跟踪数量
	TotalAmount     float64 `json:"totalAmount"`     // 总量
}

func GetCirculatingSupplyInfo(logger i_logger.ILogger, currency string) (*CirculatingSupplyInfoType, error) {
	var httpResult struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Data    struct {
			Progress struct {
				UnlockedAmount  float64 `json:"unlockedAmount"`
				LockedAmount    float64 `json:"lockedAmount"`
				UntrackedAmount float64 `json:"untrackedAmount"`
			} `json:"progress"`
		} `json:"data"`
	}
	_, _, err := go_http.NewHttpRequester(go_http.WithLogger(logger), go_http.WithTimeout(10*time.Second)).GetForStruct(
		&go_http.RequestParams{
			Url: "https://www.binance.com/bapi/apex/v1/public/apex/marketing/token-unlock/detail",
			Queries: map[string]string{
				"symbol": currency,
			},
		},
		&httpResult,
	)
	if err != nil {
		return nil, err
	}
	if httpResult.Code != "000000" {
		return nil, fmt.Errorf("GetCirculatingSupplyInfo error. code: %s, message: %s", httpResult.Code, httpResult.Message)
	}
	return &CirculatingSupplyInfoType{
		UnlockedAmount:  httpResult.Data.Progress.UnlockedAmount,
		LockedAmount:    httpResult.Data.Progress.LockedAmount,
		UntrackedAmount: httpResult.Data.Progress.UntrackedAmount,
		TotalAmount:     httpResult.Data.Progress.UnlockedAmount + httpResult.Data.Progress.LockedAmount + httpResult.Data.Progress.UntrackedAmount,
	}, nil
}
