package ldap

// Authenticate func
func Authenticate(config *Config, upn, password string) (bool, error) {
	conn, err := config.Connect()
	if err != nil {
		return false, err
	}

	return conn.Bind(upn, password)
}
