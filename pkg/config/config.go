package config

import "github.com/spf13/viper"

type config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
}

var envs = []string{
	"DB_HOST", "DB_NAME", "DB_PORT", "DB_USER", "DB_PASSWORD",
}

func LoadConfig() (config, error) {
	var config config
	viper.AddConfigPath("./")
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
