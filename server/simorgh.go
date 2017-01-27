package main

import (
	"bufio"
	"flag"
	"log"
	"net"
	"strconv"
	"strings"

	"gitlab.com/ahmdrz/simon/tree"
)

var simon struct {
	tree *tree.Tree
}

func init() {
	simon.tree = tree.NewTree()
}

func main() {
	port := flag.String("port", "8080", "Server port")
	proto := flag.String("protocol", "tcp", "Server net listen protocol")
	flag.Parse()

	ln, err := net.Listen(*proto, ":"+*port)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	for {
		cmd, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			conn.Write([]byte("BAD REQUEST"))
			conn.Close()
			return
		}
		cmd = strings.TrimSpace(cmd)

		if strings.HasPrefix(cmd, "{") && strings.HasSuffix(cmd, "}") {
			cmd = cmd[1 : len(cmd)-1]
			if strings.HasPrefix(cmd, "set") {
				cmd = strings.Replace(cmd, "set", "", -1)
				cmd = strings.TrimSpace(cmd)
				parts := strings.Split(cmd, "=")
				if len(parts) == 2 {
					simon.tree.Set(parts[0], parts[1])
					conn.Write([]byte(parts[0] + " = " + parts[1] + "\n"))
				} else {
					conn.Write([]byte("INVALID\n"))
				}
			} else if strings.HasPrefix(cmd, "get") {
				cmd = strings.Replace(cmd, "get", "", -1)
				cmd = strings.TrimSpace(cmd)
				value := simon.tree.Get(cmd)
				conn.Write([]byte(cmd + " = " + value + "\n"))
			} else if strings.HasPrefix(cmd, "del") {
				cmd = strings.Replace(cmd, "del", "", -1)
				cmd = strings.TrimSpace(cmd)
				simon.tree.Del(cmd)
				conn.Write([]byte(cmd + " REMOVED\n"))
			} else if cmd == "clr" {
				n := simon.tree.Clr()
				conn.Write([]byte("MEMORY CLEARED (" + strconv.Itoa(n) + ")\n"))
			} else {
				conn.Write([]byte("INVALID\n"))
			}
		} else {
			conn.Write([]byte("INVALID\n"))
		}
	}
}
