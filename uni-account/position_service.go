package uni_account

import (
	"context"
	"encoding/json"
	"net/http"
)

// ChangeUMLeverageService change user's initial leverage of specific symbol market
type ChangeUMLeverageService struct {
	c        *Client
	symbol   string
	leverage int
}

// Symbol set symbol
func (s *ChangeUMLeverageService) Symbol(symbol string) *ChangeUMLeverageService {
	s.symbol = symbol
	return s
}

// Leverage set leverage
func (s *ChangeUMLeverageService) Leverage(leverage int) *ChangeUMLeverageService {
	s.leverage = leverage
	return s
}

// Do send request
func (s *ChangeUMLeverageService) Do(ctx context.Context, opts ...RequestOption) (res *SymbolLeverage, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/papi/v1/um/leverage",
		secType:  secTypeSigned,
	}
	r.setFormParams(params{
		"symbol":   s.symbol,
		"leverage": s.leverage,
	})
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SymbolLeverage)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SymbolLeverage define leverage info of symbol
type SymbolLeverage struct {
	Leverage         int    `json:"leverage"`
	MaxNotionalValue string `json:"maxNotionalValue"`
	Symbol           string `json:"symbol"`
}
