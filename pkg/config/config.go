package config

type Config struct {
}

func DefaultConfig() (*Config, error) {
	return &Config{}, nil
}
