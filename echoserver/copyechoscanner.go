package main 

import (
	"net"
	"io"
	"log"
)

func echo(conn net.Conn) {
	defer conn.Close()
	
	for {
		if _, err := io.Copy(conn, conn); err != nil {
			log.Fatalln("Error read string")
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}

	log.Println("Listening port: 20080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to connect")
		}

		log.Println("Received connection")

		go echo(conn)
	}
}
