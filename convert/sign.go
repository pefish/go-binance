package convert

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	go_format_any "github.com/pefish/go-format/any"
	go_http "github.com/pefish/go-http"
)

func SignRequest(
	params *go_http.RequestParams,
	apiKey string,
	secretKey string,
) (*go_http.RequestParams, error) {
	body := make(url.Values, 0)

	b, err := json.Marshal(params.Params)
	if err != nil {
		return nil, err
	}

	var m map[string]any
	err = json.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		body.Set(k, go_format_any.ToString(v))
	}
	body.Set("timestamp", go_format_any.ToString(time.Now().UnixMilli()))
	bodyString := body.Encode()

	mac := hmac.New(sha256.New, []byte(secretKey))
	_, err = mac.Write([]byte(bodyString))
	if err != nil {
		return nil, err
	}
	sig := fmt.Sprintf("%x", (mac.Sum(nil)))

	bodyString = body.Encode()
	bodyString += "&signature=" + sig

	params.Params = bodyString
	if params.Headers == nil {
		params.Headers = make(map[string]string, 0)
	}
	params.Headers["X-MBX-APIKEY"] = apiKey
	return params, nil
}
