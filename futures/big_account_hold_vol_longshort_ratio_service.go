package futures

import (
	"context"
	"encoding/json"
	"net/http"
)

type BigAccountHoldVolLongshortRatioService struct {
	c         *Client
	symbol    string
	period    string
	limit     *int
	startTime *int64
	endTime   *int64
}

// Symbol set symbol
func (s *BigAccountHoldVolLongshortRatioService) Symbol(symbol string) *BigAccountHoldVolLongshortRatioService {
	s.symbol = symbol
	return s
}

// Period set period interval
func (s *BigAccountHoldVolLongshortRatioService) Period(period string) *BigAccountHoldVolLongshortRatioService {
	s.period = period
	return s
}

// Limit set limit
func (s *BigAccountHoldVolLongshortRatioService) Limit(limit int) *BigAccountHoldVolLongshortRatioService {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *BigAccountHoldVolLongshortRatioService) StartTime(startTime int64) *BigAccountHoldVolLongshortRatioService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *BigAccountHoldVolLongshortRatioService) EndTime(endTime int64) *BigAccountHoldVolLongshortRatioService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *BigAccountHoldVolLongshortRatioService) Do(ctx context.Context, opts ...RequestOption) (res []*BigAccountHoldVolLongshortRatio, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/futures/data/topLongShortPositionRatio",
	}

	r.setParam("symbol", s.symbol)
	r.setParam("period", s.period)

	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*BigAccountHoldVolLongshortRatio{}, err
	}

	res = make([]*BigAccountHoldVolLongshortRatio, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*BigAccountHoldVolLongshortRatio{}, err
	}

	return res, nil
}

type BigAccountHoldVolLongshortRatio struct {
	Symbol         string `json:"symbol"`
	LongShortRatio string `json:"longShortRatio"`
	LongAccount    string `json:"longAccount"`
	ShortAccount   string `json:"shortAccount"`
	Timestamp      int64  `json:"timestamp"`
}
