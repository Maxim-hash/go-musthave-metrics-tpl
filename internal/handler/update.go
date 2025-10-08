package handler

import (
	"net/http"
	"strconv"

	models "github.com/Maxim-hash/go-musthave-metrics-tpl/internal/model"
	"github.com/go-chi/chi/v5"
)

func UpdateHandler(storage *models.MemStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		metricType := chi.URLParam(r, "metricType")
		metricName := chi.URLParam(r, "metricName")
		metricValue := chi.URLParam(r, "metricValue")

		if metricName == "" {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		switch metricType {
		case "counter":
			val, err := strconv.ParseInt(metricValue, 10, 64)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			storage.UpdateCounter(metricName, val)
		case "gauge":
			val, err := strconv.ParseFloat(metricValue, 64)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			storage.UpdateGauge(metricName, (val))
		default:
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	}
}
