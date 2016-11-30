package main

import (
	"log"
	"net"
	"os"
)

//
// Simple UDP server that always responds with 'sup'
//
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("The PORT environment variable is empty")
	}

	serverAddr, err := net.ResolveUDPAddr("udp", ":"+port)
	if err != nil {
		log.Fatalf("error resolving UDP addr")
	}

	conn, err := net.ListenUDP("udp", serverAddr)
	if err != nil {
		log.Fatalf("error listening on UDP addr")
	}
	defer conn.Close()

	log.Println("sup is running...")
	b := make([]byte, 1024)
	for {
		_, remoteAddr, err := conn.ReadFromUDP(b)
		if err != nil {
			log.Println("Error: ", err)
		}
		go sup(conn, remoteAddr)
	}
}

func sup(conn *net.UDPConn, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("sup"), addr)
	if err != nil {
		log.Println("Error sending response: ", err)
	}
}
