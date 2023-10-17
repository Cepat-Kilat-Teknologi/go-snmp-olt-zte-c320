package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/gosnmp/gosnmp"
	"github.com/sumitroajiprabowo/go-snmp-olt-c320/config"
	"github.com/sumitroajiprabowo/go-snmp-olt-c320/internal/model"
	"github.com/sumitroajiprabowo/go-snmp-olt-c320/internal/repository/snmp"
	"github.com/sumitroajiprabowo/go-snmp-olt-c320/pkg/utils"
	"time"
)

type PonUseCase interface {
	GetONUInfo(ctx context.Context) ([]model.ONUInfo, error)
}

type ponUsecase struct {
	snmpRepository snmp.SNMPRepository
	cfg            *config.Config
}

func NewPonUsecase(snmpRepository snmp.SNMPRepository, cfg *config.Config) PonUseCase {
	return &ponUsecase{
		snmpRepository: snmpRepository,
		cfg:            cfg,
	}
}

// GetONUInfo melakukan SNMP Walk, mengambil data ONUInfo, dan mengonversinya menjadi JSON
func (u *ponUsecase) GetONUInfo(ctx context.Context) ([]model.ONUInfo, error) {
	// Tentukan timeout menggunakan context
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	// OID dasar Anda
	baseOID := u.cfg.OltCfg.BaseOID

	// OID yang Anda ingin gunakan untuk mengambil ONU ID
	onuOID := u.cfg.OltCfg.OnuIDName

	// Menggunakan SNMP Walk dengan timeout
	var onuInfoList []model.ONUInfo

	err := u.snmpRepository.Walk(baseOID+onuOID, func(pdu gosnmp.SnmpPDU) error {
		var onuInfo model.ONUInfo
		onuInfo.ID = utils.ExtractONUID(pdu.Name)
		onuInfo.Name = utils.ExtractName(pdu.Value)
		onuInfoList = append(onuInfoList, onuInfo)
		return nil
	})

	if err != nil {
		// Periksa apakah terjadi error atau context dibatalkan
		if errors.Is(err, context.Canceled) {
			return nil, fmt.Errorf("SNMP operation canceled: %w", err)
		} else if errors.Is(context.DeadlineExceeded, err) {
			return nil, fmt.Errorf("SNMP operation timed out: %w", err)
		}
		return nil, fmt.Errorf("failed to walk OID: %w", err)
	}

	return onuInfoList, nil

}
