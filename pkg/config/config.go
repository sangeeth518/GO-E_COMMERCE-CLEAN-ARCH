package config

import "github.com/spf13/viper"

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	// JWTToken   string `mapstructure:"JWTTOKEN"`
	JWTToken     string `mapstructure:"JWTTOKEN"`      // admin access token secret
	UserJWTToken string `mapstructure:"USER_JWTTOKEN"` // user access token secret
	RefreshToken string `mapstructure:"REFRESH_TOKEN"` // refresh token secret
	AWSRegion    string `mapstructure:"AWS_REGION"`
	AWSKey       string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AWSSecret    string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	BucketName   string `mapstructure:"AWS_BUCKET_NAME"`
}

var envs = []string{
	"DB_HOST", "DB_NAME", "DB_PORT", "DB_USER", "DB_PASSWORD", "JWTTOKEN", "USER_JWTTOKEN", "REFRESH_TOKEN",
	"AWS_REGION", "AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY", "AWS_BUCKET_NAME",
}

func LoadConfig() (Config, error) {
	var config Config
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
