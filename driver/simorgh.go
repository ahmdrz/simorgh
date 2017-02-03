package simorgh

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"net"
	"strings"
	"time"
)

type Simorgh struct {
	conn net.Conn
	key  string
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
	var _session []byte = make([]byte, base64.StdEncoding.EncodedLen(len(session)))
	base64.StdEncoding.Encode(_session, session)
	return &Simorgh{
		conn: conn,
		key:  string(_session),
	}, nil
}

func (s *Simorgh) Close() error {
	s.send("exit")
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
	type Invalid struct {
		Text  string
		Error error
	}
	invalid := make(chan Invalid, 1)
	go func() {
		s.conn.Write([]byte("{" + string(s.key) + "-&-" + cmd + "}\n"))
		reader := bufio.NewReader(s.conn)
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-1]
		if strings.HasPrefix(text, "{") && strings.HasSuffix(text, "}") {
			text = text[1 : len(text)-1]
			if text == "INVALID" || text == "UNDEFINED" {
				invalid <- Invalid{Text: text, Error: fmt.Errorf("(INVALID or UNDEFINED) SEND REQUEST")}
			}
		}
		invalid <- Invalid{Text: text, Error: nil}
	}()

	select {
	case result := <-invalid:
		return result.Text, result.Error
	case <-time.After(10 * time.Second):
		return "TIMEOUT", fmt.Errorf("REQUEST_TIME_OUT")
	}
}
