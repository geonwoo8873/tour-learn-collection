package pkgdata

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type pkgData struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

func PackageData(url string) ([]pkgData, error) {
	var packages []pkgData

	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	if req.Header.Get("Content-Type") != "application/json" {
		return nil, fmt.Errorf("unexpected content type: %s", req.Header.Get("Content-Type"))
	}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &packages)
	if err != nil {
		return nil, err
	}

	return packages, nil
}
