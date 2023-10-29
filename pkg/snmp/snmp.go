package snmp

import (
	"context"
	"fmt"
	"github.com/gosnmp/gosnmp"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/config"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/internal/utils"
	"os"
	"time"
)

var (
	snmpHost      string
	snmpPort      uint16
	snmpCommunity string
)

// SetupSnmpConnection is a function to set up snmp connection
func SetupSnmpConnection(ctx context.Context, config *config.Config) (*gosnmp.GoSNMP, error) {

	if os.Getenv("APP_ENV") == "development" || os.Getenv("APP_ENV") == "production" {
		snmpHost = os.Getenv("SNMP_HOST")
		snmpPort = utils.ConvertStringToUint16(os.Getenv("SNMP_PORT"))
		snmpCommunity = os.Getenv("SNMP_COMMUNITY")
	} else {
		snmpHost = config.SnmpCfg.Ip
		snmpPort = config.SnmpCfg.Port
		snmpCommunity = config.SnmpCfg.Community
	}

	target := &gosnmp.GoSNMP{
		Target:    snmpHost,
		Port:      snmpPort,
		Community: snmpCommunity,
		Context:   ctx,
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
