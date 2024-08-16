package model

type OltConfig struct {
	BaseOID                   string
	OnuIDNameOID              string
	OnuTypeOID                string
	OnuSerialNumberOID        string
	OnuRxPowerOID             string
	OnuTxPowerOID             string
	OnuStatusOID              string
	OnuIPAddressOID           string
	OnuDescriptionOID         string
	OnuLastOnlineOID          string
	OnuLastOfflineOID         string
	OnuLastOfflineReasonOID   string
	OnuGponOpticalDistanceOID string
}

type ONUInfo struct {
	ID   string `json:"onu_id"`
	Name string `json:"name"`
}

type ONUInfoPerBoard struct {
	Board        int    `json:"board"`
	PON          int    `json:"pon"`
	ID           int    `json:"onu_id"`
	Name         string `json:"name"`
	OnuType      string `json:"onu_type"`
	SerialNumber string `json:"serial_number"`
	RXPower      string `json:"rx_power"`
	Status       string `json:"status"`
}

type ONUCustomerInfo struct {
	Board                int    `json:"board"`
	PON                  int    `json:"pon"`
	ID                   int    `json:"onu_id"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	OnuType              string `json:"onu_type"`
	SerialNumber         string `json:"serial_number"`
	RXPower              string `json:"rx_power"`
	TXPower              string `json:"tx_power"`
	Status               string `json:"status"`
	IPAddress            string `json:"ip_address"`
	LastOnline           string `json:"last_online"`
	LastOffline          string `json:"last_offline"`
	Uptime               string `json:"uptime"`
	LastDownTimeDuration string `json:"last_down_time_duration"`
	LastOfflineReason    string `json:"offline_reason"`
	GponOpticalDistance  string `json:"gpon_optical_distance"`
}

type OnuID struct {
	Board int `json:"board"`
	PON   int `json:"pon"`
	ID    int `json:"onu_id"`
}

type OnuOnlyID struct {
	ID int `json:"onu_id"`
}

type SNMPWalkTask struct {
	BaseOID   string
	TargetOID string
	BoardID   int
	PON       int
}

type OnuSerialNumber struct {
	Board        int    `json:"board"`
	PON          int    `json:"pon"`
	ID           int    `json:"onu_id"`
	SerialNumber string `json:"serial_number"`
}
