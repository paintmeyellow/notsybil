package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	APIKey     string `mapstructure:"api_key" json:"api_key"`
	SecretKey  string `mapstructure:"secret_key" json:"secret_key"`
	Passphrase string `mapstructure:"passphrase" json:"passphrase"`
}

func Load(configPath, configName, configType string) (*Config, error) {
	v := viper.New()
	v.AddConfigPath(configPath)
	v.SetConfigName(configName)
	v.SetConfigType(configType)

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	var conf Config
	if err := v.Unmarshal(&conf); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}

	return &conf, nil
}
