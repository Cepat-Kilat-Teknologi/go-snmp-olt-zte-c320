package snmp

import (
	"fmt"
	"github.com/gosnmp/gosnmp"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/config"
	"time"
)

// SetupSnmpConnection is a function to setup snmp connection
func SetupSnmpConnection(config *config.Config) (*gosnmp.GoSNMP, error) {
	target := &gosnmp.GoSNMP{
		Target:    config.SnmpCfg.Ip,
		Port:      config.SnmpCfg.Port,
		Community: config.SnmpCfg.Community,
		Version:   gosnmp.Version2c,
		Timeout:   time.Duration(300) * time.Second,
		//Logger:    gosnmp.NewLogger(log.New(os.Stdout, "", 0)),
	}

	err := target.Connect()
	if err != nil {
		return nil, fmt.Errorf("gagal terhubung: %w", err)
	}

	return target, nil
}
