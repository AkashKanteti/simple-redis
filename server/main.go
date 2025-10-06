package main

import (
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":6969")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}

		handleFunction(conn)
	}
}

func handleFunction(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println(err)
		return
	}

	resp := commandMux(string(buf[:n]))

	conn.Write([]byte(resp))
}

func commandMux(cmd string) string {
	cmd = strings.TrimSpace(cmd)

	cmdArr := strings.Split(cmd, " ")

	switch cmdArr[0] {
	case "PING":
		return "PONG"
	case "ECHO":
		return strings.Join(cmdArr[1:], " ")
	}

	return ""
}
