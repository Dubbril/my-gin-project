package middleware

import (
	"bytes"
	"github.com/Dubbril/my-gin-project/app/config"
	"github.com/Dubbril/my-gin-project/app/helper"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"io"
	"strconv"
)

type RestClient struct {
	Client *resty.Client
}

func NewRestClient(client *resty.Client) *RestClient {
	getConfig := config.GetConfig()
	client.OnBeforeRequest(HandlerRequest)
	client.OnAfterResponse(HandlerResponse)
	client.SetTimeout(getConfig.ConnectionTimeout)
	return &RestClient{Client: client}
}

func HandlerRequest(client *resty.Client, request *resty.Request) error {
	client.Header.Set("x-correlation-id", CorrelationID)

	var requestBytes []byte
	if request.Body != nil {
		bodyReader, ok := request.Body.(io.ReadCloser)
		if ok {
			requestBytes, _ = io.ReadAll(bodyReader)
			request.Body = io.NopCloser(bytes.NewBuffer(requestBytes))
		}
	}

	buildLogReq := log.Info().
		Str("2_CORRELATION_ID", CorrelationID).
		Str("3_METHOD", request.Method).
		Str("4_URL", request.URL)

	// Handle response body plaintext on json
	if helper.IsValidJSON(requestBytes) {
		buildLogReq.RawJSON("5_BODY", requestBytes)
	} else {
		buildLogReq.Str("5_BODY", string(requestBytes))
	}

	buildLogReq.Msg("LOG_STEP_2")

	return nil
}

func HandlerResponse(_ *resty.Client, resp *resty.Response) error {
	buildLogResp := log.Info().
		Str("2_CORRELATION_ID", CorrelationID).
		Str("3_RESPONSE_STATUS", strconv.Itoa(resp.StatusCode())).
		Str("4_FULL_REQUEST_TIME", resp.Time().String())

	// Handle response body plaintext on json
	if helper.IsValidJSON(resp.Body()) {
		buildLogResp.RawJSON("5_BODY", resp.Body())
	} else {
		buildLogResp.Str("5_BODY", string(resp.Body()))
	}

	buildLogResp.Msg("LOG_STEP_3")
	return nil
}
