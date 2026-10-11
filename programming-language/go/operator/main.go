package main

import "fmt"

/*

%08b: 8자리 2진수로 출력, 빈 자리는 0으로 채움

*/

func main() {
	var a int8 = 7
	// Result: 00000111 [7]
	fmt.Printf("%d = %08b,\t %16b\n", a, a, uint8(a))

	var b int8 = a
	// `>>`는 오른쪽으로 비트 이동을 n만큼 수행하지만 그만큼 빈자리의 비트는 부호 비트로 채워진다.
	b = a >> 1
	// Result: 00000011 [3]
	fmt.Printf("%d = %08b,\t %16b\n", b, b, uint8(b))

	// `|`는 비트 단위 OR 연산을 수행
	if a|b != 0 {
		fmt.Println("a | b is not 0")
	}
}
