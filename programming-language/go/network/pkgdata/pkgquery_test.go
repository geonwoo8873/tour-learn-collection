package pkgdata

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func packageServer() *httptest.Server {
	pkgData := `[
		{
			"name": "example-package",
			"version": "1.0.0",
			"description": "An example package"
		}
	]`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		fmt.Fprint(w, pkgData)
	}))

	return ts
}

func TestPackageData(t *testing.T) {
	ts := packageServer()
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("Failed to get response: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected status OK, got %v", res.Status)
	}
}
