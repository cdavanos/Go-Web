package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	// listen to the tcp connection on 8080
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer li.Close()

	for {
		// accept the connection
		conn, err := li.Accept()
		if err != nil {
			panic(err)
		}
		// handle the conn in a goroutine
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		// convert to lower case, then store in a byteslice and
		// send it off to rot13
		ln := strings.ToLower(scanner.Text())
		bs := []byte(ln)
		r := rot13(bs)

		// print out the return along with the original line
		fmt.Fprintf(conn, "%s - %s\n\n", ln, r)
	}
}

func rot13(bs []byte) []byte {
	// make a slicde of bytes of length equal to bs passed in
	// slice byte of line of text passed in
	var r13 = make([]byte, len(bs))
	// range over the byteslice
	for i, v := range bs {
		// ascii 97-122
		if v <= 109 {
			r13[i] = v + 13
		} else {
			r13[i] = v - 13
		}
	}
	return r13
}
