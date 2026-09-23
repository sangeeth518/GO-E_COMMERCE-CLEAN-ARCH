package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	RabbitMQURL  string `mapstructure:"RABBITMQ_URL"`
	RedisHost    string `mapstructure:"REDIS_HOST"`
	RedisPort    string `mapstructure:"REDIS_PORT"`
	SMTPHost     string `mapstructure:"SMTP_HOST"`
	SMTPPort     int    `mapstructure:"SMTP_PORT"`
	SMTPEmail    string `mapstructure:"SMTP_EMAIL"`
	SMTPPassword string `mapstructure:"SMTP_PASSWORD"`
	SMTPName     string `mapstructure:"SMTP_NAME"`
}

var envs = []string{
	"RABBITMQ_URL",
	"REDIS_HOST",
	"REDIS_PORT",
	"SMTP_HOST",
	"SMTP_PORT",
	"SMTP_EMAIL",
	"SMTP_PASSWORD",
	"SMTP_NAME",
}

func LoadConfig() (Config, error) {
	var config Config
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	for _, val := range envs {
		if err := viper.BindEnv(val); err != nil {
			return config, err
		}

	}
	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}
	return config, nil
}
