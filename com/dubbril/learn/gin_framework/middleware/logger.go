package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

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

		// Create a custom response writer
		w := &responseLogger{body: bytes.NewBuffer(nil), ResponseWriter: c.Writer}

		// Get the value of the X-Correlation-ID header
		correlationID := c.GetHeader("X-Correlation-ID")

		// Continue with processing the request
		c.Writer = w
		c.Next()

		// Capture response details
		responseBytes := w.body.Bytes()
		responseStatus := w.status
		duration := time.Since(startTime)

		// Log information about the request
		var prettyRequestJSON bytes.Buffer
		_ = json.Compact(&prettyRequestJSON, requestBytes)

		fmt.Printf("[EDGE-USER-SERVICE] %s | Correlation-ID : %s | Client IP : %s | Method : %s | %s | Request Body: %s\n",
			startTime.Format("2006/01/02 - 15:04:05.000"),
			correlationID,
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL.Path,
			prettyRequestJSON.String(),
		)

		// Log information about the response
		var prettyResponseJSON bytes.Buffer
		_ = json.Compact(&prettyResponseJSON, responseBytes)

		fmt.Printf("[EDGE-USER-SERVICE] %s | Correlation-ID : %s | Response Status : %d | Full Request Time : %s | Response Body: %s\n",
			time.Now().Format("2006/01/02 - 15:04:05.000"),
			correlationID,
			responseStatus,
			duration,
			prettyResponseJSON.String(),
		)
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
