package config

import (
	"fmt"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	MongoConfig struct {
		ConnectionURI string `env:"URI" envDefault:"mongodb://localhost:27017"`
		Database      string `env:"DB" envDefault:"inte"`
		Timeout       uint64 `env:"TIMEOUT" envDefault:"30"`
	} `envPrefix:"INTE_MONGO_"`

	ApiConfig struct {
		SessionSecret      string `env:"SESSION_SECRET"`
		AdminSessionSecret string `env:"ADMIN_SESSION_SECRET"`
		Port               string `env:"PORT" envDefault:":8080"`
		LocalToken         string `env:"LOCAL_TOKEN"`
	} `envPrefix:"INTE_API_"`

	OauthConfig struct {
		GoogleClientID     string `env:"GOOGLE_CLIENT_ID"`
		GoogleClientSecret string `env:"GOOGLE_CLIENT_SECRET"`
	} `envPrefix:"INTE_OAUTH_"`

	BaseURI struct {
		Front string `env:"FRONT"`
		Back string `env:"BACK"`
	} `envPrefix:"INTE_BASE_URI_"`

	LogLevel string `env:"INTE_LOG_LEVEL" envDefault:"info"`
}

var config Config

func GetConfig() Config {
	return config
}

func init() {
	godotenv.Load()
	if err := env.Parse(&config); err != nil {
		logrus.Fatal(err)
	}

	logrus.SetLevel(logrus.InfoLevel)
	if config.LogLevel == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
	}

	logrus.Info("Loaded config: ", fmt.Sprintf("%+v", config))
}