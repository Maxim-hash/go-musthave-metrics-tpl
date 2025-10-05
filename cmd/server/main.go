package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	handlers "github.com/Maxim-hash/go-musthave-metrics-tpl/internal/handler"
	models "github.com/Maxim-hash/go-musthave-metrics-tpl/internal/model"
)

func main() {
	ParseFlags()

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	r := chi.NewRouter()
	storage := models.NewMemStorage()

	r.Post("/update/{metricType}/{metricName}/{metricValue}", handlers.UpdateHandler(storage))

	r.Route("/", func(r chi.Router) {
		r.Get("/", handlers.GetMetricsHandler(storage))
		r.Route("/value/{metricType}/{metricName}", func(r chi.Router) {
			r.Get("/", handlers.GetMetricValueHandler(storage))
		})
	})
	fmt.Println("Server running on ", flagRunAddr)
	return http.ListenAndServe(flagRunAddr, r)
}
