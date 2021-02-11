package ldap

import "fmt"

// IsAccountExpired func
func (c *Conn) IsAccountExpired(upn string) (bool, error) {
	filter := fmt.Sprintf("(userPrincipalName=%s)", upn)
	attrs := []string{"accountExpires"}
	
	entry, err := c.SearchOne(filter, attrs)

	if err != nil {
		return true, fmt.Errorf("Account expiry: return false with exception: %v", err)
	}

	// 64-bit signed integer of 0xFFFFFFFFF -> 9223372036854775807
	if entry.Attributes[0].Values[0] == "9223372036854775807" {
		return false, nil
	}

	return true, nil
}
