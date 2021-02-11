package ldap

import (
	"testing"
)

// TestExpiry func
func TestExpiry(t *testing.T) {
	if testConfig.Server == "" {
		t.Skip("ADTEST_SERVER not set")
		return
	}

	if testConfig.BindUPN == "" || testConfig.BindPass == "" {
		t.Skip("ADTEST_BIND_UPN or ADTEST_BIND_PASS not set")
		return
	}

	if testConfig.BaseDN == "" {
		t.Skip("ADTEST_BASEDN not set")
		return
	}

	config := &Config{Server: testConfig.Server, Port: testConfig.Port, BaseDN: testConfig.BaseDN}
	conn, err := config.Connect()
	if err != nil {
		t.Fatal("Error connecting to server:", err)
	}

	if expired, _ := conn.IsAccountExpired(testConfig.BindUPN); expired != false {
		t.Error("IsAccountExpired: Expected false while account is never expired")
	}
}
