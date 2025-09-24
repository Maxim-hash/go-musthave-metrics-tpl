package main

import (
	"net/http"

	handlers "github.com/Maxim-hash/go-musthave-metrics-tpl/internal/handler"
	models "github.com/Maxim-hash/go-musthave-metrics-tpl/internal/model"
)

func main() {
	storage := models.NewMemStorage()
	mux := http.NewServeMux()

	mux.HandleFunc("/update/", handlers.UpdateHandler(storage))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
