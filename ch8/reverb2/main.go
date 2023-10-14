// Clock1 is a TCP server that periodically writes the time.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle one connection at a time
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintf(c, "server > %s\n", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintf(c, "server > %s\n", shout)
	time.Sleep(delay)
	fmt.Fprintf(c, "server > %s\n", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}
