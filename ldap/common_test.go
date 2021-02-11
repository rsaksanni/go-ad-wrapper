package ldap

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var testConfig struct {
	Server      string
	Port        int
	BindUPN     string
	BindPass    string
	BaseDN      string
	PasswordUPN string
}

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	testConfig.Server = os.Getenv("ADTEST_SERVER")

	if port, err := strconv.Atoi(os.Getenv("ADTEST_PORT")); err == nil {
		testConfig.Port = port
	} else {
		testConfig.Port = 389
	}

	testConfig.BindUPN = os.Getenv("ADTEST_BIND_UPN")
	testConfig.BindPass = os.Getenv("ADTEST_BIND_PASS")

	testConfig.BaseDN = os.Getenv("ADTEST_BASEDN")
	testConfig.PasswordUPN = os.Getenv("ADTEST_PASSWORD_UPN")
}
