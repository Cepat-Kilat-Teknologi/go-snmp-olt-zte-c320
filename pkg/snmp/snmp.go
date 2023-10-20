package snmp

import (
	"fmt"
	"github.com/gosnmp/gosnmp"
	"github.com/megadata-dev/go-snmp-olt-c320/config"
	"time"
)

// SetupSnmpConnection digunakan untuk membuat koneksi SNMP dan mengembalikan objek *gosnmp.GoSNMP
func SetupSnmpConnection(config *config.Config) (*gosnmp.GoSNMP, error) {
	target := &gosnmp.GoSNMP{
		Target:    config.SnmpCfg.Ip,
		Port:      161,
		Community: config.SnmpCfg.Community,
		Version:   gosnmp.Version2c,
		Timeout:   time.Duration(300) * time.Second,
		Retries:   3,
	}

	err := target.Connect()
	if err != nil {
		return nil, fmt.Errorf("gagal terhubung: %w", err)
	}

	return target, nil
}
