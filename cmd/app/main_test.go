package main

import (
	"net/http"
	"net/http/httptest"
	"os"
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

func TestServesAtRootWhenUnset(t *testing.T) {
	os.Unsetenv("BASE_PATH")
	if got := get(t, "/health"); got != http.StatusOK {
		t.Fatalf("GET /health = %d, want 200", got)
	}
}

func TestServesUnderPrefix(t *testing.T) {
	t.Setenv("BASE_PATH", "/direct/agent-7:3000")
	if got := get(t, "/direct/agent-7:3000/health"); got != http.StatusOK {
		t.Fatalf("prefixed health = %d, want 200", got)
	}
	if got := get(t, "/health"); got != http.StatusNotFound {
		t.Fatalf("bare health = %d, want 404", got)
	}
}

func TestPrefixIsNormalised(t *testing.T) {
	t.Setenv("BASE_PATH", "direct/agent-7:3000/")
	if got := basePath(); got != "/direct/agent-7:3000" {
		t.Fatalf("basePath() = %q", got)
	}
}
