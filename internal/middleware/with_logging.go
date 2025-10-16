package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/Maxim-hash/go-musthave-metrics-tpl/internal/logger"
)

func WithLogging(h http.Handler) http.Handler {
	logFn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		responseData := &logger.ResponseData{
			Status: 0,
			Size:   0,
		}
		lw := logger.LoggingResponseWriter{
			ResponseWriter: w,            // встраиваем оригинальный http.ResponseWriter
			ResponseData:   responseData, // Change to the correct exported field name if available
		}
		h.ServeHTTP(&lw, r) // внедряем реализацию http.ResponseWriter

		duration := time.Since(start)

		logger.Log.Info(
			"request completed",
			zap.String("uri", r.RequestURI),
			zap.String("method", r.Method),
			zap.Int("status", responseData.Status), // получаем перехваченный код статуса ответа
			zap.Duration("duration", duration),
			zap.Int("size", responseData.Size), // получаем перехваченный размер ответа
		)
	}
	return http.HandlerFunc(logFn)
}
