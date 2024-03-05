package util

import "github.com/spf13/viper"

type Config struct {
	SERVER_ADDRESS string `SERVER_ADDRESS`
	KAFKA_ADDRESS  string `KAFKA_ADDRESS`
}

func LoadConfig() (config *Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env.local")
	viper.SetConfigType("dotenv")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
