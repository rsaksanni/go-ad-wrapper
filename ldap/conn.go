package ldap

import (
	"errors"
	"fmt"
	"strings"

	ldap "github.com/go-ldap/ldap/v3"
)

// Conn struct
type Conn struct {
	Conn   *ldap.Conn
	Config *Config
}

var connection *Conn

// Connect func
func (c *Config) Connect() (*Conn, error) {
	if connection == nil {
		con, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", c.Server, c.Port))
		if err != nil {
			return nil, fmt.Errorf("Connection error: %v", err)
		}
		connection = &Conn{Conn: con, Config: c}
	}

	return connection, nil

}

// Bind func
func (c *Conn) Bind(upn, password string) (bool, error) {
	if len(strings.TrimSpace(password)) <= 1 {
		return false, errors.New("Bind error: password length is not comply")
	}
	err := c.Conn.Bind(upn, password)
	if err != nil {
		return false, fmt.Errorf("Bind error (%s): %v", upn, err)
	}
	return true, nil
}
