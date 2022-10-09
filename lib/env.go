package lib

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	DB_URL string `mapstructure:"DB_URL"`
}

func NewEnv() Env {
	env := Env{}

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't read configuration file")
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("The Env file cannot be loaded -> ", err)
	}

	return env
}
