package main

import (
	"bufio"
	"flag"
	"fmt"
	"gopass"
	"log"
	"os"

	"github.com/ahmdrz/simorgh/driver"
	"strings"
)

func main() {
	addr := flag.String("address", "localhost", "Server address")
	port := flag.String("port", "8080", "Server port")
	proto := flag.String("protocol", "tcp", "Server net listen protocol")
	flag.Parse()

	var username string
	fmt.Print("Username > ")
	fmt.Scanf("%s", &username)
	fmt.Printf("Password > ")

	password, err := gopass.GetPasswd()
	if err != nil {
		fmt.Println("Cannot read input password")
		os.Exit(-1)
	}

	si, err := simorgh.New(*addr+":"+*port, username, string(password), *proto)
	if err != nil {
		log.Println(err)
		os.Exit(-3)
	}
	defer si.Close()

	for {
		fmt.Print("< ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-1]
		if text == "\\q" {
			fmt.Println("bye")
			os.Exit(0)
		}
		if strings.HasPrefix(text, "set") {
			text = text[3:]
			parts := strings.Split(text, "=")
			msg, err := si.Set(parts[0], parts[1])
			if err != nil {
				fmt.Println("> ERROR !", msg)
			} else {
				fmt.Println("> OK", msg)
			}
		} else if strings.HasPrefix(text, "get") {
			text = text[3:]
			msg, err := si.Get(text)
			if err != nil {
				fmt.Println("> ERROR !", msg)
			} else {
				fmt.Println("> OK", msg)
			}
		} else if text == "clr" {
			msg, err := si.Clr()
			if err != nil {
				fmt.Println("> ERROR !", msg)
			} else {
				fmt.Println("> OK", msg)
			}
		} else if strings.HasPrefix(text, "del") {
			text = text[3:]
			msg, err := si.Del(text)
			if err != nil {
				fmt.Println("> ERROR !", msg)
			} else {
				fmt.Println("> OK", msg)
			}
		}
	}
	os.Exit(0)
}
