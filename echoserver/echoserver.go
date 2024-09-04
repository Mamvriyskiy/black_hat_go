package main

import (
	"net"
	"log"
	"io"
)

func echo(conn net.Conn) {
	defer conn.Close()
	b := make([]byte, 255)

	for {
		n, err := conn.Read(b)
		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}
		if err != nil {
			log.Println("Unexpected error")
			break
		}
		log.Printf("Received %d bytes: %s\n", n, string(b))
		_, err = conn.Write(b)
		if err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}

func main () {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}

	log.Println("Listening on port 0.0.0.0:20080")

	for {
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}

		go echo(conn)
	}

}
