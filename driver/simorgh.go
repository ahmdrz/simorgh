package simorgh

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Simorgh struct {
	conn net.Conn
	key  []byte
}

func New(address, username, password, proto string) (*Simorgh, error) {
	conn, err := net.Dial(proto, address)
	if err != nil {
		return nil, err
	}
	session, err := validate([]byte(username), []byte(password), conn)
	if err != nil {
		return nil, err
	}
	return &Simorgh{
		conn: conn,
		key:  session,
	}, nil
}

func (s *Simorgh) Close() error {
	return s.conn.Close()
}

func (s *Simorgh) Set(key, value string) (string, error) {
	return s.send("set " + key + "=" + value)
}

func (s *Simorgh) Get(key string) (string, error) {
	return s.send("get " + key)
}

func (s *Simorgh) Del(key string) (string, error) {
	return s.send("del " + key)
}

func (s *Simorgh) Clr() (string, error) {
	return s.send("clr")
}

func (s *Simorgh) send(cmd string) (string, error) {
	s.conn.Write([]byte("{" + string(s.key) + "-&-" + cmd + "}\n"))
	reader := bufio.NewReader(s.conn)
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]
	if strings.HasPrefix(text, "{") && strings.HasSuffix(text, "}") {
		text = text[1 : len(text)-1]
		if text == "INVALID" || text == "UNDEFINED" {
			return text, fmt.Errorf("(INVALID or UNDEFINED) SEND REQUEST")
		}
	}
	return text, nil
}