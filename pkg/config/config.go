package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port    int
	AuthUrl string
}

func init() {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func DefaultConfig() (*Config, error) {
	return &Config{}, nil
}

func LoadConfig() (*Config, error) {
	var c *Config
	err := viper.Unmarshal(&c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Config) string() string {
	return fmt.Sprintf("%d, %s", c.Port, c.AuthUrl)
}
