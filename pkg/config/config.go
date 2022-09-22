package config

import "github.com/spf13/viper"

type Config struct {
	Port     string `mapstructure:"PORT"`
	Username string `mapstructure:"USR"`
	Password string `mapstructure:"PASS"`
	Url      string `mapstructure:"URL"`
	DbName   string `mapstructure:"DBNAME"`
}

// LoadConfig Load db config, returns empty config on error
func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
