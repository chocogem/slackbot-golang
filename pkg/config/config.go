package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	SlackAuthToken  string `mapstructure:"SLACK_AUTH_TOKEN" `
	SlackChannelId  string `mapstructure:"SLACK_CHANNEL_ID" `
	SlackMessageUrl string `mapstructure:"SLACK_MESSAGE_URL" `
	LogrusLogLevel  string `mapstructure:"LOGRUS_LOG_LEVEL" `
	AppPort         string `mapstructure:"APP_PORT" `
}

// List of environment variables to fetch for
var envs = []string{
	"SLACK_AUTH_TOKEN","SLACK_CHANNEL_ID","LOGRUS_LOG_LEVEL",
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile("local.env")
	viper.ReadInConfig()

	// Bind from environment variable
	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	return config, nil
}
