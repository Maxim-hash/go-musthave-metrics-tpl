package handler

import (
	"net/http"
	"strconv"
	"strings"

	models "github.com/Maxim-hash/go-musthave-metrics-tpl/internal/model"
)

func UpdateHandler(storage *models.MemStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		path := strings.TrimPrefix(r.URL.Path, "/update/")
		parts := strings.Split(path, "/")

		if len(parts) != 3 {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		metricType, metricName, metricValue := parts[0], parts[1], parts[2]

		if metricName == "" {
			http.Error(w, "bad request", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		switch metricType {
		case "counter":
			val, err := strconv.ParseInt(metricValue, 10, 64)
			if err != nil {
				http.Error(w, "bad request", http.StatusBadRequest)
				return
			}
			storage.UpdateCounter(metricName, val)
		case "gauge":
			val, err := strconv.ParseFloat(metricValue, 64)
			if err != nil {
				http.Error(w, "bad request", http.StatusBadRequest)
				return
			}
			storage.UpdateGauge(metricName, (val))
		default:
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
	}
}
