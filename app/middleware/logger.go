package middleware

import (
	"bytes"
	"github.com/Dubbril/my-gin-project/app/helper"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"strconv"
	"time"
)

var CorrelationID string

func InitLogger() {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339Nano}).
		Level(zerolog.TraceLevel).
		With().
		Str("1_SERVICE", "MY_GIN_SERVICE").
		Timestamp().
		Logger()
	zerolog.TimeFieldFormat = "2006-01-02T15:04:05.999Z07:00"
	log.Logger = logger
}

// LogHandler is a middleware function to log information about incoming requests and outgoing responses.
func LogHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Capture request details
		startTime := time.Now()
		var requestBytes []byte
		if c.Request.Body != nil {
			requestBytes, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBytes)) // Reset the request body for further use
		}

		// Get the value of the X-Correlation-ID header
		CorrelationID = c.GetHeader("X-Correlation-ID")
		buildLogReq := log.Info().
			Str("2_CORRELATION_ID", CorrelationID).
			Str("3_METHOD", c.Request.Method).
			Str("4_URL", c.Request.URL.RequestURI()).
			Str("5_CLIENT_IP", c.ClientIP())

		// Handle response body plaintext on json
		if helper.IsValidJSON(requestBytes) {
			buildLogReq.RawJSON("6_BODY", requestBytes)
		} else {
			buildLogReq.Str("6_BODY", string(requestBytes))
		}

		buildLogReq.Msg("LOG_STEP_1")

		// Create a custom response writer
		w := &responseLogger{body: bytes.NewBuffer(nil), ResponseWriter: c.Writer}
		// Continue with processing the request
		c.Writer = w
		c.Next()

		// Capture response details
		responseBytes := w.body.Bytes()
		responseStatus := w.status
		duration := time.Since(startTime)

		buildLogResp := log.Info().
			Str("2_CORRELATION_ID", CorrelationID).
			Str("3_RESPONSE_STATUS", strconv.Itoa(responseStatus)).
			Str("4_FULL_REQUEST_TIME", duration.String())

		// Handle response body plaintext on json
		if helper.IsValidJSON(responseBytes) {
			buildLogResp.RawJSON("5_BODY", responseBytes)
		} else {
			buildLogResp.Str("5_BODY", string(responseBytes))
		}

		buildLogResp.Msg("LOG_STEP_4")

		CorrelationID = ""
	}
}

// responseLogger is a custom response writer to capture the response body
type responseLogger struct {
	gin.ResponseWriter
	body   *bytes.Buffer
	status int
}

// Write is called to write the response body
func (w *responseLogger) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// WriteHeader is called to set the response status code
func (w *responseLogger) WriteHeader(statusCode int) {
	w.status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// WriteString is a helper function to write a string to the response body
func (w *responseLogger) WriteString(s string) (int, error) {
	return io.WriteString(w, s)
}
