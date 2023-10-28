package common

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() (config Config, err error) {

	viper.AddConfigPath("./")
	viper.AddConfigPath("./app")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		log.Println("err reading .env viper", err)
		return
	}

	if err = viper.Unmarshal(&config); err != nil {
		log.Println("err unmarshaling config viper", err)
		return
	}

	return
}
