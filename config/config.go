package config

import (
	"errors"
	"github.com/spf13/viper"
)

type Config struct {
	SnmpCfg    SnmpConfig
	RedisCfg   RedisConfig
	OltCfg     OltConfig
	Board1Pon1 Board1Pon1
	Board1Pon2 Board1Pon2
	Board1Pon3 Board1Pon3
	Board1Pon4 Board1Pon4
	Board1Pon5 Board1Pon5
	Board1Pon6 Board1Pon6
	Board1Pon7 Board1Pon7
	Board1Pon8 Board1Pon8
	Board2Pon1 Board2Pon1
	Board2Pon2 Board2Pon2
	Board2Pon3 Board2Pon3
	Board2Pon4 Board2Pon4
	Board2Pon5 Board2Pon5
	Board2Pon6 Board2Pon6
	Board2Pon7 Board2Pon7
	Board2Pon8 Board2Pon8
}

type SnmpConfig struct {
	Ip        string `mapstructure:"ip"`
	Port      uint16 `mapstructure:"port"`
	Community string `mapstructure:"community"`
}

type RedisConfig struct {
	Host               string `mapstructure:"host"`
	Port               string `mapstructure:"port"`
	Password           string `mapstructure:"password"`
	DB                 int    `mapstructure:"db"`
	DefaultDB          int    `mapstructure:"default_db"`
	MinIdleConnections int    `mapstructure:"min_idle_connections"`
	PoolSize           int    `mapstructure:"pool_size"`
	PoolTimeout        int    `mapstructure:"pool_timeout"`
}

type OltConfig struct {
	BaseOID1        string `mapstructure:"base_oid_1"`
	BaseOID2        string `mapstructure:"base_oid_2"`
	OnuIDNameAllPon string `mapstructure:"onu_id_name"`
	OnuTypeAllPon   string `mapstructure:"onu_type"`
}

type Board1Pon1 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

type Board1Pon2 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

type Board1Pon3 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

type Board1Pon4 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

type Board1Pon5 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

type Board1Pon6 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

type Board1Pon7 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

type Board1Pon8 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

type Board2Pon1 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

type Board2Pon2 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

type Board2Pon3 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

type Board2Pon4 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

type Board2Pon5 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

type Board2Pon6 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

type Board2Pon7 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

type Board2Pon8 struct {
	OnuIDNameOID              string `mapstructure:"onu_id_name"`
	OnuTypeOID                string `mapstructure:"onu_type"`
	OnuSerialNumberOID        string `mapstructure:"onu_serial_number"`
	OnuRxPowerOID             string `mapstructure:"onu_rx_power"`
	OnuTxPowerOID             string `mapstructure:"onu_tx_power"`
	OnuStatusOID              string `mapstructure:"onu_status_id"`
	OnuIPAddressOID           string `mapstructure:"onu_ip_address"`
	OnuDescriptionOID         string `mapstructure:"onu_description"`
	OnuLastOnlineOID          string `mapstructure:"onu_last_online_time"`
	OnuLastOfflineOID         string `mapstructure:"onu_last_offline_time"`
	OnuLastOfflineReasonOID   string `mapstructure:"onu_last_offline_reason"`
	OnuGponOpticalDistanceOID string `mapstructure:"onu_gpon_optical_distance"`
}

// LoadConfig file from given path using viper
func LoadConfig(filename string) (*Config, error) {

	// Initialize viper
	v := viper.New()

	// Set config file name
	v.SetConfigName(filename)

	// Set config path in current directory
	v.AddConfigPath(".")

	// Allow environment variables to override config
	v.AutomaticEnv()

	// Read config file
	if err := v.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError // Initialize config file not found error
		if errors.As(err, &configFileNotFoundError) {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	var cfg Config // Initialize config variable

	// Unmarshal config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
