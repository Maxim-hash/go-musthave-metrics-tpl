package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	handlers "github.com/Maxim-hash/go-musthave-metrics-tpl/internal/handler"
	models "github.com/Maxim-hash/go-musthave-metrics-tpl/internal/model"
)

func main() {
	r := chi.NewRouter()
	storage := models.NewMemStorage()

	r.Post("/update/{metricType}/{metricName}/{metricValue}", handlers.UpdateHandler(storage))

	r.Route("/", func(r chi.Router) {
		r.Get("/", handlers.GetMetricsHandler(storage))
		r.Route("/value/{metricType}/{metricName}", func(r chi.Router) {
			r.Get("/", handlers.GetMetricValueHandler(storage))
		})
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
