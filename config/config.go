package config

import (
	"fmt"
	"os"
)

func GetDsn() string {
	user := os.Getenv("DB_GOMIC_USER")
	pass := os.Getenv("DB_GOMIC_PASS")
	name := os.Getenv("DB_GOMIC_NAME")
	host := os.Getenv("DB_GOMIC_HOST")
	return fmt.Sprintf("%s:%s@%s/%s", user, pass, host, name)
}
