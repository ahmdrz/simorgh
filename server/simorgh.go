package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"go-srp/src/srp"
	"tree"
)

const SAFE_PRIME_BITS = 2048
const LIMIT_TIME = 600

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

	go removeUnusedRequests()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleRequest(conn)
	}
}

func removeUnusedRequests() {
	simorgh.auth.mutex.Lock()
	for key, value := range simorgh.auth.accepts {
		if value+LIMIT_TIME < time.Now().Unix() {
			delete(simorgh.auth.accepts, key)
		}
	}
	simorgh.auth.mutex.Unlock()
	time.AfterFunc(5*time.Second, removeUnusedRequests)
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
				simorgh.auth.mutex.Lock()
				simorgh.auth.keys[conn.RemoteAddr().String()] = s
				simorgh.auth.mutex.Unlock()
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
				simorgh.auth.mutex.Lock()
				delete(simorgh.auth.keys, conn.RemoteAddr().String())
				var session []byte = s.RawKey()
				var _session []byte = make([]byte, base64.StdEncoding.EncodedLen(len(session)))
				base64.StdEncoding.Encode(_session, session)
				simorgh.auth.accepts[string(_session)] = time.Now().Unix()
				simorgh.auth.mutex.Unlock()
				conn.Write([]byte("{--proof--" + proof + "--proof--}\n"))
			}

			if _, valid := simorgh.auth.accepts[key]; valid {
				simorgh.auth.mutex.Lock()
				simorgh.auth.accepts[key] = time.Now().Unix()
				simorgh.auth.mutex.Unlock()
			} else {
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
				value, mode := simorgh.tree.Get(cmd)
				conn.Write([]byte("{" + value + "-mode:" + mode + "}\n"))
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
