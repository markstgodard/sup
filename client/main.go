package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	p := make([]byte, 1024)
	conn, err := net.Dial("udp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	defer conn.Close()

	msg := "hey"

	fmt.Printf("Send: %s\n", msg)

	fmt.Fprintf(conn, msg)

	_, err = bufio.NewReader(conn).Read(p)
	if err != nil {
		log.Fatalf("error %v\n", err)
	}
	fmt.Printf("Resp: %s\n", p)
}
