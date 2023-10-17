package snmp

import (
	"github.com/gosnmp/gosnmp"
)

type SNMPRepository interface {
	GetSnmp() *gosnmp.GoSNMP
	Connect() error
	GetIpOlt() (string, error)
	Walk(oid string, walkFunc func(pdu gosnmp.SnmpPDU) error) error
}

type snmpRepository struct {
	snmp *gosnmp.GoSNMP
}

func NewPonRepository(snmp *gosnmp.GoSNMP) SNMPRepository {
	return &snmpRepository{
		snmp: snmp,
	}
}

func (r *snmpRepository) GetSnmp() *gosnmp.GoSNMP {
	return r.snmp
}

func (r *snmpRepository) Connect() error {
	return r.snmp.Connect()
}

func (r *snmpRepository) GetIpOlt() (string, error) {
	return r.snmp.Target, nil
}

func (r *snmpRepository) Walk(oid string, walkFunc func(pdu gosnmp.SnmpPDU) error) error {
	return r.snmp.Walk(oid, walkFunc)
}
