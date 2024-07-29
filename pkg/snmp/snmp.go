package snmp

import (
	"fmt"
	"github.com/gosnmp/gosnmp"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/config"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/internal/utils"
	"log"
	"os"
	"time"
)

var (
	snmpHost      string
	snmpPort      uint16
	snmpCommunity string
	logSnmp       gosnmp.Logger
)

// SetupSnmpConnection is a function to set up snmp connection
func SetupSnmpConnection(config *config.Config) (*gosnmp.GoSNMP, error) {

	if os.Getenv("APP_ENV") == "development" || os.Getenv("APP_ENV") == "production" {
		snmpHost = os.Getenv("SNMP_HOST")
		snmpPort = utils.ConvertStringToUint16(os.Getenv("SNMP_PORT"))
		snmpCommunity = os.Getenv("SNMP_COMMUNITY")
		logSnmp = gosnmp.Logger{}
	} else {
		snmpHost = config.SnmpCfg.Ip
		snmpPort = config.SnmpCfg.Port
		snmpCommunity = config.SnmpCfg.Community
		logSnmp = gosnmp.NewLogger(log.New(os.Stdout, "", 0))
	}

	target := &gosnmp.GoSNMP{
		Target:    snmpHost,
		Port:      snmpPort,
		Community: snmpCommunity,
		Version:   gosnmp.Version2c,
		Timeout:   time.Duration(30) * time.Second,
		//Retries:   3
		Logger: logSnmp,
	}

	err := target.Connect()
	if err != nil {
		return nil, fmt.Errorf("gagal terhubung: %w", err)
	}

	return target, nil
}
