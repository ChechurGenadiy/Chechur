package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP: []byte{0, 0, 0, 0}, Port: 10001})
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	message := "Hello World"
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Println(err)
	}

}
