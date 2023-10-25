package repository

import (
	"github.com/gosnmp/gosnmp"
)

type SnmpInterface interface {
	Get(oids []string) (result *gosnmp.SnmpPacket, err error)
	Walk(oid string, walkFunc func(pdu gosnmp.SnmpPDU) error) error
	BulkWalk(oid string, walkFunc func(pdu gosnmp.SnmpPDU) error) error
	BulkWalkAll(oid string) ([]gosnmp.SnmpPDU, error)
}

type snmpRepository struct {
	snmp *gosnmp.GoSNMP
}

func NewPonRepository(snmp *gosnmp.GoSNMP) SnmpInterface {
	return &snmpRepository{
		snmp: snmp,
	}
}

func (r *snmpRepository) Get(oids []string) (result *gosnmp.SnmpPacket, err error) {
	return r.snmp.Get(oids)
}

func (r *snmpRepository) Walk(oid string, walkFunc func(pdu gosnmp.SnmpPDU) error) error {
	return r.snmp.Walk(oid, walkFunc)
}

func (r *snmpRepository) BulkWalk(oid string, walkFunc func(pdu gosnmp.SnmpPDU) error) error {
	return r.snmp.BulkWalk(oid, walkFunc)
}

func (r *snmpRepository) BulkWalkAll(oid string) ([]gosnmp.SnmpPDU, error) {
	return r.snmp.BulkWalkAll(oid)
}
