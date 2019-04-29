package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := strings.ToLower(scanner.Text())
		bs := []byte(ln) // byte conversion
		r := rot13(bs)

		fmt.Fprintf(conn, "%s - %s\n", ln, r)
	}
}

func rot13(bs []byte) []byte {
	var r13 = make([]byte, len(bs))
	for i, v := range bs { // rotate ascii value by 13
		if v <= 109 {
			r13[i] = v + 13
		} else {
			r13[i] = v - 13
		}
	}
	return r13
}

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn)
	}
}
