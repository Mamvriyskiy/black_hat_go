package main 

import (
	"net"
	"bufio"
	"log"
)

func echo(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	defer writer.Flush()
	
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln("Error read string")
		}

		log.Println("Server:", str)
		_, err = writer.WriteString(str)
		if err != nil {
			log.Fatalln("Error write string")
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
