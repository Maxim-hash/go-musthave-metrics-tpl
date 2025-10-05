// internal/handler/update_test.go
package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	handler "github.com/Maxim-hash/go-musthave-metrics-tpl/internal/handler"
	models "github.com/Maxim-hash/go-musthave-metrics-tpl/internal/model"
)

func TestUpdateHandler_OK_Gauge(t *testing.T) {
	st := models.NewMemStorage()
	h := handler.UpdateHandler(st)

	req := httptest.NewRequest(http.MethodPost, "/update/gauge/Alloc/123.5", nil)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("want 200, got %d", w.Code)
	}
	v, ok := st.GetGauge("Alloc")
	if !ok || v != 123.5 {
		t.Fatalf("storage not updated: v=%v ", v)
	}
}
