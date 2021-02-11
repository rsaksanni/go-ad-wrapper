package ldap

import (
	"fmt"
	"testing"
)

func TestConnSearchOne(t *testing.T) {
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
	defer conn.Conn.Close()

	status, err := conn.Bind(testConfig.BindUPN, testConfig.BindPass)
	if err != nil {
		t.Fatal("Error binding to server:", err)
	}

	if !status {
		t.Fatal("Error binding to server: invalid credentials")
	}

	if _, err = conn.SearchOne(fmt.Sprintf("(userPrincipalName=%s)", testConfig.BindUPN), []string{}); err != nil {
		t.Error("SearchOne: valid search: Expected err to be nil but got:", err)
	}

	if entry, _ := conn.SearchOne(fmt.Sprintf("(userPrincipalName=%s)", testConfig.BindUPN), []string{"accountExpires"}); entry != nil {
		// fmt.Println(entry.Attributes[0].Values[0])
		if entry.Attributes[0].Values[0] != "9223372036854775807" {
			t.Error("SearchOne: valid search: Expected max 64-bit signed integer as not expired but got expired")
		}

	}
}
