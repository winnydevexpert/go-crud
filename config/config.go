package config

import "github.com/spf13/viper"

var AppConfig *ConfigStructure

type ConfigStructure struct {
	Mode               string `mapstructure:"MODE"`
	Port               string `mapstructure:"PORT"`
	DatabaseHost       string `mapstructure:"DB_HOST"`
	DatabasePort       string `mapstructure:"DB_PORT"`
	DatabaseName       string `mapstructure:"DB_NAME"`
	DatabaseUsername   string `mapstructure:"DB_USERNAME"`
	DatabasePassword   string `mapstructure:"DB_PASSWORD"`
	AccessSecretKey    string `mapstructure:"ACCESS_SECREAT_KEY"`
	AccessTokenExpire  string `mapstructure:"ACCESS_TOKEN_EXPIRE"`
	RefreshSecretKey   string `mapstructure:"REFRESH_SECREAT_KEY"`
	RefreshTokenExpire string `mapstructure:"REFRESH_TOKEN_EXPIRE"`
}

func LoadConfig(path string) (config *ConfigStructure, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	errViperReadConfig := viper.ReadInConfig()
	if errViperReadConfig != nil {
		return
	}

	errViperUnmarshal := viper.Unmarshal(&config)
	if errViperUnmarshal != nil {
		return
	}

	return
}
