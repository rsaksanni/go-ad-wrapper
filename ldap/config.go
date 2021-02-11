package ldap

import (
	"errors"
	"strings"
)

// Config for AD connection
type Config struct {
	Server string
	Port   int
	BaseDN string
}

// OU func return current Organizational Unit name
func (c *Config) OU() (string, error) {
	var ou string
	for _, val := range strings.Split(strings.ToLower(c.BaseDN), ",") {
		if trimmed := strings.TrimSpace(val); strings.HasPrefix(trimmed, "ou=") {
			ou = trimmed[3:]
		}
	}
	if ou == "" {
		return "", errors.New("Config error, BaseDN is invalid or OU is not set")
	}

	return ou, nil
}
