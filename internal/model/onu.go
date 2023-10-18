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
