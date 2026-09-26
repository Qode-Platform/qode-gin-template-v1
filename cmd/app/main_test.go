package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func get(t *testing.T, path string) int {
	t.Helper()
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	newRouter().ServeHTTP(w, httptest.NewRequest(http.MethodGet, path, nil))
	return w.Code
}

func TestServesAtRoot(t *testing.T) {
	for _, path := range []string{"/health", "/"} {
		if got := get(t, path); got != http.StatusOK {
			t.Fatalf("GET %s = %d, want 200", path, got)
		}
	}
}
