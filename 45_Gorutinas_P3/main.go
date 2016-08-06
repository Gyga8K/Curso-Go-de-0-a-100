package main

import (
	"io"
	"log"
	"net"
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
			log.Print(err) //Por ejemplo, conexión abortada
			continue
		}
		go manejarCon(conn)
	}

}

func manejarCon(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
		if err != nil {
			return // ejemplo, si el cliente se desconecta
		}
		time.Sleep(1 * time.Second)
	}
}
