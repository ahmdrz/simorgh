package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	addr := flag.String("address", "localhost", "Server address")
	port := flag.String("port", "8080", "Server port")
	proto := flag.String("protocol", "tcp", "Server net listen protocol")
	flag.Parse()

	conn, err := net.Dial(*proto, *addr+":"+*port)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		fmt.Print("< ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-1]
		if text == "\\q" {
			return
		}
		fmt.Fprintf(conn, "{"+text+"}\n")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		message = strings.TrimSpace(message)

		if strings.HasPrefix(message, "{") && strings.HasSuffix(message, "}") {
			message = message[1 : len(message)-1]
			fmt.Println("> " + message)
		} else {
			fmt.Println("> ERROR ON READING DATA")
			return
		}
	}
}
