package config

import (
	"errors"
	"github.com/spf13/viper"
)

type Config struct {
	SnmpCfg   SnmpConfig
	ServerCfg ServerConfig
}

type SnmpConfig struct {
	IpOlt     string
	Community string
}

type ServerConfig struct {
	Host string
	Port uint16
}

// LoadConfig file from given path using viper
func LoadConfig(filename string) (*Config, error) {
	v := viper.New()
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	var cfg Config

	// Unmarshal config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
