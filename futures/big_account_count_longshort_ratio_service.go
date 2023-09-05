package futures

import (
	"context"
	"encoding/json"
	"net/http"
)

type BigAccountCountLongshortRatioService struct {
	c         *Client
	symbol    string
	period    string
	limit     *int
	startTime *int64
	endTime   *int64
}

// Symbol set symbol
func (s *BigAccountCountLongshortRatioService) Symbol(symbol string) *BigAccountCountLongshortRatioService {
	s.symbol = symbol
	return s
}

// Period set period interval
func (s *BigAccountCountLongshortRatioService) Period(period string) *BigAccountCountLongshortRatioService {
	s.period = period
	return s
}

// Limit set limit
func (s *BigAccountCountLongshortRatioService) Limit(limit int) *BigAccountCountLongshortRatioService {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *BigAccountCountLongshortRatioService) StartTime(startTime int64) *BigAccountCountLongshortRatioService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *BigAccountCountLongshortRatioService) EndTime(endTime int64) *BigAccountCountLongshortRatioService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *BigAccountCountLongshortRatioService) Do(ctx context.Context, opts ...RequestOption) (res []*BigAccountCountLongshortRatio, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/futures/data/topLongShortAccountRatio",
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
		return []*BigAccountCountLongshortRatio{}, err
	}

	res = make([]*BigAccountCountLongshortRatio, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*BigAccountCountLongshortRatio{}, err
	}

	return res, nil
}

type BigAccountCountLongshortRatio struct {
	Symbol         string `json:"symbol"`
	LongShortRatio string `json:"longShortRatio"`
	LongAccount    string `json:"longAccount"`
	ShortAccount   string `json:"shortAccount"`
	Timestamp      int64  `json:"timestamp"`
}
