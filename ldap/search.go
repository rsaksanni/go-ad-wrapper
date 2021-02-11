package ldap

import (
	"fmt"

	ldap "github.com/go-ldap/ldap/v3"
)

// SearchOne ..
func (c *Conn) SearchOne(filter string, attrs []string) (*ldap.Entry, error) {
	search := ldap.NewSearchRequest(
		c.Config.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.DerefAlways,
		1,
		0,
		false,
		filter,
		attrs,
		nil,
	)

	result, err := c.Conn.Search(search)
	if err != nil {
		return nil, fmt.Errorf(`Search error "%s": %v`, filter, err)
	}

	if len(result.Entries) == 0 {
		return nil, fmt.Errorf(`Search error "%s": no entries returned`, filter)
	}

	return result.Entries[0], nil
}
