package main

import "fmt"

/*
//// Scanf
Ref lib code : func Scanf(format string, a ...any) (n int, err error)
`Scanf()`는 표준 입력에서 서식 형태로 `Printf`와 유사하게 값을 읽어오는 것이 아닌 입력받는다.

//// Scanln
Ref lib code : func Scanln(a ...any) (n int, err error)
`Scanln()`는 표준 입력에서 한 줄 단위로 값을 읽어오는 함수로, 입력값이 공백으로 구분되어 여러 개일 경우에도 각 값을 변수에 저장할 수 있다.

FIFO [First In First Out]는 먼저 입력된 데이터가 먼저 읽히는 데이터 구조를 의미하지만, 
가장 먼저 입력한 데이터부터 읽기 때문에 데이터가 반대로 출력되는 경우가 발생할 수 있다.
특히 FIFO는 표준 입력 스트림에서 사용되고 있기에, 첫 글자가 먼저 입력되더라도
다른 입력값이 남아 있는 경우에는 예상과 다르게 처리될 수 있다. 즉 읽은 데이터는 다시 읽어올 수 없는 것이다.
*/

func main() {
	var command, options string

	//// Scan
	// `Scan`는 사용자가 입력한 값을 String 형태로 읽어오는 함수지만 `Print`와 같은 줄 바꿈 없이 출력되어 주로 명령어 입력을 받을 때 사용
	// 하지만 사용자가 입력한 값이 공백으로 구분되어 여러 개일 경우, 각각의 값을 변수에 저장하려면 Scan 함수에 해당 변수들을 순서대로 전달해야 한다.
	// Ref lib code : func Scan(a ...any) (n int, err error)
	input, err := fmt.Scan(&command, &options)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if command == "docker" {
		if options == "run" {
			fmt.Println("Executing docker run command")
		}

		if options == "stop" {
			fmt.Println("Executing docker stop command")
		}

		if options == "" {
			fmt.Println("Executing docker command without options")
		}
	}

	fmt.Println("User input:", input)
}
