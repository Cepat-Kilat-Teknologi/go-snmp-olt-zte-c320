package snmp

import (
	"fmt"
	"github.com/gosnmp/gosnmp"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/config"
	"log"
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
		AppOpts: make(map[string]interface{}),
	}

	err := target.Connect()
	if err != nil {
		return nil, fmt.Errorf("gagal terhubung: %w", err)
	}

	// Function handles for collecting metrics on query latencies.
	var sent time.Time
	target.OnSent = func(x *gosnmp.GoSNMP) {
		sent = time.Now()
	}
	target.OnRecv = func(x *gosnmp.GoSNMP) {
		log.Println("Query latency in seconds:", time.Since(sent).Seconds())
	}
	target.OnFinish = func(x *gosnmp.GoSNMP) {
		log.Println("Total query latency in seconds:", time.Since(sent).Seconds())
	}

	return target, nil
}
