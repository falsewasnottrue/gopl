package main

import (
	"fmt"
	"net"
	"log"
	"bufio"
	"time"
)

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}

func echo(c net.Conn, input string, delay time.Duration) {
	fmt.Fprintln(c, input)
	time.Sleep(delay)
	fmt.Fprintln(c, input)
	time.Sleep(delay)
	fmt.Fprintln(c, input)
}