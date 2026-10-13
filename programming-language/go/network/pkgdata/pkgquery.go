package pkgdata

import (
	"encoding/json"
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
