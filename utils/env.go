package utils

import (
	"github.com/spf13/viper"
)

type Env struct {
	// Database Configuration
	DBHost     string `mapstructure:"DB_HOST"`
	DBUser     string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBPort     string `mapstructure:"DB_PORT"`

	// Endpoint configuration
	Port string `mapstructure:"PORT"`
}

func NewEnv() (*Env, error) {
	env := Env{}

	viper.SetConfigFile("../.env")
	if err := viper.ReadInConfig(); err != nil {

		return nil, err

		// if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		// 	return nil, ErrEnvFileNotFound
		// } else {
		// 	return nil, err
		// }
	}

	return &env, nil
}
