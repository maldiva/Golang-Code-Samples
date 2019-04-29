package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"strings"
	"text/template"
)

type Request struct {
	URL  string
	Type string
}

func handle(conn net.Conn) {

	defer conn.Close()

	r := request(conn)

	respond(conn, r)
}

func request(conn net.Conn) Request {

	i := 0
	scanner := bufio.NewScanner(conn)
	var r Request

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if i == 0 {
			req := strings.Fields(ln)
			fmt.Println("URL of the request is: ", req[1])
			m := req[0]
			fmt.Println("***METHOD", m)

			r.Type = req[0]
			r.URL = req[1]

		}

		if ln == "" { // there is a blank line between headers and body
			break
		}
		i++
	}
	return r
}

func abort(conn net.Conn) {
	fmt.Fprint(conn, "HTTP/1.1 404 Not Found\n")
	fmt.Fprintf(conn, "Content-Length: %d\n", 26)
	fmt.Fprint(conn, "Content-Type: text/html\n")
	fmt.Fprint(conn, "\r\n")
}

func handleRequest(conn net.Conn, r Request) {

	body := "<html><header><title>This is title</title></header><body>This is {{.Type}} request on URL: {{.URL}}</body></html>"

	t := template.Must(template.New("body").Parse(body))
	var tpl bytes.Buffer
	err := t.Execute(&tpl, r)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprint(conn, "HTTP/1.1 200 OK\n")
	fmt.Fprintf(conn, "Content-Length: %d\n", len(tpl.String()))
	fmt.Fprint(conn, "Content-Type: text/html\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, tpl.String())

}

func respond(conn net.Conn, r Request) {

	if r.Type == "GET" || r.Type == "POST" {
		handleRequest(conn, r)
	} else {
		abort(conn)
	}

}

func main() {

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go handle(conn)
	}
}
