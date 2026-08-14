package middlewares

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"strings"
)

func CompressionMiddleware(next http.Handler) http.Handler {
	fmt.Println("Compression Middleware Initialized")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Compression Middleware Executed")
		// Start measuring the response time
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		// Set the response headers for gzip compression
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		w = &gzipResponseWriter{ResponseWriter: w, Writer: gz}
		defer gz.Close()
		w.Header().Set("Vary", "Accept-Encoding")
		next.ServeHTTP(w, r)
		fmt.Println("Sent response from Compression Middleware")
	})
}

type gzipResponseWriter struct {
	http.ResponseWriter
	Writer *gzip.Writer
}

func (g *gzipResponseWriter) Write(b []byte) (int, error) {
	return g.Writer.Write(b)
}
