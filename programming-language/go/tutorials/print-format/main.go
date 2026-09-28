package main

import "fmt"

func main() {
	example_print("`Print` does not add a new line after output")

	example_println("`Println` is a new line after output")

	example_printf("This is an formatted string")
}

/*
줄 바꿈 없이 출력 [Print without newline]
`Print`는 줄 바꿈 없이 값들이 출력되어 주로 값은 이어서 출력될 때나 혹은 변경되어 출력하고자 할 때 사용 = Top와 같은 리소스 상황 확인할 때 사용
*/
func example_print(str string) {
	fmt.Print(str)
}

/*
줄 바꿈 있는 출력 [Print with newline]
`Println`은 출력 후 자동으로 줄 바꿈이 되어 다음 출력이 새로운 줄에서 시작될 때 사용 = 주로 Installation 과정이나 로그 출력 시 활용
*/
func example_println(str string) {
	fmt.Println(str)
}

/*
서식 지정자 [Format specifier]:
`Printf`는 서식 지정자를 포함한 형식화된 출력 방식으로, 다양한 데이터 타입이 지정되어 출력을 하고자 하거나 혹은 출력 형식을 세밀하게 제어하고자 할 때 사용

%v : 데이터 타입에 맞춘 기본 형태로 출력 [Value of the value]
%T : 데이터 타입 출력 [Type of the value]
%t : Boolean 참 혹은 거짓 형태로 출력 [True or False]
%d : 10진수 정수 형태로 출력 [Decimal integer value]
%b : 2진수 정수 형태로 출력 [Binary integer value]
%x \ %X : 16진수 정수 형태로 출력 [Hexadecimal integer value] Note: x/X는 10+ 이상 값은 a-f 소/대 문자로 표시
%c : 유니코드 단일 문자 형태로 출력 [Character value] Note: 정수 타입만 가능
%o \ %O : 8진수 정수 형태로 출력 [Octal integer value] Note: o/O는 10+ 이상 값은 0-7로 표시하지만 O는 앞에 8진수임을 명시하여 값을 출력
%e \ %E : 지수 형태로 출력 [Scientific notation] Note: e/E는 소수점 이하 값과 지수부를 포함하여 출력, E는 지수부를 대문자로 표시
%f \ %F : 소수 형태로 출력 [Floating-point number] Note: f/F는 소수점 이하 값을 포함하여 출력, F는 소수점 이하 값을 대문자로 표시
%g \ %G : 가장 간결한 형태로 출력 [Compact representation] Note: g/G는 소수점 이하 값과 지수부를 포함하여 출력, G는 지수부를 대문자로 표시
%s : 문자열 형태로 출력 [String value]
%q : 따옴표로 감싼 문자열 형태로 출력 [Quoted string value]
%p : 포인터 주소 형태로 출력 [Pointer address] | Example: fmt.Printf("%p", &variable)
*/
func example_printf(str string) {
	// `Printf`는 문자열을 위한 `%s`, 정수를 위한 `%d`와 같은 서식 지정자를 포함할 수 있는 형식화된 출력 방식
	fmt.Printf("%s\n", str)
}
