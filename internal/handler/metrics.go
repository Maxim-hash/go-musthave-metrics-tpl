package handler

import (
	"fmt"
	"net/http"
	"strconv"

	models "github.com/Maxim-hash/go-musthave-metrics-tpl/internal/model"
	"github.com/go-chi/chi/v5"
)

func GetMetricsHandler(storage *models.MemStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		metrics := storage.GetAllMetrics()
		var str string

		for name, value := range metrics {
			str += name + ": " + fmt.Sprintf("%v", value) + "\n"
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(str))
	}
}

func GetMetricValueHandler(storage *models.MemStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		metricType := chi.URLParam(r, "metricType")
		metricName := chi.URLParam(r, "metricName")

		switch metricType {
		case "counter":
			val, ok := storage.GetCounter(metricName)
			if !ok {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(strconv.FormatInt(val, 10)))
		case "gauge":
			val, ok := storage.GetGauge(metricName)
			if !ok {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(strconv.FormatFloat(val, 'f', -1, 64)))
		default:
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
	}
}
