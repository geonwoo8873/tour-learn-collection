package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func RemoteResource(url string) ([]byte, error) {
	// `Get`은 지정된 URL로 HTTP GET 요청을 하고 Response 타입 객체와 error 값을 반환한다.
	// Response 객체의 Body는 반드시 사용 후 닫아야 한다. 여기서는 `io.ReadAll`을 사용하여 Body를 읽고 반환한다.
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	// io 라이브러리의 `ReadAll`는 Content Body를 읽고 해당 함수의 반환 값을 그대로 반환한다.
	return io.ReadAll(res.Body)
}

func main() {
	// `Args`는 `Application Name`을 포함하여 `Command line` 인자를 담고 있는 Slice이다.
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stdout, "HTTP URL to get data from\nUsage: %s <url>\n", os.Args[0])
		os.Exit(1)
	}

	body, err := RemoteResource(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stdout, "%v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "%s\n", body)
}
