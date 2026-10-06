package main

import (
	"fmt"
	// `bufio`는 버퍼링(Buffering)된 I/O를 제공하여 시스템 호출을 최소화하여 성능을 향상시키는 기능을 제공
	"bufio"
	// `errors`는 오류 처리를 위한 표준 라이브러리
	"errors"
	// `io`는 입출력(I/O) 관련 기능을 제공
	"io"
	// `strconv`는 문자열을 다른 기본 자료형으로 변환하는 기능을 제공
	"strconv"
	// `os`는 운영체제와 상호작용하기 위한 기능을 제공
	"os"
)

func printUsage(write io.Writer) {
	usage := "Usage: program <numTimes>\n"
	fmt.Fprint(write, usage)
}

func validateArgs(cfg config) error {
	if !(cfg.numTime > 0) {
		return errors.New("numTime must be greater than 0")
	}

	return nil
}

// Reader, Writer = 데이터 읽기 및 쓰기를 위한 인터페이스
func getName(read io.Reader, write io.Writer) (string, error) {
	msg := "Please enter your input: "
	fmt.Fprint(write, msg)

	if write != nil {
		fmt.Println("Debug: Writer is not nil.")
	}

	// 사용자 입력을 받기 위해 bufio.NewScanner를 사용하여 읽기 가능한 스캐너를 생성
	scanner := bufio.NewScanner(read)
	scanner.Scan()

	// 입력 받은 값이 없으면 공백과 함께 에러를 반환하도록 처리
	if err := scanner.Err(); err != nil {
		return "", err
	}

	// Text 입력 값을 받을 수 있도록 scanner.Text()를 사용
	getInput := scanner.Text()
	// 입력 받은 값의 길이가 0이면 공백과 함께 오류 메시지가 반환되도록 처리
	if len(getInput) == 0 {
		return "", errors.New("input cannot be empty")
	}

	return getInput, nil
}

type config struct {
	numTime    int
	printUsage bool
}

// `args`는 명령줄 인자를 String Slice 형태로 전달 받도록 설계된 매개변수
func parseArgs(args []string) (config, error) {
	var numTimes int
	var err error

	cfg := config{}
	if len(args) != 1 {
		return cfg, errors.New("invalid number of arguments")
	}

	// 명령줄 인자가 하나인지 확인하고 지정된 인
	// 3수가 아닌 경우 오류를 반환하도록 처리
	if args[0] == "-h" || args[0] == "--help" {
		cfg.printUsage = true
		return cfg, nil
	}

	// `Atoi` Ref lib code : func strconv.Atoi(s string) (int, error)
	// `strconv.Atoi`는 문자열을 정수로 변환하며, 변환에 실패하면 오류를 반환
	numTimes, err = strconv.Atoi(args[0])
	if err != nil {
		return cfg, err
	}

	cfg.numTime = numTimes
	return cfg, nil
}

func greetUser(write io.Writer, name string, cfg config) {
	msg := fmt.Sprintf("Hello, %s\n", name)

	for i := 0; i < cfg.numTime; i++ {
		fmt.Fprint(write, msg)
	}
}

func runCMD(read io.Reader, write io.Writer, cfg config) error {
	if cfg.printUsage {
		printUsage(write)
		return nil
	}

	name, err := getName(read, write)
	if err != nil {
		return err
	}

	for i := 0; i < cfg.numTime; i++ {
		fmt.Fprintln(write, "Hello,", name)
	}

	greetUser(write, name, cfg)
	return nil
}

func main() {
	cfg, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Println("Error parsing arguments:", err)
		printUsage(io.Writer(os.Stdout))
		os.Exit(1)
	}
	

	err = validateArgs(cfg)
	if err != nil {
		fmt.Println("Error validating arguments:", err)
		os.Exit(1)
	}

	err = runCMD(os.Stdin, os.Stdout, cfg)
	if err != nil {
		fmt.Println("Error running command:", err)
		os.Exit(1)
	}
}
