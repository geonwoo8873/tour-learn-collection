package postnetwork

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type pkgData struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type registerResult struct {
	Id string `json:"id"`
}

func registerPackage(url string, data pkgData) (registerResult, error) {
	pkg := registerResult{}

	// `Marshal`는 요청한 Body에 JSON 데이터로 전송할 객체를 Byte Slice로 변환시킨다.
	body, err := json.Marshal(data)
	if err != nil {
		return pkg, err
	}

	reader := bytes.NewReader(body)
	r, err := http.Post(url, "application/json", reader)
	if err != nil {
		return pkg, err
	}

	defer r.Body.Close()

	resData, err := io.ReadAll(r.Body)
	if err != nil {
		return pkg, err
	}

	err = json.Unmarshal(resData, &pkg)
	if err != nil {
		return pkg, errors.New(string(resData))
	}

	return pkg, nil
}
