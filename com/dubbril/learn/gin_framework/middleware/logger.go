package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"io"
	"os"
	"strconv"
	"time"
)

func logger() zerolog.Logger {
	return zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Logger()
}

// LogHandler is a middleware function to log information about incoming requests and outgoing responses.
func LogHandler() gin.HandlerFunc {
	logger := logger()
	return func(c *gin.Context) {
		// Capture request details
		startTime := time.Now()
		var requestBytes []byte
		if c.Request.Body != nil {
			requestBytes, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBytes)) // Reset the request body for further use
		}

		// Get the value of the X-Correlation-ID header
		correlationID := c.GetHeader("X-Correlation-ID")
		logger.Info().
			Str("SERVICE", "MY_GIN_SERVICE").
			Str("CORRELATION_ID", correlationID).
			Str("METHOD", c.Request.Method).
			Str("URL", c.Request.URL.RequestURI()).
			Str("CLIENT_IP", c.ClientIP()).
			Msg(string(requestBytes))

		// Create a custom response writer
		w := &responseLogger{body: bytes.NewBuffer(nil), ResponseWriter: c.Writer}
		// Continue with processing the request
		c.Writer = w
		c.Next()

		// Capture response details
		responseBytes := w.body.Bytes()
		responseStatus := w.status
		duration := time.Since(startTime)

		logger.Info().
			Str("SERVICE", "MY_GIN_SERVICE").
			Str("CORRELATION_ID", correlationID).
			Str("RESPONSE_STATUS", strconv.Itoa(responseStatus)).
			Str("FULL_REQUEST_TIME", duration.String()).
			Msg(string(responseBytes))
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
