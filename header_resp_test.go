package traefikplugindemo

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDemo(t *testing.T) {
	cfg := CreateConfig()
	cfg.ValueStrCase = "snake"
	cfg.DefaultValue = "test demo"

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := New(ctx, next, cfg, "axiaoxin-plugin-demo")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	assertHeader(t, recorder, "resp", "test_demo")
}

func assertHeader(t *testing.T, recorder *httptest.ResponseRecorder, key, expected string) {
	t.Helper()

	if recorder.HeaderMap.Get(key) != expected {
		t.Errorf("invalid header value: %s", recorder.HeaderMap.Get(key))
	}
}
