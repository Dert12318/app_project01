package env

import (
	"/models"
)

var Config models.ServerConfig



func loadConfig() (err error) {
	err = godotenv.Load()
	if err != nil {
		logrus.Fatal(err, " config/env: load gotdotenv")
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