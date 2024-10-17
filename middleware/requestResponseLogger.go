package middleware

import (
	"bytes"
	"io"
	"log"

	"github.com/gin-gonic/gin"
)

// Custom middleware to log request and response
func RequestResponseLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log request details
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Println("Failed to read request body:", err)
		}
		log.Printf("Request: %s %s\nBody: %s\n", c.Request.Method, c.Request.URL, string(bodyBytes))

		// Set the body back because it was consumed by io.ReadAll
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// Capture the response body
		responseBody := new(bytes.Buffer)
		writer := c.Writer
		c.Writer = &responseWriter{ResponseWriter: writer, body: responseBody}

		// Process the request
		c.Next()

		// Log response details
		statusCode := c.Writer.Status()
		log.Printf("Response Status: %d\nResponse Body: %s\n", statusCode, responseBody.String())
	}
}

// Custom response writer to capture the response body
type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (rw *responseWriter) Write(data []byte) (int, error) {
	rw.body.Write(data)                  // Write response data to buffer
	return rw.ResponseWriter.Write(data) // Write response to the client
}
