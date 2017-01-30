package main

import (
	"bufio"
	"flag"
	"log"
	"net"
	"strconv"
	"strings"

	"go-srp/src/srp"
	"tree"
)

const SAFE_PRIME_BITS = 2048

var DEFAULT_USERNAME = []byte("simorgh")
var DEFAULT_PASSWORD = []byte("simorgh")

var simorgh struct {
	tree *tree.Tree
	auth *Authentication
}

func init() {
	simorgh.tree = tree.NewTree()
	Ih, salt, v, err := srp.Verifier(DEFAULT_USERNAME, DEFAULT_PASSWORD, SAFE_PRIME_BITS)
	if err != nil {
		panic(err)
	}
	simorgh.auth = InitAuthentication()
	simorgh.auth.salt = salt
	simorgh.auth.verifier = v
	simorgh.auth.identityHash = Ih
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
			var key string = ""
			if strings.Contains(cmd, "-&-") {
				key = strings.Split(cmd, "-&-")[0]
				cmd = cmd[strings.Index(cmd, "-&-")+3:]
			}

			if strings.HasPrefix(cmd, "--cred--") && strings.HasSuffix(cmd, "--cred--") {
				creds := strings.Replace(cmd, "--cred--", "", -1)
				I, A, err := srp.ServerBegin(creds)
				s, err := srp.NewServer(I, simorgh.auth.salt, simorgh.auth.verifier, A, SAFE_PRIME_BITS)
				if err != nil {
					conn.Write([]byte("{--validate--0--validate--}\n"))
					continue
				}
				simorgh.auth.keys[conn.RemoteAddr().String()] = s
				s_creds := s.Credentials()
				conn.Write([]byte("{--cred--" + s_creds + "--cred--}\n"))
			} else if strings.HasPrefix(cmd, "--auth--") && strings.HasSuffix(cmd, "--auth--") {
				auth := strings.Replace(cmd, "--auth--", "", -1)
				s, ok := simorgh.auth.keys[conn.RemoteAddr().String()]
				if !ok {
					conn.Write([]byte("{--validate--0--validate--}\n"))
					continue
				}
				proof, err := s.ClientOk(auth)
				if err != nil {
					conn.Write([]byte("{--validate--0--validate--}\n"))
					continue
				}
				simorgh.auth.accepts[string(s.RawKey())] = true
				conn.Write([]byte("{--proof--" + proof + "--proof--}\n"))
			}

			if _, valid := simorgh.auth.accepts[key]; !valid {
				continue
			}

			if strings.HasPrefix(cmd, "set") {
				cmd = strings.Replace(cmd, "set", "", -1)
				cmd = strings.TrimSpace(cmd)
				parts := strings.Split(cmd, "=")
				if len(parts) == 2 {
					simorgh.tree.Set(parts[0], parts[1])
					conn.Write([]byte("{OK}\n"))
				} else {
					conn.Write([]byte("{INVALID}\n"))
				}
			} else if strings.HasPrefix(cmd, "get") {
				cmd = strings.Replace(cmd, "get", "", -1)
				cmd = strings.TrimSpace(cmd)
				value := simorgh.tree.Get(cmd)
				conn.Write([]byte("{" + value + "}\n"))
			} else if strings.HasPrefix(cmd, "del") {
				cmd = strings.Replace(cmd, "del", "", -1)
				cmd = strings.TrimSpace(cmd)
				simorgh.tree.Del(cmd)
				conn.Write([]byte("{OK}\n"))
			} else if cmd == "clr" {
				n := simorgh.tree.Clr()
				conn.Write([]byte("{MEMORY CLEARED (" + strconv.Itoa(n) + ")}\n"))
			} else {
				conn.Write([]byte("{INVALID}\n"))
			}
		} else {
			conn.Write([]byte("{INVALID}\n"))
		}
	}
}
