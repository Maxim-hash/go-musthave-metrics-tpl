package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.com/Maxim-hash/go-musthave-metrics-tpl/internal/config/flags"
	handlers "github.com/Maxim-hash/go-musthave-metrics-tpl/internal/handler"
	"github.com/Maxim-hash/go-musthave-metrics-tpl/internal/logger"
	"github.com/Maxim-hash/go-musthave-metrics-tpl/internal/middleware"
	models "github.com/Maxim-hash/go-musthave-metrics-tpl/internal/model"
)

func main() {
	cfg := flags.ParseFlags()

	if err := run(cfg); err != nil {
		panic(err)
	}
}

func run(cfg *flags.Config) error {
	if err := logger.NewLogger(cfg.FlagLogLevel); err != nil {
		return err
	}
	logger.Log.Info("Running server", zap.String("address", cfg.FlagRunAddr))
	r := chi.NewRouter()
	r.Use(middleware.WithLogging)
	storage := models.NewMemStorage()

	r.Post("/update/{metricType}/{metricName}/{metricValue}", handlers.UpdateHandler(storage))

	r.Route("/", func(r chi.Router) {
		r.Get("/", handlers.GetMetricsHandler(storage))
		r.Route("/value/{metricType}/{metricName}", func(r chi.Router) {
			r.Get("/", handlers.GetMetricValueHandler(storage))
		})
	})

	return http.ListenAndServe(cfg.FlagRunAddr, r)
}
