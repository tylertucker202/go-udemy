package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func ResponseTimeMiddleware(next http.Handler) http.Handler {
	fmt.Println("Response Time Middleware initialized")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Start measuring the response time
		start := time.Now()

		wrappedWriter := &responseWriter{ResponseWriter: w, status: http.StatusOK}

		duration := time.Since(start)
		w.Header().Set("X-Response-Time", duration.String())
		next.ServeHTTP(wrappedWriter, r)
		//calc duration
		//log the request details

		fmt.Printf("Method: %s, Path: %s, Status: %d, Duration: %s\n", r.Method, r.URL.Path, wrappedWriter.status, duration.String())
		fmt.Println("Sent Response from Response Time Middleware")
	})
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) writeHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}
