package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// `httptest`는 HTTP Server Status를 테스트하기 위한 패키지로,
// 실제 서버를 띄우지 않고도 HTTP 요청과 응답을 시뮬레이션할 수 있다.
func setHTTPServer() *httptest.Server {
	ts := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, "Hello World")
			}))

	return ts
}

// Test 함수명은 테스트 진행될 함수임을 나타내기 때문에, 다른 Test 함수가 존재하면 실행대상이 된다.
// 즉 Test**** 함수 네이밍은 진행할 테스트 케이스외에 붙이지 말아야된다.
func Test(t *testing.T) {
	ts := setHTTPServer()
	defer ts.Close()

	expected := "Hello World"
	data, err := RemoteResource(ts.URL)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if expected != string(data) {
		t.Errorf("Expected %q, got %q", expected, string(data))
	}
}

//// CMD
// go test -v
//// Result
// === RUN   TestRemoteResource
// --- PASS: TestRemoteResource (0.00s)
// PASS
