package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"openclaw/internal/store"
)

func TestCommerceReadEndpointsReturnSeedData(t *testing.T) {
	srv := NewServer(store.New(), nil, http.NotFoundHandler(), "http://agent.test")
	routes := srv.Routes()

	for _, path := range []string{"/api/products", "/api/orders", "/api/knowledge", "/api/mcp-connectors", "/api/ai-partners", "/api/skillhub", "/api/paperclip"} {
		t.Run(path, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, path, nil)
			rec := httptest.NewRecorder()

			routes.ServeHTTP(rec, req)

			if rec.Code != http.StatusOK {
				t.Fatalf("expected 200, got %d with %s", rec.Code, rec.Body.String())
			}
			if rec.Body.Len() == 0 {
				t.Fatal("expected response body")
			}
		})
	}
}
