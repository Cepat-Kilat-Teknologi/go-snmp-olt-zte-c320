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
	"sort"
	"strconv"
	"time"
)

type OnuUseCase interface {
	GetAllONUInfo(ctx context.Context) ([]model.ONUInfo, error)
	GetByPonID(ctx context.Context, ponID int) ([]model.ONUInformation, error)
	GetByGtGoIDAndPonID(ctx context.Context, gtGoID, ponID int) ([]model.ONUInfoPerGTGO, error)
}

type onuUsecase struct {
	snmpRepository snmp.RepositorySNMP
	cfg            *config.Config
}

func NewOnuUsecase(snmpRepository snmp.RepositorySNMP, cfg *config.Config) OnuUseCase {
	return &onuUsecase{
		snmpRepository: snmpRepository,
		cfg:            cfg,
	}
}

// GetAllONUInfo melakukan SNMP Walk, mengambil data ONUInfo, dan mengonversinya menjadi JSON
func (u *onuUsecase) GetAllONUInfo(ctx context.Context) ([]model.ONUInfo, error) {
	// Tentukan timeout menggunakan context
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	// OID dasar
	baseOID := u.cfg.OltCfg.BaseOID1

	// OID yang ingin gunakan untuk mengambil ONU ID
	onuOID := u.cfg.OltCfg.OnuIDNameAllPon

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

func (u *onuUsecase) GetByGtGoIDAndPonID(ctx context.Context, gtGoID, ponID int) ([]model.ONUInfoPerGTGO, error) {
	// Tentukan timeout menggunakan context
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	var baseOID string
	var onuIDNameOID string
	var onuTypeOID string
	var onuSerialNumberOID string
	var onuRxPowerOID string
	var onuStatusOID string

	switch gtGoID {

	case 1:
		switch ponID {
		case 1:
			baseOID = u.cfg.OltCfg.BaseOID1
			onuIDNameOID = u.cfg.Board1Pon1.OnuIDNameOID
			onuTypeOID = u.cfg.Board1Pon1.OnuTypeOID
			onuSerialNumberOID = u.cfg.Board1Pon1.OnuSerialNumberOID
			onuRxPowerOID = u.cfg.Board1Pon1.OnuRxPowerOID
			onuStatusOID = u.cfg.Board1Pon1.OnuStatusOID
		case 2:
			baseOID = u.cfg.OltCfg.BaseOID1
			onuIDNameOID = u.cfg.Board1Pon2.OnuIDNameOID
			onuTypeOID = u.cfg.Board1Pon2.OnuTypeOID
			onuSerialNumberOID = u.cfg.Board1Pon2.OnuSerialNumberOID
			onuRxPowerOID = u.cfg.Board1Pon2.OnuRxPowerOID
			onuStatusOID = u.cfg.Board1Pon2.OnuStatusOID
		case 3:
			baseOID = u.cfg.OltCfg.BaseOID1
			onuIDNameOID = u.cfg.Board1Pon3.OnuIDNameOID
			onuTypeOID = u.cfg.Board1Pon3.OnuTypeOID
			onuSerialNumberOID = u.cfg.Board1Pon3.OnuSerialNumberOID
			onuRxPowerOID = u.cfg.Board1Pon3.OnuRxPowerOID
			onuStatusOID = u.cfg.Board1Pon3.OnuStatusOID
		case 4:
			baseOID = u.cfg.OltCfg.BaseOID1
			onuIDNameOID = u.cfg.Board1Pon4.OnuIDNameOID
			onuTypeOID = u.cfg.Board1Pon4.OnuTypeOID
			onuSerialNumberOID = u.cfg.Board1Pon4.OnuSerialNumberOID
			onuRxPowerOID = u.cfg.Board1Pon4.OnuRxPowerOID
			onuStatusOID = u.cfg.Board1Pon4.OnuStatusOID
		case 5:
			baseOID = u.cfg.OltCfg.BaseOID1
			onuIDNameOID = u.cfg.Board1Pon5.OnuIDNameOID
			onuTypeOID = u.cfg.Board1Pon5.OnuTypeOID
			onuSerialNumberOID = u.cfg.Board1Pon5.OnuSerialNumberOID
			onuRxPowerOID = u.cfg.Board1Pon5.OnuRxPowerOID
			onuStatusOID = u.cfg.Board1Pon5.OnuStatusOID
		case 6:
			baseOID = u.cfg.OltCfg.BaseOID1
			onuIDNameOID = u.cfg.Board1Pon6.OnuIDNameOID
			onuTypeOID = u.cfg.Board1Pon6.OnuTypeOID
			onuSerialNumberOID = u.cfg.Board1Pon6.OnuSerialNumberOID
			onuRxPowerOID = u.cfg.Board1Pon6.OnuRxPowerOID
			onuStatusOID = u.cfg.Board1Pon6.OnuStatusOID
		case 7:
			baseOID = u.cfg.OltCfg.BaseOID1
			onuIDNameOID = u.cfg.Board1Pon7.OnuIDNameOID
			onuTypeOID = u.cfg.Board1Pon7.OnuTypeOID
			onuSerialNumberOID = u.cfg.Board1Pon7.OnuSerialNumberOID
			onuRxPowerOID = u.cfg.Board1Pon7.OnuRxPowerOID
			onuStatusOID = u.cfg.Board1Pon7.OnuStatusOID
		case 8:
			baseOID = u.cfg.OltCfg.BaseOID1
			onuIDNameOID = u.cfg.Board1Pon8.OnuIDNameOID
			onuTypeOID = u.cfg.Board1Pon8.OnuTypeOID
			onuSerialNumberOID = u.cfg.Board1Pon8.OnuSerialNumberOID
			onuRxPowerOID = u.cfg.Board1Pon8.OnuRxPowerOID
			onuStatusOID = u.cfg.Board1Pon8.OnuStatusOID
		default:
			baseOID = u.cfg.OltCfg.BaseOID1
			onuIDNameOID = u.cfg.Board1Pon1.OnuIDNameOID
			onuTypeOID = u.cfg.Board1Pon1.OnuTypeOID
			onuSerialNumberOID = u.cfg.Board1Pon1.OnuSerialNumberOID
			onuRxPowerOID = u.cfg.Board1Pon1.OnuRxPowerOID
			onuStatusOID = u.cfg.Board1Pon1.OnuStatusOID
		}
	case 2:
		switch ponID {
		case 1:
			baseOID = u.cfg.OltCfg.BaseOID1
			onuIDNameOID = u.cfg.Board1Pon1.OnuIDNameOID
			onuTypeOID = u.cfg.Board1Pon1.OnuTypeOID
			onuSerialNumberOID = u.cfg.Board1Pon1.OnuSerialNumberOID
			onuRxPowerOID = u.cfg.Board1Pon1.OnuRxPowerOID
			onuStatusOID = u.cfg.Board1Pon1.OnuStatusOID
		case 2:
			baseOID = u.cfg.OltCfg.BaseOID1
			onuIDNameOID = u.cfg.Board1Pon2.OnuIDNameOID
			onuTypeOID = u.cfg.Board1Pon2.OnuTypeOID
			onuSerialNumberOID = u.cfg.Board1Pon2.OnuSerialNumberOID
			onuRxPowerOID = u.cfg.Board1Pon2.OnuRxPowerOID
			onuStatusOID = u.cfg.Board1Pon2.OnuStatusOID
		case 3:
			baseOID = u.cfg.OltCfg.BaseOID1
			onuIDNameOID = u.cfg.Board1Pon3.OnuIDNameOID
			onuTypeOID = u.cfg.Board1Pon3.OnuTypeOID
			onuSerialNumberOID = u.cfg.Board1Pon3.OnuSerialNumberOID
			onuRxPowerOID = u.cfg.Board1Pon3.OnuRxPowerOID
			onuStatusOID = u.cfg.Board1Pon3.OnuStatusOID
		case 4:
			baseOID = u.cfg.OltCfg.BaseOID1
			onuIDNameOID = u.cfg.Board1Pon4.OnuIDNameOID
			onuTypeOID = u.cfg.Board1Pon4.OnuTypeOID
			onuSerialNumberOID = u.cfg.Board1Pon4.OnuSerialNumberOID
			onuRxPowerOID = u.cfg.Board1Pon4.OnuRxPowerOID
			onuStatusOID = u.cfg.Board1Pon4.OnuStatusOID
		case 5:
			baseOID = u.cfg.OltCfg.BaseOID1
			onuIDNameOID = u.cfg.Board1Pon5.OnuIDNameOID
			onuTypeOID = u.cfg.Board1Pon5.OnuTypeOID
			onuSerialNumberOID = u.cfg.Board1Pon5.OnuSerialNumberOID
			onuRxPowerOID = u.cfg.Board1Pon5.OnuRxPowerOID
			onuStatusOID = u.cfg.Board1Pon5.OnuStatusOID
		case 6:
			baseOID = u.cfg.OltCfg.BaseOID1
			onuIDNameOID = u.cfg.Board1Pon6.OnuIDNameOID
			onuTypeOID = u.cfg.Board1Pon6.OnuTypeOID
			onuSerialNumberOID = u.cfg.Board1Pon6.OnuSerialNumberOID
			onuRxPowerOID = u.cfg.Board1Pon6.OnuRxPowerOID
			onuStatusOID = u.cfg.Board1Pon6.OnuStatusOID
		case 7:
			baseOID = u.cfg.OltCfg.BaseOID1
			onuIDNameOID = u.cfg.Board1Pon7.OnuIDNameOID
			onuTypeOID = u.cfg.Board1Pon7.OnuTypeOID
			onuSerialNumberOID = u.cfg.Board1Pon7.OnuSerialNumberOID
			onuRxPowerOID = u.cfg.Board1Pon7.OnuRxPowerOID
			onuStatusOID = u.cfg.Board1Pon7.OnuStatusOID
		case 8:
			baseOID = u.cfg.OltCfg.BaseOID1
			onuIDNameOID = u.cfg.Board1Pon8.OnuIDNameOID
			onuTypeOID = u.cfg.Board1Pon8.OnuTypeOID
			onuSerialNumberOID = u.cfg.Board1Pon8.OnuSerialNumberOID
			onuRxPowerOID = u.cfg.Board1Pon8.OnuRxPowerOID
			onuStatusOID = u.cfg.Board1Pon8.OnuStatusOID
		default:
			baseOID = u.cfg.OltCfg.BaseOID1
			onuIDNameOID = u.cfg.Board1Pon1.OnuIDNameOID
			onuTypeOID = u.cfg.Board1Pon1.OnuTypeOID
			onuSerialNumberOID = u.cfg.Board1Pon1.OnuSerialNumberOID
			onuRxPowerOID = u.cfg.Board1Pon1.OnuRxPowerOID
			onuStatusOID = u.cfg.Board1Pon1.OnuStatusOID
		}
	}

	// Menggunakan SNMP Walk dengan timeout
	var onuInformationList []model.ONUInfoPerGTGO

	// Buat map untuk menyimpan hasil SNMP Walk
	snmpDataMap := make(map[string]gosnmp.SnmpPDU)

	err := u.snmpRepository.Walk(baseOID+onuIDNameOID, func(pdu gosnmp.SnmpPDU) error {
		// Simpan hasil SNMP Walk dalam map dengan ID sebagai kunci
		snmpDataMap[utils.ExtractONUID(pdu.Name)] = pdu
		return nil
	})

	if err != nil {
		panic(err)
	}

	// Lakukan pengolahan berdasarkan data SNMP yang tersimpan di map snmpDataMap
	for _, pdu := range snmpDataMap {
		onuInfo := model.ONUInfoPerGTGO{
			GTGO: gtGoID,
			PON:  ponID,
			ID:   utils.ExtractONUID(pdu.Name),
			Name: utils.ExtractName(pdu.Value),
		}

		// Lakukan pengolahan berdasarkan data SNMP yang tersimpan
		onuType, err := u.getONUType(ctx, onuTypeOID, onuInfo.ID)
		if err == nil {
			onuInfo.OnuType = onuType
		}

		onuSerialNumber, err := u.getSerialNumber(ctx, onuSerialNumberOID, onuInfo.ID)
		if err == nil {
			onuInfo.SerialNumber = onuSerialNumber
		}

		onuRXPower, err := u.getRxPower(ctx, onuRxPowerOID, onuInfo.ID)
		if err == nil {
			onuInfo.RXPower = onuRXPower
		}

		onuStatus, err := u.getStatus(ctx, onuStatusOID, onuInfo.ID)
		if err == nil {
			onuInfo.Status = onuStatus
		}

		onuInformationList = append(onuInformationList, onuInfo)
	}

	// Urutkan berdasarkan ID
	sort.Slice(onuInformationList, func(i, j int) bool {
		onuID1, _ := strconv.Atoi(onuInformationList[i].ID)
		onuID2, _ := strconv.Atoi(onuInformationList[j].ID)
		return onuID1 < onuID2
	})
	return onuInformationList, nil
}

func (u *onuUsecase) GetByPonID(ctx context.Context, ponID int) ([]model.ONUInformation, error) {
	// Tentukan timeout menggunakan context
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	var baseOID string
	var onuIDNameOID string
	var onuTypeOID string
	var onuSerialNumberOID string
	var onuRxPowerOID string
	//var onuTxPowerOID string
	var onuStatusOID string

	switch ponID {
	case 1:
		baseOID = u.cfg.OltCfg.BaseOID1
		onuIDNameOID = u.cfg.Board1Pon1.OnuIDNameOID
		onuTypeOID = u.cfg.Board1Pon1.OnuTypeOID
		onuSerialNumberOID = u.cfg.Board1Pon1.OnuSerialNumberOID
		onuRxPowerOID = u.cfg.Board1Pon1.OnuRxPowerOID
		onuStatusOID = u.cfg.Board1Pon1.OnuStatusOID
	case 2:
		baseOID = u.cfg.OltCfg.BaseOID1
		onuIDNameOID = u.cfg.Board1Pon2.OnuIDNameOID
		onuTypeOID = u.cfg.Board1Pon2.OnuTypeOID
		onuSerialNumberOID = u.cfg.Board1Pon2.OnuSerialNumberOID
		onuRxPowerOID = u.cfg.Board1Pon2.OnuRxPowerOID
		onuStatusOID = u.cfg.Board1Pon2.OnuStatusOID
	case 3:
		baseOID = u.cfg.OltCfg.BaseOID1
		onuIDNameOID = u.cfg.Board1Pon3.OnuIDNameOID
		onuTypeOID = u.cfg.Board1Pon3.OnuTypeOID
		onuSerialNumberOID = u.cfg.Board1Pon3.OnuSerialNumberOID
		onuRxPowerOID = u.cfg.Board1Pon3.OnuRxPowerOID
		onuStatusOID = u.cfg.Board1Pon3.OnuStatusOID
	case 4:
		baseOID = u.cfg.OltCfg.BaseOID1
		onuIDNameOID = u.cfg.Board1Pon4.OnuIDNameOID
		onuTypeOID = u.cfg.Board1Pon4.OnuTypeOID
		onuSerialNumberOID = u.cfg.Board1Pon4.OnuSerialNumberOID
		onuRxPowerOID = u.cfg.Board1Pon4.OnuRxPowerOID
		onuStatusOID = u.cfg.Board1Pon4.OnuStatusOID
	case 5:
		baseOID = u.cfg.OltCfg.BaseOID1
		onuIDNameOID = u.cfg.Board1Pon5.OnuIDNameOID
		onuTypeOID = u.cfg.Board1Pon5.OnuTypeOID
		onuSerialNumberOID = u.cfg.Board1Pon5.OnuSerialNumberOID
		onuRxPowerOID = u.cfg.Board1Pon5.OnuRxPowerOID
		onuStatusOID = u.cfg.Board1Pon5.OnuStatusOID
	case 6:
		baseOID = u.cfg.OltCfg.BaseOID1
		onuIDNameOID = u.cfg.Board1Pon6.OnuIDNameOID
		onuTypeOID = u.cfg.Board1Pon6.OnuTypeOID
		onuSerialNumberOID = u.cfg.Board1Pon6.OnuSerialNumberOID
		onuRxPowerOID = u.cfg.Board1Pon6.OnuRxPowerOID
		onuStatusOID = u.cfg.Board1Pon6.OnuStatusOID
	case 7:
		baseOID = u.cfg.OltCfg.BaseOID1
		onuIDNameOID = u.cfg.Board1Pon7.OnuIDNameOID
		onuTypeOID = u.cfg.Board1Pon7.OnuTypeOID
		onuSerialNumberOID = u.cfg.Board1Pon7.OnuSerialNumberOID
		onuRxPowerOID = u.cfg.Board1Pon7.OnuRxPowerOID
		onuStatusOID = u.cfg.Board1Pon7.OnuStatusOID
	case 8:
		baseOID = u.cfg.OltCfg.BaseOID1
		onuIDNameOID = u.cfg.Board1Pon8.OnuIDNameOID
		onuTypeOID = u.cfg.Board1Pon8.OnuTypeOID
		onuSerialNumberOID = u.cfg.Board1Pon8.OnuSerialNumberOID
		onuRxPowerOID = u.cfg.Board1Pon8.OnuRxPowerOID
		onuStatusOID = u.cfg.Board1Pon8.OnuStatusOID
	default:
		baseOID = u.cfg.OltCfg.BaseOID1
		onuIDNameOID = u.cfg.Board1Pon1.OnuIDNameOID
		onuTypeOID = u.cfg.Board1Pon1.OnuTypeOID
		onuSerialNumberOID = u.cfg.Board1Pon1.OnuSerialNumberOID
		onuRxPowerOID = u.cfg.Board1Pon1.OnuRxPowerOID
		onuStatusOID = u.cfg.Board1Pon1.OnuStatusOID
	}

	// Menggunakan SNMP Walk dengan timeout
	var onuInformationList []model.ONUInformation

	// Buat map untuk menyimpan hasil SNMP Walk
	snmpDataMap := make(map[string]gosnmp.SnmpPDU)

	err := u.snmpRepository.Walk(baseOID+onuIDNameOID, func(pdu gosnmp.SnmpPDU) error {
		// Simpan hasil SNMP Walk dalam map dengan ID sebagai kunci
		snmpDataMap[utils.ExtractONUID(pdu.Name)] = pdu
		return nil
	})

	if err != nil {
		panic(err)
	}

	// Lakukan pengolahan berdasarkan data SNMP yang tersimpan di map snmpDataMap
	for _, pdu := range snmpDataMap {
		onuInfo := model.ONUInformation{
			PON:  ponID,
			ID:   utils.ExtractONUID(pdu.Name),
			Name: utils.ExtractName(pdu.Value),
		}

		// Lakukan pengolahan berdasarkan data SNMP yang tersimpan
		onuType, err := u.getONUType(ctx, onuTypeOID, onuInfo.ID)
		if err == nil {
			onuInfo.OnuType = onuType
		}

		onuSerialNumber, err := u.getSerialNumber(ctx, onuSerialNumberOID, onuInfo.ID)
		if err == nil {
			onuInfo.SerialNumber = onuSerialNumber
		}

		onuRXPower, err := u.getRxPower(ctx, onuRxPowerOID, onuInfo.ID)
		if err == nil {
			onuInfo.RXPower = onuRXPower
		}

		//onuTXPower, err := u.getTxPower(ctx, onuTxPowerOID, onuInfo.ID)
		//if err == nil {
		//	onuInfo.TXPower = onuTXPower
		//}

		onuStatus, err := u.getStatus(ctx, onuStatusOID, onuInfo.ID)
		if err == nil {
			onuInfo.Status = onuStatus
		}

		onuInformationList = append(onuInformationList, onuInfo)
	}

	// Urutkan berdasarkan ID
	sort.Slice(onuInformationList, func(i, j int) bool {
		onuID1, _ := strconv.Atoi(onuInformationList[i].ID)
		onuID2, _ := strconv.Atoi(onuInformationList[j].ID)
		return onuID1 < onuID2
	})

	return onuInformationList, nil
}

func (u *onuUsecase) getONUType(ctx context.Context, onuTypeOID, onuID string) (string, error) {
	var onuType string

	// Gunakan context.WithTimeout untuk mengatur batasan waktu pada operasi SNMP Walk
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	// OID dasar
	baseOID := u.cfg.OltCfg.BaseOID2

	// Lakukan SNMP Walk untuk ONU Type
	err := u.snmpRepository.Walk(baseOID+onuTypeOID+"."+onuID, func(pdu gosnmp.SnmpPDU) error {
		onuType = utils.ExtractName(pdu.Value)
		return nil
	})

	if err != nil {
		return "", err
	}

	return onuType, nil
}

func (u *onuUsecase) getSerialNumber(ctx context.Context, onuSerialNumberOID, onuID string) (string, error) {
	var onuSerialNumber string

	// Gunakan context.WithTimeout untuk mengatur batasan waktu pada operasi SNMP Walk
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	// OID dasar
	baseOID := u.cfg.OltCfg.BaseOID1

	// Lakukan SNMP Walk untuk ONU Serial Number
	err := u.snmpRepository.Walk(baseOID+onuSerialNumberOID+"."+onuID, func(pdu gosnmp.SnmpPDU) error {
		onuSerialNumber = utils.ExtractSerialNumber(pdu.Value)
		return nil
	})

	if err != nil {
		return "", err
	}

	return onuSerialNumber, nil
}

func (u *onuUsecase) getRxPower(ctx context.Context, onuRxPowerOID, onuID string) (string, error) {
	var onuRxPower string

	// Gunakan context.WithTimeout untuk mengatur batasan waktu pada operasi SNMP Walk
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	// OID dasar
	baseOID := u.cfg.OltCfg.BaseOID1

	// Lakukan SNMP Walk untuk ONU Serial Number
	err := u.snmpRepository.Walk(baseOID+onuRxPowerOID+"."+onuID+"."+"1", func(pdu gosnmp.SnmpPDU) error {

		// convert interface to string
		onuRxPower, _ = utils.ConvertAndMultiply(pdu.Value)
		return nil
	})

	if err != nil {
		return "", err
	}

	return onuRxPower, nil
}

func (u *onuUsecase) getTxPower(ctx context.Context, onuTxPowerOID, onuID string) (string, error) {
	var onuTxPower string

	// Gunakan context.WithTimeout untuk mengatur batasan waktu pada operasi SNMP Walk
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	// OID dasar
	baseOID := u.cfg.OltCfg.BaseOID2

	// Lakukan SNMP Walk untuk ONU Serial Number
	err := u.snmpRepository.Walk(baseOID+onuTxPowerOID+"."+onuID+"."+"1", func(pdu gosnmp.SnmpPDU) error {
		onuTxPower, _ = utils.ConvertAndMultiply(pdu.Value)
		return nil
	})

	if err != nil {
		return "", err
	}

	return onuTxPower, nil
}

func (u *onuUsecase) getStatus(ctx context.Context, onuStatusOID, onuID string) (string, error) {
	var onuStatus string

	// Gunakan context.WithTimeout untuk mengatur batasan waktu pada operasi SNMP Walk
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	// OID dasar
	baseOID := u.cfg.OltCfg.BaseOID1

	// Lakukan SNMP Walk untuk ONU Serial Number
	err := u.snmpRepository.Walk(baseOID+onuStatusOID+"."+onuID, func(pdu gosnmp.SnmpPDU) error {
		onuStatus = utils.ExtractAndGetStatus(pdu.Value)
		return nil
	})

	if err != nil {
		return "", err
	}

	return onuStatus, nil
}
