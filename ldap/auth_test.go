package ldap

import (
	"testing"
)

// TestAuthenticate ...
func TestAuthenticate(t *testing.T) {

	if testConfig.Server == "" {
		t.Skip("ADTEST_SERVER not set")
		return
	}

	if testConfig.BaseDN == "" {
		t.Skip("ADTEST_BASEDN not set")
		return
	}

	config := &Config{Server: testConfig.Server, Port: testConfig.Port, BaseDN: testConfig.BaseDN}

	status, err := Authenticate(config, "go-ad-wrapper", "invalid password")
	if err == nil {
		t.Error("Invalid credentials: Expected err to be nil but got:", err)
	}
	if status {
		t.Error("Invalid credentials: Expected authenticate status to be false")
	}

	status, err = Authenticate(config, testConfig.BindUPN, testConfig.BindPass)
	if err != nil {
		t.Fatal("Valid UPN: Expected err to be nil but got:", err)
	}
	if !status {
		t.Error("Valid UPN: Expected authenticate status to be true")
	}

}
