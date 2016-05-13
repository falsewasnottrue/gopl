package main

import (
	"fmt"
	"net"
	"log"
	"bufio"
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
		// echo(c, input.Text(), 1*time.Second)
		fmt.Printf("received %v\n", input.Text())
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}

