package uni_account

import (
	"context"
	"encoding/json"
	"net/http"
)

// CreateUMOrderService create order
type CreateUMOrderService struct {
	c                *Client
	symbol           string
	side             SideType
	positionSide     *PositionSideType
	orderType        OrderType
	timeInForce      *TimeInForceType
	quantity         string
	reduceOnly       *bool
	price            *string
	newClientOrderID *string
	stopPrice        *string
	workingType      *WorkingType
	activationPrice  *string
	callbackRate     *string
	priceProtect     *bool
	newOrderRespType NewOrderRespType
	closePosition    *bool
}

// Symbol set symbol
func (s *CreateUMOrderService) Symbol(symbol string) *CreateUMOrderService {
	s.symbol = symbol
	return s
}

// Side set side
func (s *CreateUMOrderService) Side(side SideType) *CreateUMOrderService {
	s.side = side
	return s
}

// PositionSide set side
func (s *CreateUMOrderService) PositionSide(positionSide PositionSideType) *CreateUMOrderService {
	s.positionSide = &positionSide
	return s
}

// Type set type
func (s *CreateUMOrderService) Type(orderType OrderType) *CreateUMOrderService {
	s.orderType = orderType
	return s
}

// TimeInForce set timeInForce
func (s *CreateUMOrderService) TimeInForce(timeInForce TimeInForceType) *CreateUMOrderService {
	s.timeInForce = &timeInForce
	return s
}

// Quantity set quantity of coin, not usdt
func (s *CreateUMOrderService) Quantity(quantity string) *CreateUMOrderService {
	s.quantity = quantity
	return s
}

// ReduceOnly set reduceOnly
func (s *CreateUMOrderService) ReduceOnly(reduceOnly bool) *CreateUMOrderService {
	s.reduceOnly = &reduceOnly
	return s
}

// Price set price
func (s *CreateUMOrderService) Price(price string) *CreateUMOrderService {
	s.price = &price
	return s
}

// NewClientOrderID set newClientOrderID
func (s *CreateUMOrderService) NewClientOrderID(newClientOrderID string) *CreateUMOrderService {
	s.newClientOrderID = &newClientOrderID
	return s
}

// StopPrice set stopPrice
func (s *CreateUMOrderService) StopPrice(stopPrice string) *CreateUMOrderService {
	s.stopPrice = &stopPrice
	return s
}

// WorkingType set workingType
func (s *CreateUMOrderService) WorkingType(workingType WorkingType) *CreateUMOrderService {
	s.workingType = &workingType
	return s
}

// ActivationPrice set activationPrice
func (s *CreateUMOrderService) ActivationPrice(activationPrice string) *CreateUMOrderService {
	s.activationPrice = &activationPrice
	return s
}

// CallbackRate set callbackRate
func (s *CreateUMOrderService) CallbackRate(callbackRate string) *CreateUMOrderService {
	s.callbackRate = &callbackRate
	return s
}

// PriceProtect set priceProtect
func (s *CreateUMOrderService) PriceProtect(priceProtect bool) *CreateUMOrderService {
	s.priceProtect = &priceProtect
	return s
}

// NewOrderResponseType set newOrderResponseType
func (s *CreateUMOrderService) NewOrderResponseType(newOrderResponseType NewOrderRespType) *CreateUMOrderService {
	s.newOrderRespType = newOrderResponseType
	return s
}

// ClosePosition set closePosition
func (s *CreateUMOrderService) ClosePosition(closePosition bool) *CreateUMOrderService {
	s.closePosition = &closePosition
	return s
}

func (s *CreateUMOrderService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, header *http.Header, err error) {

	r := &request{
		method:   http.MethodPost,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol":   s.symbol,
		"side":     s.side,
		"type":     s.orderType,
		"quantity": s.quantity,
	}
	if s.positionSide != nil {
		m["positionSide"] = *s.positionSide
	}
	if s.timeInForce != nil {
		m["timeInForce"] = *s.timeInForce
	}
	if s.reduceOnly != nil {
		m["reduceOnly"] = *s.reduceOnly
	}
	if s.price != nil {
		m["price"] = *s.price
	}
	if s.newClientOrderID != nil {
		m["newClientOrderId"] = *s.newClientOrderID
	}
	if s.stopPrice != nil {
		m["stopPrice"] = *s.stopPrice
	}
	if s.workingType != nil {
		m["workingType"] = *s.workingType
	}
	if s.priceProtect != nil {
		m["priceProtect"] = *s.priceProtect
	}
	if s.activationPrice != nil {
		m["activationPrice"] = *s.activationPrice
	}
	if s.callbackRate != nil {
		m["callbackRate"] = *s.callbackRate
	}
	if s.closePosition != nil {
		m["closePosition"] = *s.closePosition
	}
	r.setFormParams(m)
	data, header, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, &http.Header{}, err
	}
	return data, header, nil
}

