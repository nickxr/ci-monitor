package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port         string `mapstructure:"PORT"`
	DBHost       string `mapstructure:"DB_HOST"`
	DBPort       string `mapstructure:"DB_PORT"`
	DBUser       string `mapstructure:"DB_USER"`
	DBPassword   string `mapstructure:"DB_PASSWORD"`
	DBName       string `mapstructure:"DB_NAME"`
	GitHubSecret string `mapstructure:"GITHUB_SECRET"`
}

func LoadConfig() (Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("No .env file found, using env variables")
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}
