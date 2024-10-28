package config

import (
	"log"
	"os"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	PostgreSQL `yaml:"postgresql"`
}

type PostgreSQL struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     uint64 `yaml:"port" env-default:"5432"`
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
	DBName   string `yaml:"db_name" env-required:"true"`
	SSLMode  string `yaml:"ssl_mode" env-default:"disable"`
}

var (
	config Config
	once   sync.Once
)

func MustLoad() *Config {
	once.Do(func() {
		configPath := os.Getenv("CONFIG_PATH")
		if configPath == "" {
			log.Fatal("CONFIG_PATH is not set")
		}

		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			log.Fatalf("config file does not exist: %s", err)
		}

		if err := cleanenv.ReadConfig(configPath, &config); err != nil {
			log.Fatalf("cannot read config: %s", err)
		}
	})

	return &config
}
