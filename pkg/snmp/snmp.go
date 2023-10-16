package snmp

import (
	"fmt"
	"github.com/gosnmp/gosnmp"
	"github.com/sumitroajiprabowo/go-snmp-olt-c320/config"
	"time"
)

// SetupSnmpConnection  for olt c320 using snmp v2c and gosnmp library
func SetupSnmpConnection(config *config.Config) (*gosnmp.GoSNMP, error) {
	target := &gosnmp.GoSNMP{
		Target:    config.SnmpCfg.IpOlt,
		Port:      161,
		Community: config.SnmpCfg.Community,
		Version:   gosnmp.Version2c,
		Timeout:   time.Duration(2) * time.Second,
		//Logger:    log.New(os.Stdout, "", 0),
	}

	err := target.Connect()
	if err != nil {
		return nil, fmt.Errorf("connect: %w", err)
	}

	return target, nil
}
