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

type ConfigOption struct {
	AppEnv string
}

type config struct {
	Database struct {
		DBHost     string `mapstructure:"host"`
		DBName     string `mapstructure:"dbname"`
		DBUser     string `mapstructure:"user"`
		DBPassword string `mapstructure:"password"`
	}

	Server struct {
		Port string `mapstructure:"port"`
	}
}

func ReadConfig(option ConfigOption) {
	Config := &C

	viper.AddConfigPath(filepath.Join(rootDir(), "config"))
	viper.SetConfigName("config")

	viper.SetConfigType("yml")
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
