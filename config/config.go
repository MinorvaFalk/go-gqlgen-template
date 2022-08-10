package config

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var C config

type config struct {
	// Database Configuration
	DBHost     string `mapstructure:"DB_HOST"`
	DBUser     string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBPort     string `mapstructure:"DB_PORT"`

	// Endpoint configuration
	Port string `mapstructure:"PORT"`
}

func ReadConfig() {
	Config := &C

	viper.AddConfigPath(filepath.Join(rootDir(), "db"))
	viper.SetConfigName("db")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Fetch root directory path
func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
