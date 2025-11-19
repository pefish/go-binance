package ws_service

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"encoding/hex"
	"encoding/json"

	"github.com/gorilla/websocket"
	i_logger "github.com/pefish/go-interface/i-logger"
	"github.com/pkg/errors"
)

const (
	baseWsMainURL = "wss://api.binance.com/sapi/wss"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandString(n int) string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}

func HmacSHA256(message, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}

var (
	// WebsocketTimeout is an interval for sending ping/pong messages if WebsocketKeepalive is enabled
	WebsocketTimeout = time.Second * 30
)

type WSService struct {
	apiKey    string
	apiSecret string
	logger    i_logger.ILogger
}

func New(apiKey, apiSecret string, logger i_logger.ILogger) *WSService {
	return &WSService{
		apiKey:    apiKey,
		apiSecret: apiSecret,
		logger:    logger,
	}
}

type WsAnnouncementEvent struct {
	CatalogId   int    `json:"catalogId"`
	CatalogName string `json:"catalogName"`
	PublishDate int64  `json:"publishDate"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	Disclaimer  string `json:"disclaimer"`
}

type WsAnnouncementHandler func(data *WsAnnouncementEvent)

type ResponseType struct {
	Type_   string `json:"type"`
	Data    string `json:"data"`
	SubType string `json:"subType"`
}

func (t *WSService) WatchAnnouncement(
	ctx context.Context,
	handler WsAnnouncementHandler,
) error {
	return WsServeLoop(
		ctx,
		t.logger,
		&WsConfig{
			EndpointFunc: func() string {
				paramsExceptSign := fmt.Sprintf(
					"random=%s&topic=%s&recvWindow=%s&timestamp=%d",
					RandString(32),
					"com_announcement_en",
					"30000",
					time.Now().UnixMilli(),
				)
				signature := HmacSHA256(paramsExceptSign, t.apiSecret)

				return fmt.Sprintf("%s?%s&signature=%s", baseWsMainURL, paramsExceptSign, signature)
			},
			ApiKey: t.apiKey,
		},
		func(data []byte) {
			event := new(WsAnnouncementEvent)
			err := json.Unmarshal(data, &event)
			if err != nil {
				t.logger.ErrorF("event Unmarshal error: %v", err)
				return
			}
			handler(event)
		},
	)
}

// WsHandler handle raw websocket message
type WsHandler func(message []byte)

// ErrHandler handles errors
type ErrHandler func(err error)

// WsConfig webservice configuration
type WsConfig struct {
	EndpointFunc func() string
	ApiKey       string
}

func WsServeLoop(
	ctx context.Context,
	logger i_logger.ILogger,
	cfg *WsConfig,
	handler WsHandler,
) error {
	url := cfg.EndpointFunc()
	wsServeChan := make(chan bool, 1)
	wsServeChan <- true
	var doneC chan struct{}
	var stopC chan struct{}
	var err error
	newCtx, cancel := context.WithCancel(ctx)
	for {
		select {
		case <-wsServeChan:
			logger.InfoF("Connecting <%s>...", url)
			doneC, stopC, err = WsServe(
				logger,
				cfg,
				handler,
				func(err error) {
					if strings.Contains(err.Error(), "connection timed out") ||
						strings.Contains(err.Error(), "i/o timeout") ||
						strings.Contains(err.Error(), "close 1006 (abnormal closure)") ||
						strings.Contains(err.Error(), "connection reset by peer") {
						logger.InfoF("Connection <%s> closed, reconnect.", url)
						wsServeChan <- true
					} else {
						// 如果是其他错误，直接中断
						logger.ErrorF("Connection <%s> error: %v", url, err)
						cancel()
					}
				},
			)
			if err != nil {
				return err
			}
			logger.InfoF("Connect <%s> done.", url)
		case <-doneC:
			logger.InfoF("Connection <%s> closed, to reconnect...", url)
			wsServeChan <- true
			doneC = nil // 阻止这个分支被多次执行
			continue
		case <-ctx.Done():
			stopC <- struct{}{}
			return nil
		case <-newCtx.Done():
			return nil
		}
	}
}

func WsServe(logger i_logger.ILogger, cfg *WsConfig, handler WsHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	Dialer := websocket.Dialer{
		Proxy:             http.ProxyFromEnvironment,
		HandshakeTimeout:  45 * time.Second,
		EnableCompression: false,
	}

	requestHeader := http.Header{}
	requestHeader.Add("X-MBX-APIKEY", cfg.ApiKey)
	c, _, err := Dialer.Dial(cfg.EndpointFunc(), requestHeader)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to dial websocket")
	}
	c.SetReadLimit(655350)
	doneC = make(chan struct{})
	stopC = make(chan struct{})
	go func() {
		// This function will exit either on error from
		// websocket.Conn.ReadMessage or when the stopC channel is
		// closed by the client.
		defer close(doneC)
		keepAlive(logger, c, WebsocketTimeout)
		// Wait for the stopC channel to be closed.  We do that in a
		// separate goroutine because ReadMessage is a blocking
		// operation.
		silent := false
		go func() {
			select {
			case <-stopC:
				logger.Debug("stopC received.")
				silent = true
			case <-doneC:
				logger.Debug("doneC received.")
			}
			c.Close()
			logger.Debug("connection closed.")
		}()
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				if !silent {
					errHandler(errors.Wrap(err, "failed to read websocket message"))
				}
				return
			}
			logger.DebugF("received msg: %s", string(message))
			response := new(ResponseType)
			err = json.Unmarshal(message, &response)
			if err != nil {
				if !silent {
					errHandler(errors.Errorf("response Unmarshal error. %+v", err))
				}
				return
			}
			if response.Type_ == "COMMAND" && response.SubType == "REGISTER" {
				if !silent {
					errHandler(errors.Errorf("response error. %+v", response))
				}
				return
			}
			handler([]byte(response.Data))
		}
	}()
	return
}

func keepAlive(logger i_logger.ILogger, c *websocket.Conn, timeout time.Duration) {
	ticker := time.NewTicker(timeout)

	c.SetPongHandler(func(msg string) error {
		logger.Debug("pong received.")
		c.SetReadDeadline(time.Now().Add(timeout * 2))
		return nil
	})

	go func() {
		defer ticker.Stop()
		for {
			deadline := time.Now().Add(10 * time.Second)
			err := c.WriteControl(websocket.PingMessage, []byte{}, deadline)
			if err != nil {
				logger.ErrorF("send ping error. err: %s", err.Error())
				return
			}
			logger.Debug("ping sended.")
			<-ticker.C
		}
	}()
}
