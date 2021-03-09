package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/willroberts/minecraft-client"
)

var (
	hostport string
	password string

	quitCommands = map[string]struct{}{
		"exit": {},
		"quit": {},
	}
)

func init() {
	flag.StringVar(&hostport, "hostport", "127.0.0.1:25575", "hostname and port to connect to")
	flag.StringVar(&password, "password", "minecraft", "rcon password")
	flag.Parse()
}

func main() {
	client, err := minecraft.NewClient(hostport)
	if err != nil {
		log.Fatal("failed to connect:", err)
	}
	defer client.Close()

	if err := client.Authenticate(password); err != nil {
		log.Fatal("failed to authenticate:", err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Starting RCON shell. Use 'exit', 'quit', or Ctrl-C to exit.")
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("failed to read input:", err)
		}
		command := strings.TrimRight(input, "\r\n")

		if _, ok := quitCommands[command]; ok {
			break
		}

		resp, err := client.SendCommand(command)
		if err != nil {
			log.Fatal("failed to send command:", err)
		}
		fmt.Println(resp)
	}
}
