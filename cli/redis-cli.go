package main

import (
	"bufio"
	"fmt"
	"github.com/AkashKanteti/simple-redis/serializer"
	"net/http"
	"os"
	"strings"
)

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("%v", err)
		}

		if text == "exit\n" {
			break
		}

		cmds := strings.Split(strings.TrimSpace(text), " ")

		resp, err := http.Post("localhost:6379", "application/json", strings.NewReader(serializer.SerializeArray(cmds)))
		if err != nil {
			fmt.Printf("%v", err)
		}

		fmt.Println(resp.Body)
	}
}
