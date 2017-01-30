package simorgh

import (
	"bufio"
	"fmt"
	"go-srp/src/srp"
	"net"
	"strings"
)

const SAFE_PRIME_BITS = 2048

func validate(username, password []byte, conn net.Conn) ([]byte, error) {
	c, err := srp.NewClient(username, password, SAFE_PRIME_BITS)
	if err != nil {
		return nil, fmt.Errorf("SRP AUTHENTICATION FAILED")
	}

	conn.Write([]byte("{--cred--" + c.Credentials() + "--cred--}\n"))
	message, _ := bufio.NewReader(conn).ReadString('\n')
	message = strings.TrimSpace(message)

	if strings.HasPrefix(message, "{") && strings.HasSuffix(message, "}") {
		message = message[1 : len(message)-1]
		if strings.HasPrefix(message, "--cred--") && strings.HasSuffix(message, "--cred--") {
			server_creds := strings.Replace(message, "--cred--", "", -1)
			auth, err := c.Generate(server_creds)
			if err != nil {
				return nil, err
			}
			conn.Write([]byte("{--auth--" + auth + "--auth--}\n"))
		} else if message == "--validate--0--validate--" {
			return nil, fmt.Errorf("INVALID USERNAME OR PASSWORD")
		}
	}

	message, _ = bufio.NewReader(conn).ReadString('\n')
	message = strings.TrimSpace(message)
	if strings.HasPrefix(message, "{") && strings.HasSuffix(message, "}") {
		message = message[1 : len(message)-1]
		if strings.HasPrefix(message, "--proof--") && strings.HasSuffix(message, "--proof--") {
			proof := strings.Replace(message, "--proof--", "", -1)
			err := c.ServerOk(proof)
			if err != nil {
				return nil, err
			}

			key := c.RawKey()
			return key, nil
		} else if message == "--validate--0--validate--" {
			return nil, fmt.Errorf("INVALID USERNAME OR PASSWORD")
		}
	}

	return nil, fmt.Errorf("Undefined error")
}
