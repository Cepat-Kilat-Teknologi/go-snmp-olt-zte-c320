package model

type ONUInfo struct {
	ID   string `json:"onu_id"`
	Name string `json:"name"`
}

type ONUInfoPerGTGO struct {
	GTGO         int    `json:"gtgo"`
	PON          int    `json:"pon"`
	ID           int    `json:"onu_id"`
	Name         string `json:"name"`
	OnuType      string `json:"onu_type"`
	SerialNumber string `json:"serial_number"`
	RXPower      string `json:"rx_power"`
	Status       string `json:"status"`
}

type ONUCustomerInfo struct {
	GTGO         int    `json:"gtgo"`
	PON          int    `json:"pon"`
	ID           int    `json:"onu_id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	OnuType      string `json:"onu_type"`
	SerialNumber string `json:"serial_number"`
	RXPower      string `json:"rx_power"`
	TXPower      string `json:"tx_power"`
	Status       string `json:"status"`
	IPAddress    string `json:"ip_address"`
}

type OnuID struct {
	GTGO int `json:"gtgo"`
	PON  int `json:"pon"`
	ID   int `json:"onu_id"`
}

type OnuOnlyID struct {
	ID int `json:"onu_id"`
}
