package middleware

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

// RequestResponseLogger logs details about the incoming request and outgoing response
func RequestResponseLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log request details
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("Failed to read request body:", err)
		}
		log.Printf("Request: %s %s\nBody: %s\n", r.Method, r.URL, string(bodyBytes))

		// Set the body back because it was consumed by io.ReadAll
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// Capture the response body by using a custom ResponseWriter
		responseBody := new(bytes.Buffer)
		writer := &responseWriter{ResponseWriter: w, body: responseBody}

		// Pass control to the next handler in the chain
		next.ServeHTTP(writer, r)

		// Log response details
		statusCode := writer.status
		if statusCode == 0 {
			statusCode = http.StatusOK
		}
		log.Printf("Response Status: %d\nResponse Body: %s\n", statusCode, responseBody.String())
	})
}

// Custom response writer to capture the response body
type responseWriter struct {
	http.ResponseWriter
	body   *bytes.Buffer
	status int
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.status = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *responseWriter) Write(data []byte) (int, error) {
	rw.body.Write(data)                   // Write response data to buffer
	return rw.ResponseWriter.Write(data)  // Write response to the client
}
