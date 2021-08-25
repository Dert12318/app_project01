package env

import (
	"app_project01/models"
	"fmt"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

var Config models.ServerConfig

func init() {
	err := loadConfig()
	if err != nil {
		fmt.Printf(err)
	}
}
func loadConfig() (err error) {
	err = godotenv.Load()
	if err != nil {

	}

	err = env.Parse(&Config)
	if err != nil {
		return err
	}

	err = env.Parse(&Config.PostgresConfig)
	if err != nil {
		return err
	}

	return err
}