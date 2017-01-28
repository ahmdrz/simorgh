package simorgh

import (
	"log"
	"net"
)

type Simorgh struct {
	conn net.Conn
}

func New(connString string, proto string) (*Simorgh, error) {
	conn, err := net.Dial(proto, connString)
	if err != nil {
		log.Println(err)
		return nil,err
	}
	return &Simorgh{
		conn : conn,
	},
}

func (s *Simorgh) Close() error {
	return s.conn.Close()
}

func (s *Simorgh) Set(key, value string) {
	s.conn.Write([]byte("{set " + a + "=" + b + "}\n")
}
