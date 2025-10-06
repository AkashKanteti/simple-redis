package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":6969")

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("%v", err)
	}

	//cmds := strings.Split(strings.TrimSpace(text), " ")

	_, err = conn.Write([]byte(text))
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
