package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

func handle(conn net.Conn) {
	defer conn.Close()

	io.WriteString(conn, "\nIN-MEMORY DATABASE\n\n"+
		"USAGE:\n"+
		"SET key value\n"+
		"GET key value\n"+
		"DEL key\n\n"+
		"EXAMPLE:\n"+
		"SET fav chocolate\n"+
		"GET fav\n\n\n")

	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln) // breaks the sting into individual words and stores them into slice of strings

		fs[0] = strings.ToLower(fs[0])
		switch fs[0] {
		case "get":
			k := fs[1]
			v := data[k]
			fmt.Fprintf(conn, "%s\n", v)
		case "set":
			if len(fs) != 3 {
				fmt.Fprintln(conn, "Wrong number of items in query, expected 3")
				continue
			}
			k := fs[1]
			v := fs[2]
			data[k] = v
			fmt.Fprintln(conn, "Succesfully inserted in DB")
		case "del":
			k := fs[1]
			delete(data, k)
			fmt.Fprintln(conn, "Succesfully removed from DB")
		default:
			fmt.Fprintln(conn, "INVALID COMMAND")
		}
	}
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
