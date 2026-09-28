package main

import (
	"fmt"
	"net"
)

func main() {
	// `Dial`
	conn, err := net.Dial("tcp", "localhost:50001")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: localhost:50001\r\n\r\n")

	// `Listen`
	ln, err := net.Listen("tcp", "localhost:50002")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go func(c net.Conn) {
			defer c.Close()
			buf := make([]byte, 1024)
			n, err := c.Read(buf)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(buf[:n]))
		}(conn)
	}
}

/*
1. func net.Dial(network string, address string) (net.Conn, error)
-> dial tcp [::1]:50001: connectex: No connection could be made because the target machine actively refused it.

2. func net.Listen(network string, address string) (net.Listener, error)
->
*/
