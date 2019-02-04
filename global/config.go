package global

import (
	"fmt"
)

// ConfigData - config variables of app
type ConfigData struct {
	AppPort  int
	Host     string
	DBPort   int
	Database string
	User     string
	Password string
}

func (c ConfigData) String() string {
	pwd := c.Password
	if pwd != "" {
		pwd = "************"
	}
	return fmt.Sprintf("{ Host: %v; Port: %v; Database: %v; User: %v; Password: %v}", c.Host, c.DBPort, c.Database, c.User, pwd)
}
