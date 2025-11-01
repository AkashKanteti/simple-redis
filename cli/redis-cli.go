package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/AkashKanteti/simple-redis/serializer"
)

func main() {
	conn, err := net.Dial("tcp", ":6969")

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("%v", err)
	}

	text = strings.TrimSpace(text)
	// serialize using risp protocol
	serializedText := serializer.SerializeString(text)

	_, err = conn.Write([]byte(serializedText))
	if err != nil {
		fmt.Printf("%v", err)
	}

	var buf = make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Print(string(buf[:n]))
}
