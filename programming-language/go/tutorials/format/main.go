package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := readFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(content)

	// `Printf`는 형식화된 문자열을 출력할 때 활용되고, 주로 터미널에서 결과를 확인할 때나 로그 및 안내에 사용된다.
	// `\n`은 Println과 유사한 줄바꿈을 수행하지만, Printf에서 문자열 내에 직접 포함시켜야 한다.
	fmt.Printf(
		"tour-learn-collection is a repository dedicated to recording the continuous learning undertaken to grow as a developer and engineer.\n" +
		"Beyond simply serving as a space for personal study records, it aims to foster a culture of sharing and disseminating diverse knowledge,\n" +
		"including programming languages and credentials. Building knowledge through learning,\n" +
		"sharing it, and striving for a better career that is the dream and direction of this project.\n")
}

// func {function_name}(variable_name datatype) return_datatype {...}
// `return`에 대한 데이터 타입을 명시해야 하며, 함수 내부에서 반드시 해당 타입의 값을 반환해야 한다.
// `os.ReadFile`는 지정된 파일을 읽어 바이트 슬라이스(`[]byte Slice`)로 반환하며, 오류(`nil = error`)가 발생할 경우 해당 오류를 반환한다.
func readFile(filename string) (string, error) {
	readBytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return "", err
	}
	s := string(readBytes)
	return s, nil
}