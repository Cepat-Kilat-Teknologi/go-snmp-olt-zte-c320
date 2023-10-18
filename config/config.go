package config

import (
	"errors"
	"github.com/spf13/viper"
)

type Config struct {
	SnmpCfg    SnmpConfig
	ServerCfg  ServerConfig
	OltCfg     OltConfig
	Board1Pon1 Board1Pon1
	Board1Pon2 Board1Pon2
	Board1Pon3 Board1Pon3
	Board1Pon4 Board1Pon4
	Board1Pon5 Board1Pon5
	Board1Pon6 Board1Pon6
	Board1Pon7 Board1Pon7
	Board1Pon8 Board1Pon8
}

type SnmpConfig struct {
	Ip        string `mapstructure:"ip"`
	Community string `mapstructure:"community"`
}

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port uint16 `mapstructure:"port"`
}

type OltConfig struct {
	BaseOID1        string `mapstructure:"base_oid_1"`
	BaseOID2        string `mapstructure:"base_oid_2"`
	OnuIDNameAllPon string `mapstructure:"onu_id_name"`
	OnuTypeAllPon   string `mapstructure:"onu_type"`
}

type Board1Pon1 struct {
	OnuIDNameOID       string `mapstructure:"onu_id_name"`
	OnuTypeOID         string `mapstructure:"onu_type"`
	OnuSerialNumberOID string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID      string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID      string `mapstructure:"onu_tx_power"`
	OnuStatusOID       string `mapstructure:"onu_status_id"`
}

type Board1Pon2 struct {
	OnuIDNameOID       string `mapstructure:"onu_id_name"`
	OnuTypeOID         string `mapstructure:"onu_type"`
	OnuSerialNumberOID string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID      string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID      string `mapstructure:"onu_tx_power"`
	OnuStatusOID       string `mapstructure:"onu_status_id"`
}

type Board1Pon3 struct {
	OnuIDNameOID       string `mapstructure:"onu_id_name"`
	OnuTypeOID         string `mapstructure:"onu_type"`
	OnuSerialNumberOID string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID      string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID      string `mapstructure:"onu_tx_power"`
	OnuStatusOID       string `mapstructure:"onu_status_id"`
}

type Board1Pon4 struct {
	OnuIDNameOID       string `mapstructure:"onu_id_name"`
	OnuTypeOID         string `mapstructure:"onu_type"`
	OnuSerialNumberOID string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID      string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID      string `mapstructure:"onu_tx_power"`
	OnuStatusOID       string `mapstructure:"onu_status_id"`
}

type Board1Pon5 struct {
	OnuIDNameOID       string `mapstructure:"onu_id_name"`
	OnuTypeOID         string `mapstructure:"onu_type"`
	OnuSerialNumberOID string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID      string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID      string `mapstructure:"onu_tx_power"`
	OnuStatusOID       string `mapstructure:"onu_status_id"`
}

type Board1Pon6 struct {
	OnuIDNameOID       string `mapstructure:"onu_id_name"`
	OnuTypeOID         string `mapstructure:"onu_type"`
	OnuSerialNumberOID string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID      string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID      string `mapstructure:"onu_tx_power"`
	OnuStatusOID       string `mapstructure:"onu_status_id"`
}

type Board1Pon7 struct {
	OnuIDNameOID       string `mapstructure:"onu_id_name"`
	OnuTypeOID         string `mapstructure:"onu_type"`
	OnuSerialNumberOID string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID      string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID      string `mapstructure:"onu_tx_power"`
	OnuStatusOID       string `mapstructure:"onu_status_id"`
}

type Board1Pon8 struct {
	OnuIDNameOID       string `mapstructure:"onu_id_name"`
	OnuTypeOID         string `mapstructure:"onu_type"`
	OnuSerialNumberOID string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID      string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID      string `mapstructure:"onu_tx_power"`
	OnuStatusOID       string `mapstructure:"onu_status_id"`
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