// Do send request
func (s *CreateUMOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CreateUMOrderResponse, err error) {
	data, header, err := s.createOrder(ctx, "/papi/v1/um/order", opts...)
	if err != nil {
		return nil, err
	}
	res = new(CreateUMOrderResponse)
	err = json.Unmarshal(data, res)
	res.RateLimitOrder10s = header.Get("X-Mbx-Order-Count-10s")
	res.RateLimitOrder1m = header.Get("X-Mbx-Order-Count-1m")

	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUMOrderResponse define create order response
type CreateUMOrderResponse struct {
	Symbol            string           `json:"symbol"`
	OrderID           int64            `json:"orderId"`
	ClientOrderID     string           `json:"clientOrderId"`
	Price             string           `json:"price"`
	OrigQuantity      string           `json:"origQty"`
	ExecutedQuantity  string           `json:"executedQty"`
	CumQuote          string           `json:"cumQuote"`
	ReduceOnly        bool             `json:"reduceOnly"`
	Status            OrderStatusType  `json:"status"`
	StopPrice         string           `json:"stopPrice"`
	TimeInForce       TimeInForceType  `json:"timeInForce"`
	Type              OrderType        `json:"type"`
	Side              SideType         `json:"side"`
	UpdateTime        int64            `json:"updateTime"`
	WorkingType       WorkingType      `json:"workingType"`
	ActivatePrice     string           `json:"activatePrice"`
	PriceRate         string           `json:"priceRate"`
	AvgPrice          string           `json:"avgPrice"`
	PositionSide      PositionSideType `json:"positionSide"`
	ClosePosition     bool             `json:"closePosition"`
	PriceProtect      bool             `json:"priceProtect"`
	RateLimitOrder10s string           `json:"rateLimitOrder10s,omitempty"`
	RateLimitOrder1m  string           `json:"rateLimitOrder1m,omitempty"`
}

// GetUMOrderService get an order
type GetUMOrderService struct {
	c                 *Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
}

// Symbol set symbol
func (s *GetUMOrderService) Symbol(symbol string) *GetUMOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *GetUMOrderService) OrderID(orderID int64) *GetUMOrderService {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID set origClientOrderID
func (s *GetUMOrderService) OrigClientOrderID(origClientOrderID string) *GetUMOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// Do send request
func (s *GetUMOrderService) Do(ctx context.Context, opts ...RequestOption) (res *Order, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/order",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.setParam("origClientOrderId", *s.origClientOrderID)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(Order)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Order define order info
type Order struct {
	Symbol           string           `json:"symbol"`
	OrderID          int64            `json:"orderId"`
	ClientOrderID    string           `json:"clientOrderId"`
	Price            string           `json:"price"`
	ReduceOnly       bool             `json:"reduceOnly"`
	OrigQuantity     string           `json:"origQty"`
	ExecutedQuantity string           `json:"executedQty"`
	CumQuantity      string           `json:"cumQty"`
	CumQuote         string           `json:"cumQuote"`
	Status           OrderStatusType  `json:"status"`
	TimeInForce      TimeInForceType  `json:"timeInForce"`
	Type             OrderType        `json:"type"`
	Side             SideType         `json:"side"`
	StopPrice        string           `json:"stopPrice"`
	Time             int64            `json:"time"`
	UpdateTime       int64            `json:"updateTime"`
	WorkingType      WorkingType      `json:"workingType"`
	ActivatePrice    string           `json:"activatePrice"`
	PriceRate        string           `json:"priceRate"`
	AvgPrice         string           `json:"avgPrice"`
	OrigType         string           `json:"origType"`
	PositionSide     PositionSideType `json:"positionSide"`
	PriceProtect     bool             `json:"priceProtect"`
	ClosePosition    bool             `json:"closePosition"`
}
