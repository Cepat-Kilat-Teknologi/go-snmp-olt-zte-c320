package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/gosnmp/gosnmp"
	"github.com/megadata-dev/go-snmp-olt-c320/config"
	"github.com/megadata-dev/go-snmp-olt-c320/internal/model"
	"github.com/megadata-dev/go-snmp-olt-c320/internal/repository"
	"github.com/megadata-dev/go-snmp-olt-c320/pkg/utils"
	"sort"
	"strconv"
	"time"
)

type OnuUseCase interface {
	GetByGtGoIDAndPonID(ctx context.Context, gtGoID, ponID int) ([]model.ONUInfoPerGTGO, error)
	GetByGtGoIDPonIDAndOnuID(ctx context.Context, gtGoID, ponID, onuID int) (model.ONUCustomerInfo, error)
	GetEmptyOnuID(ctx context.Context, gtGoID, ponID int) ([]model.OnuID, error)
	GetByGtGoIDAndPonIDWithPagination(ctx context.Context, gtGoID, ponID, page, pageSize int) (
		[]model.ONUInfoPerGTGO, int,
	)
}

type onuUsecase struct {
	snmpRepository  repository.RepositorySNMP
	redisRepository repository.OnuRedisRepo
	cfg             *config.Config
}

func NewOnuUsecase(
	snmpRepository repository.RepositorySNMP, redisRepository repository.OnuRedisRepo, cfg *config.Config,
) OnuUseCase {
	return &onuUsecase{
		snmpRepository:  snmpRepository,
		redisRepository: redisRepository,
		cfg:             cfg,
	}
}

func (u *onuUsecase) GetByGtGoIDAndPonID(ctx context.Context, gtGoID, ponID int) ([]model.ONUInfoPerGTGO, error) {

	ctx, cancel := context.WithTimeout(ctx, time.Second*30) // Create context with timeout 30 seconds
	defer cancel()                                          // Cancel context when function is done

	var (
		baseOID            string // Base OID variable
		onuIDNameOID       string // ONU ID Name OID variable
		onuTypeOID         string // ONU Type OID variable
		onuSerialNumberOID string // ONU Serial Number OID variable
		onuRxPowerOID      string // ONU RX Power OID variable
		onuStatusOID       string // ONU Status OID variable
	)

	// Determine base OID and other OID based on GTGO ID and PON ID
	switch gtGoID {
	case 0: // GTGO 0
		switch ponID {
		case 1: // PON 1
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon1.OnuIDNameOID             // ONU ID Name OID variable get from config
			onuTypeOID = u.cfg.Board1Pon1.OnuTypeOID                 // ONU Type OID variable get from config
			onuSerialNumberOID = u.cfg.Board1Pon1.OnuSerialNumberOID // ONU Serial Number OID variable get from config
			onuRxPowerOID = u.cfg.Board1Pon1.OnuRxPowerOID           // ONU RX Power OID variable get from config
			onuStatusOID = u.cfg.Board1Pon1.OnuStatusOID             // ONU Status OID variable get from config
		case 2: // PON 2
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon2.OnuIDNameOID             // ONU ID Name OID variable get from config
			onuTypeOID = u.cfg.Board1Pon2.OnuTypeOID                 // ONU Type OID variable get from config
			onuSerialNumberOID = u.cfg.Board1Pon2.OnuSerialNumberOID // ONU Serial Number OID variable get from config
			onuRxPowerOID = u.cfg.Board1Pon2.OnuRxPowerOID           // ONU RX Power OID variable get from config
			onuStatusOID = u.cfg.Board1Pon2.OnuStatusOID             // ONU Status OID variable get from config
		case 3: // PON 3
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon3.OnuIDNameOID             // ONU ID Name OID variable get from config
			onuTypeOID = u.cfg.Board1Pon3.OnuTypeOID                 // ONU Type OID variable get from config
			onuSerialNumberOID = u.cfg.Board1Pon3.OnuSerialNumberOID // ONU Serial Number OID variable get from config
			onuRxPowerOID = u.cfg.Board1Pon3.OnuRxPowerOID           // ONU RX Power OID variable get from config
			onuStatusOID = u.cfg.Board1Pon3.OnuStatusOID             // ONU Status OID variable get from config
		case 4: // PON 4
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon4.OnuIDNameOID             // ONU ID Name OID variable get from config
			onuTypeOID = u.cfg.Board1Pon4.OnuTypeOID                 // ONU Type OID variable get from config
			onuSerialNumberOID = u.cfg.Board1Pon4.OnuSerialNumberOID // ONU Serial Number OID variable get from config
			onuRxPowerOID = u.cfg.Board1Pon4.OnuRxPowerOID           // ONU RX Power OID variable get from config
			onuStatusOID = u.cfg.Board1Pon4.OnuStatusOID             // ONU Status OID variable get from config
		case 5: // PON 5
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon5.OnuIDNameOID             // ONU ID Name OID variable get from config
			onuTypeOID = u.cfg.Board1Pon5.OnuTypeOID                 // ONU Type OID variable get from config
			onuSerialNumberOID = u.cfg.Board1Pon5.OnuSerialNumberOID // ONU Serial Number OID variable get from config
			onuRxPowerOID = u.cfg.Board1Pon5.OnuRxPowerOID           // ONU RX Power OID variable get from config
			onuStatusOID = u.cfg.Board1Pon5.OnuStatusOID             // ONU Status OID variable get from config
		case 6: // PON 6
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon6.OnuIDNameOID             // ONU ID Name OID variable get from config
			onuTypeOID = u.cfg.Board1Pon6.OnuTypeOID                 // ONU Type OID variable get from config
			onuSerialNumberOID = u.cfg.Board1Pon6.OnuSerialNumberOID // ONU Serial Number OID variable get from config
			onuRxPowerOID = u.cfg.Board1Pon6.OnuRxPowerOID           // ONU RX Power OID variable get from config
			onuStatusOID = u.cfg.Board1Pon6.OnuStatusOID             // ONU Status OID variable get from config
		case 7: // PON 7
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon7.OnuIDNameOID             // ONU ID Name OID variable get from config
			onuTypeOID = u.cfg.Board1Pon7.OnuTypeOID                 // ONU Type OID variable get from config
			onuSerialNumberOID = u.cfg.Board1Pon7.OnuSerialNumberOID // ONU Serial Number OID variable get from config
			onuRxPowerOID = u.cfg.Board1Pon7.OnuRxPowerOID           // ONU RX Power OID variable get from config
			onuStatusOID = u.cfg.Board1Pon7.OnuStatusOID             // ONU Status OID variable get from config
		case 8: // PON 8
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon8.OnuIDNameOID             // ONU ID Name OID variable get from config
			onuTypeOID = u.cfg.Board1Pon8.OnuTypeOID                 // ONU Type OID variable get from config
			onuSerialNumberOID = u.cfg.Board1Pon8.OnuSerialNumberOID // ONU Serial Number OID variable get from config
			onuRxPowerOID = u.cfg.Board1Pon8.OnuRxPowerOID           // ONU RX Power OID variable get from config
			onuStatusOID = u.cfg.Board1Pon8.OnuStatusOID             // ONU Status OID variable get from config
		default: // Invalid PON ID
			return nil, errors.New("invalid PON ID") // Return error
		}
	case 1: // GTGO 1
		switch ponID {
		case 1: // PON 1
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon1.OnuIDNameOID             // ONU ID Name OID variable get from config
			onuTypeOID = u.cfg.Board1Pon1.OnuTypeOID                 // ONU Type OID variable get from config
			onuSerialNumberOID = u.cfg.Board1Pon1.OnuSerialNumberOID // ONU Serial Number OID variable get from config
			onuRxPowerOID = u.cfg.Board1Pon1.OnuRxPowerOID           // ONU RX Power OID variable get from config
			onuStatusOID = u.cfg.Board1Pon1.OnuStatusOID             // ONU Status OID variable get from config
		case 2: // PON 2
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon2.OnuIDNameOID             // ONU ID Name OID variable get from config
			onuTypeOID = u.cfg.Board1Pon2.OnuTypeOID                 // ONU Type OID variable get from config
			onuSerialNumberOID = u.cfg.Board1Pon2.OnuSerialNumberOID // ONU Serial Number OID variable get from config
			onuRxPowerOID = u.cfg.Board1Pon2.OnuRxPowerOID           // ONU RX Power OID variable get from config
			onuStatusOID = u.cfg.Board1Pon2.OnuStatusOID             // ONU Status OID variable get from config
		case 3: // PON 3
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon3.OnuIDNameOID             // ONU ID Name OID variable get from config
			onuTypeOID = u.cfg.Board1Pon3.OnuTypeOID                 // ONU Type OID variable get from config
			onuSerialNumberOID = u.cfg.Board1Pon3.OnuSerialNumberOID // ONU Serial Number OID variable get from config
			onuRxPowerOID = u.cfg.Board1Pon3.OnuRxPowerOID           // ONU RX Power OID variable get from config
			onuStatusOID = u.cfg.Board1Pon3.OnuStatusOID             // ONU Status OID variable get from config
		case 4: // PON 4
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon4.OnuIDNameOID             // ONU ID Name OID variable get from config
			onuTypeOID = u.cfg.Board1Pon4.OnuTypeOID                 // ONU Type OID variable get from config
			onuSerialNumberOID = u.cfg.Board1Pon4.OnuSerialNumberOID // ONU Serial Number OID variable get from config
			onuRxPowerOID = u.cfg.Board1Pon4.OnuRxPowerOID           // ONU RX Power OID variable get from config
			onuStatusOID = u.cfg.Board1Pon4.OnuStatusOID             // ONU Status OID variable get from config
		case 5: // PON 5
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon5.OnuIDNameOID             // ONU ID Name OID variable get from config
			onuTypeOID = u.cfg.Board1Pon5.OnuTypeOID                 // ONU Type OID variable get from config
			onuSerialNumberOID = u.cfg.Board1Pon5.OnuSerialNumberOID // ONU Serial Number OID variable get from config
			onuRxPowerOID = u.cfg.Board1Pon5.OnuRxPowerOID           // ONU RX Power OID variable get from config
			onuStatusOID = u.cfg.Board1Pon5.OnuStatusOID             // ONU Status OID variable get from config
		case 6: // PON 6
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon6.OnuIDNameOID             // ONU ID Name OID variable get from config
			onuTypeOID = u.cfg.Board1Pon6.OnuTypeOID                 // ONU Type OID variable get from config
			onuSerialNumberOID = u.cfg.Board1Pon6.OnuSerialNumberOID // ONU Serial Number OID variable get from config
			onuRxPowerOID = u.cfg.Board1Pon6.OnuRxPowerOID           // ONU RX Power OID variable get from config
			onuStatusOID = u.cfg.Board1Pon6.OnuStatusOID             // ONU Status OID variable get from config
		case 7: // PON 7
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon7.OnuIDNameOID             // ONU ID Name OID variable get from config
			onuTypeOID = u.cfg.Board1Pon7.OnuTypeOID                 // ONU Type OID variable get from config
			onuSerialNumberOID = u.cfg.Board1Pon7.OnuSerialNumberOID // ONU Serial Number OID variable get from config
			onuRxPowerOID = u.cfg.Board1Pon7.OnuRxPowerOID           // ONU RX Power OID variable get from config
			onuStatusOID = u.cfg.Board1Pon7.OnuStatusOID             // ONU Status OID variable get from config
		case 8: // PON 8
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon8.OnuIDNameOID             // ONU ID Name OID variable get from config
			onuTypeOID = u.cfg.Board1Pon8.OnuTypeOID                 // ONU Type OID variable get from config
			onuSerialNumberOID = u.cfg.Board1Pon8.OnuSerialNumberOID // ONU Serial Number OID variable get from config
			onuRxPowerOID = u.cfg.Board1Pon8.OnuRxPowerOID           // ONU RX Power OID variable get from config
			onuStatusOID = u.cfg.Board1Pon8.OnuStatusOID             // Invalid PON ID
		default: // Invalid PON ID
			return nil, errors.New("invalid PON ID") // Return error
		}
	default: // Invalid GTGO ID
		return nil, errors.New("invalid GTGO ID") // Return error
	}

	// Redis Key
	redisKey := "gtgo_" + strconv.Itoa(gtGoID) + "_pon_" + strconv.Itoa(ponID)

	// Try to get data from Redis using GetONUInfoList method with context and Redis key as parameter
	cachedOnuData, err := u.redisRepository.GetONUInfoList(ctx, redisKey)
	if err == nil && cachedOnuData != nil {
		return cachedOnuData, nil // Return cached data if error is nil and cached data is not nil
	}

	var onuInformationList []model.ONUInfoPerGTGO // Create slice to store ONU informationList

	snmpDataMap := make(map[string]gosnmp.SnmpPDU) // Create map to store SNMP data

	/*
		Perform SNMP Walk to get ONU ID and ONU Name
		based on GTGO ID and PON ID using snmpRepository Walk method
		with context and OID as parameter
	*/
	err = u.snmpRepository.Walk(baseOID+onuIDNameOID, func(pdu gosnmp.SnmpPDU) error {
		// Store SNMP data to map with ONU ID as key and PDU as value to be used later
		snmpDataMap[utils.ExtractONUID(pdu.Name)] = pdu
		return nil
	})

	if err != nil {
		return nil, err // Return error if error is not nil
	}

	/*
		Loop through SNMP data map to get ONU information based on ONU ID and ONU Name stored in map before and store
		it to slice of ONU information list to be returned later to caller function as response data
	*/
	for _, pdu := range snmpDataMap {
		onuInfo := model.ONUInfoPerGTGO{
			GTGO: gtGoID,                         // Set GTGO ID to ONU onuInfo struct GTGO field
			PON:  ponID,                          // Set PON ID to ONU onuInfo  struct PON field
			ID:   utils.ExtractIDOnuID(pdu.Name), // Set ONU ID to ONU onuInfo struct ID field
			Name: utils.ExtractName(pdu.Value),   // Set ONU Name to ONU onuInfo struct Name field
		}

		// Get ONU Type based on ONU ID and ONU Type OID and store it to ONU onuInfo struct
		onuType, err := u.getONUType(ctx, onuTypeOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.OnuType = onuType // Set ONU Type to ONU onuInfo struct OnuType field
		}

		// Get ONU Serial Number based on ONU ID and ONU Serial Number OID and store it to ONU onuInfo struct
		onuSerialNumber, err := u.getSerialNumber(ctx, onuSerialNumberOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.SerialNumber = onuSerialNumber // Set ONU Serial Number to ONU onuInfo struct SerialNumber field
		}

		// Get ONU RX Power based on ONU ID and ONU RX Power OID and store it to ONU onuInfo struct
		onuRXPower, err := u.getRxPower(ctx, onuRxPowerOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.RXPower = onuRXPower // Set ONU RX Power to ONU onuInfo struct RXPower field
		}

		// Get ONU Status based on ONU ID and ONU Status OID and store it to ONU onuInfo struct
		onuStatus, err := u.getStatus(ctx, onuStatusOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.Status = onuStatus // Set ONU Status to ONU onuInfo struct Status field
		}

		onuInformationList = append(onuInformationList, onuInfo) // Append ONU onuInfo struct to ONU information list
	}

	// Sort ONU information list based on ONU ID ascending
	sort.Slice(onuInformationList, func(i, j int) bool {
		return onuInformationList[i].ID < onuInformationList[j].ID
	})

	// Save ONU information list to Redis
	err = u.redisRepository.SaveONUInfoList(ctx, redisKey, 300, onuInformationList)
	if err != nil {
		return nil, err // Return error if error is not nil
	}

	return onuInformationList, nil // Return ONU information list and nil error
}

func (u *onuUsecase) GetByGtGoIDPonIDAndOnuID(ctx context.Context, gtGoID, ponID, onuID int) (
	model.ONUCustomerInfo, error,
) {
	// Create context with timeout 30 seconds
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel() // Cancel context when function is done

	var (
		baseOID            string // Base OID variable
		onuIDNameOID       string // ONU ID Name OID variable
		onuTypeOID         string // ONU Type OID variable
		onuSerialNumberOID string // ONU Serial Number OID variable
		onuRxPowerOID      string // ONU RX Power OID variable
		onuTXPowerOID      string // ONU TX Power OID variable
		onuStatusOID       string // ONU Status OID variable
		onuIPAddressOID    string // ONU IP Address OID variable
		onuDescriptionOID  string // ONU Description OID variable
	)

	// Determine base OID and other OID based on GTGO ID and PON ID
	switch gtGoID {
	case 0: // GTGO 0
		switch ponID {
		case 1: // PON 1
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID
			onuIDNameOID = u.cfg.Board1Pon1.OnuIDNameOID             // ONU ID Name OID
			onuTypeOID = u.cfg.Board1Pon1.OnuTypeOID                 // ONU Type OID
			onuSerialNumberOID = u.cfg.Board1Pon1.OnuSerialNumberOID // ONU Serial Number OID
			onuRxPowerOID = u.cfg.Board1Pon1.OnuRxPowerOID           // ONU RX Power OID
			onuTXPowerOID = u.cfg.Board1Pon1.OnuTxPowerOID           // ONU TX Power OID
			onuStatusOID = u.cfg.Board1Pon1.OnuStatusOID             // ONU Status OID
			onuIPAddressOID = u.cfg.Board1Pon1.OnuIPAddressOID       // ONU IP Address OID
			onuDescriptionOID = u.cfg.Board1Pon1.OnuDescriptionOID   // ONU Description OID
		case 2: // PON 2
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID
			onuIDNameOID = u.cfg.Board1Pon2.OnuIDNameOID             // ONU ID Name OID
			onuTypeOID = u.cfg.Board1Pon2.OnuTypeOID                 // ONU Type OID
			onuSerialNumberOID = u.cfg.Board1Pon2.OnuSerialNumberOID // ONU Serial Number OID
			onuRxPowerOID = u.cfg.Board1Pon2.OnuRxPowerOID           // ONU RX Power OID
			onuTXPowerOID = u.cfg.Board1Pon2.OnuTxPowerOID           // ONU TX Power OID
			onuStatusOID = u.cfg.Board1Pon2.OnuStatusOID             // ONU Status OID
			onuIPAddressOID = u.cfg.Board1Pon2.OnuIPAddressOID       // ONU IP Address OID
			onuDescriptionOID = u.cfg.Board1Pon2.OnuDescriptionOID   // ONU Description OID
		case 3: // PON 3
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID
			onuIDNameOID = u.cfg.Board1Pon3.OnuIDNameOID             // ONU ID Name OID
			onuTypeOID = u.cfg.Board1Pon3.OnuTypeOID                 // ONU Type OID
			onuSerialNumberOID = u.cfg.Board1Pon3.OnuSerialNumberOID // ONU Serial Number OID
			onuRxPowerOID = u.cfg.Board1Pon3.OnuRxPowerOID           // ONU RX Power OID
			onuTXPowerOID = u.cfg.Board1Pon3.OnuTxPowerOID           // ONU TX Power OID
			onuStatusOID = u.cfg.Board1Pon3.OnuStatusOID             // ONU Status OID
			onuIPAddressOID = u.cfg.Board1Pon3.OnuIPAddressOID       // ONU IP Address OID
			onuDescriptionOID = u.cfg.Board1Pon3.OnuDescriptionOID   // ONU Description OID
		case 4: // PON 4
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID
			onuIDNameOID = u.cfg.Board1Pon4.OnuIDNameOID             // ONU ID Name OID
			onuTypeOID = u.cfg.Board1Pon4.OnuTypeOID                 // ONU Type OID
			onuSerialNumberOID = u.cfg.Board1Pon4.OnuSerialNumberOID // ONU Serial Number OID
			onuRxPowerOID = u.cfg.Board1Pon4.OnuRxPowerOID           // ONU RX Power OID
			onuTXPowerOID = u.cfg.Board1Pon4.OnuTxPowerOID           // ONU TX Power OID
			onuStatusOID = u.cfg.Board1Pon4.OnuStatusOID             // ONU Status OID
			onuIPAddressOID = u.cfg.Board1Pon4.OnuIPAddressOID       // ONU IP Address OID
			onuDescriptionOID = u.cfg.Board1Pon4.OnuDescriptionOID   // ONU Description OID
		case 5: // PON 5
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID
			onuIDNameOID = u.cfg.Board1Pon5.OnuIDNameOID             // ONU ID Name OID
			onuTypeOID = u.cfg.Board1Pon5.OnuTypeOID                 // ONU Type OID
			onuSerialNumberOID = u.cfg.Board1Pon5.OnuSerialNumberOID // ONU Serial Number OID
			onuRxPowerOID = u.cfg.Board1Pon5.OnuRxPowerOID           // ONU RX Power OID
			onuTXPowerOID = u.cfg.Board1Pon5.OnuTxPowerOID           // ONU TX Power OID
			onuStatusOID = u.cfg.Board1Pon5.OnuStatusOID             // ONU Status OID
			onuIPAddressOID = u.cfg.Board1Pon5.OnuIPAddressOID       // ONU IP Address OID
			onuDescriptionOID = u.cfg.Board1Pon5.OnuDescriptionOID   // ONU Description OID
		case 6: // PON 6
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID
			onuIDNameOID = u.cfg.Board1Pon6.OnuIDNameOID             // ONU ID Name OID
			onuTypeOID = u.cfg.Board1Pon6.OnuTypeOID                 // ONU Type OID
			onuSerialNumberOID = u.cfg.Board1Pon6.OnuSerialNumberOID // ONU Serial Number OID
			onuRxPowerOID = u.cfg.Board1Pon6.OnuRxPowerOID           // ONU RX Power OID
			onuTXPowerOID = u.cfg.Board1Pon6.OnuTxPowerOID           // ONU TX Power OID
			onuStatusOID = u.cfg.Board1Pon6.OnuStatusOID             // ONU Status OID
			onuIPAddressOID = u.cfg.Board1Pon6.OnuIPAddressOID       // ONU IP Address OID
			onuDescriptionOID = u.cfg.Board1Pon6.OnuDescriptionOID   // ONU Description OID
		case 7: // PON 7
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID
			onuIDNameOID = u.cfg.Board1Pon7.OnuIDNameOID             // ONU ID Name OID
			onuTypeOID = u.cfg.Board1Pon7.OnuTypeOID                 // ONU Type OID
			onuSerialNumberOID = u.cfg.Board1Pon7.OnuSerialNumberOID // ONU Serial Number OID
			onuRxPowerOID = u.cfg.Board1Pon7.OnuRxPowerOID           // ONU RX Power OID
			onuTXPowerOID = u.cfg.Board1Pon7.OnuTxPowerOID           // ONU TX Power OID
			onuStatusOID = u.cfg.Board1Pon7.OnuStatusOID             // ONU Status OID
			onuIPAddressOID = u.cfg.Board1Pon7.OnuIPAddressOID       // ONU IP Address OID
			onuDescriptionOID = u.cfg.Board1Pon7.OnuDescriptionOID   // ONU Description OID
		case 8: // PON 8
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID
			onuIDNameOID = u.cfg.Board1Pon8.OnuIDNameOID             // ONU ID Name OID
			onuTypeOID = u.cfg.Board1Pon8.OnuTypeOID                 // ONU Type OID
			onuSerialNumberOID = u.cfg.Board1Pon8.OnuSerialNumberOID // ONU Serial Number OID
			onuRxPowerOID = u.cfg.Board1Pon8.OnuRxPowerOID           // ONU RX Power OID
			onuTXPowerOID = u.cfg.Board1Pon8.OnuTxPowerOID           // ONU TX Power OID
			onuStatusOID = u.cfg.Board1Pon8.OnuStatusOID             // ONU Status OID
			onuIPAddressOID = u.cfg.Board1Pon8.OnuIPAddressOID       // ONU IP Address OID
			onuDescriptionOID = u.cfg.Board1Pon8.OnuDescriptionOID   // ONU Description OID
		default: // Invalid PON ID
			return model.ONUCustomerInfo{}, errors.New("invalid PON ID") // Return error
		}
	case 1: // GTGO 1
		switch ponID {
		case 1: // PON 1
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID
			onuIDNameOID = u.cfg.Board1Pon1.OnuIDNameOID             // ONU ID Name OID
			onuTypeOID = u.cfg.Board1Pon1.OnuTypeOID                 // ONU Type OID
			onuSerialNumberOID = u.cfg.Board1Pon1.OnuSerialNumberOID // ONU Serial Number OID
			onuRxPowerOID = u.cfg.Board1Pon1.OnuRxPowerOID           // ONU RX Power OID
			onuTXPowerOID = u.cfg.Board1Pon1.OnuTxPowerOID           // ONU TX Power OID
			onuStatusOID = u.cfg.Board1Pon1.OnuStatusOID             // ONU Status OID
			onuIPAddressOID = u.cfg.Board1Pon1.OnuIPAddressOID       // ONU IP Address OID
			onuDescriptionOID = u.cfg.Board1Pon1.OnuDescriptionOID   // ONU Description OID
		case 2: // PON 2
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID
			onuIDNameOID = u.cfg.Board1Pon2.OnuIDNameOID             // ONU ID Name OID
			onuTypeOID = u.cfg.Board1Pon2.OnuTypeOID                 // ONU Type OID
			onuSerialNumberOID = u.cfg.Board1Pon2.OnuSerialNumberOID // ONU Serial Number OID
			onuRxPowerOID = u.cfg.Board1Pon2.OnuRxPowerOID           // ONU RX Power OID
			onuTXPowerOID = u.cfg.Board1Pon2.OnuTxPowerOID           // ONU TX Power OID
			onuStatusOID = u.cfg.Board1Pon2.OnuStatusOID             // ONU Status OID
			onuIPAddressOID = u.cfg.Board1Pon2.OnuIPAddressOID       // ONU IP Address OID
			onuDescriptionOID = u.cfg.Board1Pon2.OnuDescriptionOID   // ONU Description OID
		case 3: // PON 3
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID
			onuIDNameOID = u.cfg.Board1Pon3.OnuIDNameOID             // ONU ID Name OID
			onuTypeOID = u.cfg.Board1Pon3.OnuTypeOID                 // ONU Type OID
			onuSerialNumberOID = u.cfg.Board1Pon3.OnuSerialNumberOID // ONU Serial Number OID
			onuRxPowerOID = u.cfg.Board1Pon3.OnuRxPowerOID           // ONU RX Power OID
			onuTXPowerOID = u.cfg.Board1Pon3.OnuTxPowerOID           // ONU TX Power OID
			onuStatusOID = u.cfg.Board1Pon3.OnuStatusOID             // ONU Status OID
			onuIPAddressOID = u.cfg.Board1Pon3.OnuIPAddressOID       // ONU IP Address OID
			onuDescriptionOID = u.cfg.Board1Pon3.OnuDescriptionOID   // ONU Description OID
		case 4: // PON 4
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID
			onuIDNameOID = u.cfg.Board1Pon4.OnuIDNameOID             // ONU ID Name OID
			onuTypeOID = u.cfg.Board1Pon4.OnuTypeOID                 // ONU Type OID
			onuSerialNumberOID = u.cfg.Board1Pon4.OnuSerialNumberOID // ONU Serial Number OID
			onuRxPowerOID = u.cfg.Board1Pon4.OnuRxPowerOID           // ONU RX Power OID
			onuTXPowerOID = u.cfg.Board1Pon4.OnuTxPowerOID           // ONU TX Power OID
			onuStatusOID = u.cfg.Board1Pon4.OnuStatusOID             // ONU Status OID
			onuIPAddressOID = u.cfg.Board1Pon4.OnuIPAddressOID       // ONU IP Address OID
			onuDescriptionOID = u.cfg.Board1Pon4.OnuDescriptionOID   // ONU Description OID
		case 5: // PON 5
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID
			onuIDNameOID = u.cfg.Board1Pon5.OnuIDNameOID             // ONU ID Name OID
			onuTypeOID = u.cfg.Board1Pon5.OnuTypeOID                 // ONU Type OID
			onuSerialNumberOID = u.cfg.Board1Pon5.OnuSerialNumberOID // ONU Serial Number OID
			onuRxPowerOID = u.cfg.Board1Pon5.OnuRxPowerOID           // ONU RX Power OID
			onuTXPowerOID = u.cfg.Board1Pon5.OnuTxPowerOID           // ONU TX Power OID
			onuStatusOID = u.cfg.Board1Pon5.OnuStatusOID             // ONU Status OID
			onuIPAddressOID = u.cfg.Board1Pon5.OnuIPAddressOID       // ONU IP Address OID
			onuDescriptionOID = u.cfg.Board1Pon5.OnuDescriptionOID   // ONU Description OID
		case 6: // PON 6
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID
			onuIDNameOID = u.cfg.Board1Pon6.OnuIDNameOID             // ONU ID Name OID
			onuTypeOID = u.cfg.Board1Pon6.OnuTypeOID                 // ONU Type OID
			onuSerialNumberOID = u.cfg.Board1Pon6.OnuSerialNumberOID // ONU Serial Number OID
			onuRxPowerOID = u.cfg.Board1Pon6.OnuRxPowerOID           // ONU RX Power OID
			onuTXPowerOID = u.cfg.Board1Pon6.OnuTxPowerOID           // ONU TX Power OID
			onuStatusOID = u.cfg.Board1Pon6.OnuStatusOID             // ONU Status OID
			onuIPAddressOID = u.cfg.Board1Pon6.OnuIPAddressOID       // ONU IP Address OID
			onuDescriptionOID = u.cfg.Board1Pon6.OnuDescriptionOID   // ONU Description OID
		case 7: // PON 7
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID
			onuIDNameOID = u.cfg.Board1Pon7.OnuIDNameOID             // ONU ID Name OID
			onuTypeOID = u.cfg.Board1Pon7.OnuTypeOID                 // ONU Type OID
			onuSerialNumberOID = u.cfg.Board1Pon7.OnuSerialNumberOID // ONU Serial Number OID
			onuRxPowerOID = u.cfg.Board1Pon7.OnuRxPowerOID           // ONU RX Power OID
			onuTXPowerOID = u.cfg.Board1Pon7.OnuTxPowerOID           // ONU TX Power OID
			onuStatusOID = u.cfg.Board1Pon7.OnuStatusOID             // ONU Status OID
			onuIPAddressOID = u.cfg.Board1Pon7.OnuIPAddressOID       // ONU IP Address OID
			onuDescriptionOID = u.cfg.Board1Pon7.OnuDescriptionOID   // ONU Description OID
		case 8: // PON 8
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID
			onuIDNameOID = u.cfg.Board1Pon8.OnuIDNameOID             // ONU ID Name OID
			onuTypeOID = u.cfg.Board1Pon8.OnuTypeOID                 // ONU Type OID
			onuSerialNumberOID = u.cfg.Board1Pon8.OnuSerialNumberOID // ONU Serial Number OID
			onuRxPowerOID = u.cfg.Board1Pon8.OnuRxPowerOID           // ONU RX Power OID
			onuTXPowerOID = u.cfg.Board1Pon8.OnuTxPowerOID           // ONU TX Power OID
			onuStatusOID = u.cfg.Board1Pon8.OnuStatusOID             // ONU Status OID
			onuIPAddressOID = u.cfg.Board1Pon8.OnuIPAddressOID       // ONU IP Address OID
			onuDescriptionOID = u.cfg.Board1Pon8.OnuDescriptionOID   // ONU Description OID
		default: // Invalid PON ID
			return model.ONUCustomerInfo{}, errors.New("invalid PON ID") // Return error
		}
	default: // Invalid GTGO ID
		return model.ONUCustomerInfo{}, errors.New("invalid GTGO ID") // Return error
	}

	// Create a slice of ONUCustomerInfo
	var onuInformationList model.ONUCustomerInfo

	// Create a map to store SNMP Walk results
	snmpDataMap := make(map[string]gosnmp.SnmpPDU)

	// Perform SNMP Walk to get ONU ID and Name using snmpRepository Walk method with timeout context parameter
	err := u.snmpRepository.Walk(baseOID+onuIDNameOID+"."+strconv.Itoa(onuID),
		func(pdu gosnmp.SnmpPDU) error {
			// Save SNMP Walk result in map with ID as key and Name as value (extracted from SNMP PDU)
			snmpDataMap[utils.ExtractONUID(pdu.Name)] = pdu // Extract ONU ID from SNMP PDU Name and use it as key in map
			return nil
		})

	if err != nil {
		return model.ONUCustomerInfo{}, errors.New("failed to walk OID") // Return error
	}

	/*
		Loop through SNMP data map to get ONU information based on ONU ID and ONU Name stored in map before and store
		it to slice of ONU information list to be returned later to caller function as response data
	*/
	for _, pdu := range snmpDataMap {
		onuInfo := model.ONUCustomerInfo{
			GTGO: gtGoID,                         // Set GTGO ID to ONU onuInfo struct GTGO field
			PON:  ponID,                          // Set PON ID to ONU onuInfo  struct PON field
			ID:   utils.ExtractIDOnuID(pdu.Name), // Set ONU ID (extracted from SNMP PDU) to onuInfo variable (ONU ID)
			Name: utils.ExtractName(pdu.Value),   // Set ONU Name (extracted from SNMP PDU) to onuInfo variable (ONU Name)
		}

		// Get Data ONU Type from SNMP Walk using getONUType method
		onuType, err := u.getONUType(ctx, onuTypeOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.OnuType = onuType // Set ONU Type from SNMP Walk result if no error to onuInfo variable (ONU Type)
		}

		// Get Data ONU Serial Number from SNMP Walk using getSerialNumber method
		onuSerialNumber, err := u.getSerialNumber(ctx, onuSerialNumberOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.SerialNumber = onuSerialNumber // Set ONU Serial Number from SNMP Walk result to onuInfo variable (ONU Serial Number)
		}

		// Get Data ONU RX Power from SNMP Walk using getRxPower method
		onuRXPower, err := u.getRxPower(ctx, onuRxPowerOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.RXPower = onuRXPower // Set ONU RX Power from SNMP Walk result to onuInfo variable (ONU RX Power)
		}

		// Get Data ONU TX Power from SNMP Walk using getTxPower method
		onuTXPower, err := u.getTxPower(ctx, onuTXPowerOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.TXPower = onuTXPower // Set ONU TX Power from SNMP Walk result to onuInfo variable (ONU TX Power)
		}

		// Get Data ONU Status from SNMP Walk using getStatus method
		onuStatus, err := u.getStatus(ctx, onuStatusOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.Status = onuStatus // Set ONU Status from SNMP Walk result to onuInfo variable (ONU Status)
		}

		// Get Data ONU IP Address from SNMP Walk using getIPAddress method
		onuIPAddress, err := u.getIPAddress(ctx, onuIPAddressOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.IPAddress = onuIPAddress // Set ONU IP Address from SNMP Walk result to onuInfo variable (ONU IP Address)
		}

		// Get Data ONU Description from SNMP Walk using getDescription method
		onuDescription, err := u.getDescription(ctx, onuDescriptionOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.Description = onuDescription // Set ONU Description from SNMP Walk result to onuInfo variable (ONU Description)
		}

		onuInformationList = onuInfo // Append onuInfo variable to onuInformationList slice
	}

	return onuInformationList, nil
}

func (u *onuUsecase) getONUType(ctx context.Context, onuTypeOID, onuID string) (string, error) {

	var onuType string // Variable to store ONU Type

	ctx, cancel := context.WithTimeout(ctx, time.Second*30) // Create context with timeout 30 seconds
	defer cancel()                                          // Cancel context when function is done

	baseOID := u.cfg.OltCfg.BaseOID2 // Base OID variable get from config

	// Perform SNMP Walk to get ONU Type using snmpRepository Walk method with timeout context parameter
	err := u.snmpRepository.Walk(baseOID+onuTypeOID+"."+onuID, func(pdu gosnmp.SnmpPDU) error {
		onuType = utils.ExtractName(pdu.Value) // Extract ONU Type from SNMP PDU Value
		return nil
	})

	if err != nil {
		return "", errors.New("failed to perform SNMP Walk") // Return error
	}

	return onuType, nil // Return ONU Type
}

func (u *onuUsecase) getSerialNumber(ctx context.Context, onuSerialNumberOID, onuID string) (string, error) {

	var onuSerialNumber string // Variable to store ONU Serial Number

	ctx, cancel := context.WithTimeout(ctx, time.Second*30) // Create context with timeout 30 seconds
	defer cancel()                                          // Cancel context when function is done

	baseOID := u.cfg.OltCfg.BaseOID1 // Base OID variable get from config

	// Perform SNMP Walk to get ONU Serial Number using snmpRepository Walk method with timeout context parameter
	err := u.snmpRepository.Walk(baseOID+onuSerialNumberOID+"."+onuID, func(pdu gosnmp.SnmpPDU) error {
		onuSerialNumber = utils.ExtractSerialNumber(pdu.Value) // Extract ONU Serial Number from SNMP PDU Value
		return nil
	})

	if err != nil {
		return "", errors.New("failed to perform SNMP Walk") // Return error
	}

	return onuSerialNumber, nil // Return ONU Serial Number
}

func (u *onuUsecase) getRxPower(ctx context.Context, onuRxPowerOID, onuID string) (string, error) {

	var onuRxPower string // Variable to store ONU RX Power

	ctx, cancel := context.WithTimeout(ctx, time.Second*30) // Create context with timeout 30 seconds
	defer cancel()                                          // Cancel context when function is done

	baseOID := u.cfg.OltCfg.BaseOID1 // Base OID variable get from config

	// Perform SNMP Walk to get ONU RX Power using snmpRepository Walk method with timeout context parameter
	err := u.snmpRepository.Walk(baseOID+onuRxPowerOID+"."+onuID+"."+"1", func(pdu gosnmp.SnmpPDU) error {
		onuRxPower, _ = utils.ConvertAndMultiply(pdu.Value) // Extract ONU RX Power from SNMP PDU Value
		return nil
	})

	if err != nil {
		return "", errors.New("failed to perform SNMP Walk") // Return error
	}

	return onuRxPower, nil // Return ONU RX Power
}

func (u *onuUsecase) getTxPower(ctx context.Context, onuTxPowerOID, onuID string) (string, error) {

	var onuTxPower string // Variable to store ONU TX Power

	ctx, cancel := context.WithTimeout(ctx, time.Second*30) // Create context with timeout 30 seconds
	defer cancel()                                          // Cancel context when function is done

	baseOID := u.cfg.OltCfg.BaseOID2 // Base OID variable get from config

	// Perform SNMP Walk to get ONU TX Power using snmpRepository Walk method with timeout context parameter
	err := u.snmpRepository.Walk(baseOID+onuTxPowerOID+"."+onuID+"."+"1", func(pdu gosnmp.SnmpPDU) error {
		onuTxPower, _ = utils.ConvertAndMultiply(pdu.Value) // Extract ONU TX Power from SNMP PDU Value
		return nil
	})

	if err != nil {
		return "", errors.New("failed to perform SNMP Walk") // Return error
	}

	return onuTxPower, nil // Return ONU TX Power
}

func (u *onuUsecase) getStatus(ctx context.Context, onuStatusOID, onuID string) (string, error) {

	var onuStatus string // Variable to store ONU Status

	ctx, cancel := context.WithTimeout(ctx, time.Second*30) // Create context with timeout 30 seconds
	defer cancel()                                          // Cancel context when function is done

	baseOID := u.cfg.OltCfg.BaseOID1 // Base OID variable get from config

	// Perform SNMP Walk to get ONU Status using snmpRepository Walk method with timeout context parameter
	err := u.snmpRepository.Walk(baseOID+onuStatusOID+"."+onuID, func(pdu gosnmp.SnmpPDU) error {
		onuStatus = utils.ExtractAndGetStatus(pdu.Value) // Extract ONU Status from SNMP PDU Value
		return nil
	})

	if err != nil {
		return "", errors.New("failed to perform SNMP Walk") // Return error
	}

	return onuStatus, nil // Return ONU Status
}

func (u *onuUsecase) getIPAddress(ctx context.Context, onuIPAddressOID, onuID string) (string, error) {

	var onuIPAddress string // Variable to store ONU IP Address

	ctx, cancel := context.WithTimeout(ctx, time.Second*30) // Create context with timeout 30 seconds
	defer cancel()                                          // Cancel context when function is done

	baseOID := u.cfg.OltCfg.BaseOID2 // Base OID variable get from config

	// Perform SNMP Walk to get ONU IP Address using snmpRepository Walk method with timeout context parameter
	err := u.snmpRepository.Walk(baseOID+onuIPAddressOID+"."+onuID+"."+strconv.Itoa(1), func(pdu gosnmp.SnmpPDU) error {
		onuIPAddress = utils.ExtractName(pdu.Value) // Extract ONU IP Address from SNMP PDU Value
		return nil
	})

	if err != nil {
		return "", errors.New("failed to perform SNMP Walk") // Return error
	}

	return onuIPAddress, nil // Return ONU IP Address
}

func (u *onuUsecase) getDescription(ctx context.Context, onuDescriptionOID, onuID string) (string, error) {

	var onuDescription string // Variable to store ONU Description

	ctx, cancel := context.WithTimeout(ctx, time.Second*30) // Create context with timeout 30 seconds
	defer cancel()                                          // Cancel context when function is done

	baseOID := u.cfg.OltCfg.BaseOID1 // Base OID variable get from config

	// Perform SNMP Walk to get ONU Description using snmpRepository Walk method with timeout context parameter
	err := u.snmpRepository.Walk(baseOID+onuDescriptionOID+"."+onuID, func(pdu gosnmp.SnmpPDU) error {
		onuDescription = utils.ExtractName(pdu.Value) // Extract ONU Description from SNMP PDU Value
		return nil
	})

	if err != nil {
		return "", errors.New("failed to perform SNMP Walk") // Return error
	}

	return onuDescription, nil // Return ONU Description
}

func (u *onuUsecase) GetEmptyOnuID(ctx context.Context, gtGoID, ponID int) ([]model.OnuID, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*30) // Create context with timeout 30 seconds
	defer cancel()

	var (
		baseOID      string // Base OID variable
		onuIDNameOID string // ONU ID Name OID variable
	)

	// Determine base OID and other OID based on GTGO ID and PON ID
	switch gtGoID {
	case 0: // GTGO 0
		switch ponID {
		case 1: // PON 1
			baseOID = u.cfg.OltCfg.BaseOID1              // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon1.OnuIDNameOID // ONU ID Name OID variable get from config
		case 2: // PON 2
			baseOID = u.cfg.OltCfg.BaseOID1              // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon2.OnuIDNameOID // ONU ID Name OID variable get from config
		case 3: // PON 3
			baseOID = u.cfg.OltCfg.BaseOID1              // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon3.OnuIDNameOID // ONU ID Name OID variable get from config
		case 4: // PON 4
			baseOID = u.cfg.OltCfg.BaseOID1              // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon4.OnuIDNameOID // ONU ID Name OID variable get from config
		case 5: // PON 5
			baseOID = u.cfg.OltCfg.BaseOID1              // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon5.OnuIDNameOID // ONU ID Name OID variable get from config
		case 6: // PON 6
			baseOID = u.cfg.OltCfg.BaseOID1              // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon6.OnuIDNameOID // ONU ID Name OID variable get from config
		case 7: // PON 7
			baseOID = u.cfg.OltCfg.BaseOID1              // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon7.OnuIDNameOID // ONU ID Name OID variable get from config
		case 8: // PON 8
			baseOID = u.cfg.OltCfg.BaseOID1              // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon8.OnuIDNameOID // ONU ID Name OID variable get from config
		default: // Invalid PON ID
			return nil, errors.New("invalid PON ID") // Return error
		}
	case 1: // GTGO 1
		switch ponID {
		case 1: // PON 1
			baseOID = u.cfg.OltCfg.BaseOID1              // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon1.OnuIDNameOID // ONU ID Name OID variable get from config
		case 2: // PON 2
			baseOID = u.cfg.OltCfg.BaseOID1              // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon2.OnuIDNameOID // ONU ID Name OID variable get from config
		case 3: // PON 3
			baseOID = u.cfg.OltCfg.BaseOID1              // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon3.OnuIDNameOID // ONU ID Name OID variable get from config
		case 4: // PON 4
			baseOID = u.cfg.OltCfg.BaseOID1              // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon4.OnuIDNameOID // ONU ID Name OID variable get from config
		case 5: // PON 5
			baseOID = u.cfg.OltCfg.BaseOID1              // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon5.OnuIDNameOID // ONU ID Name OID variable get from config
		case 6: // PON 6
			baseOID = u.cfg.OltCfg.BaseOID1              // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon6.OnuIDNameOID // ONU ID Name OID variable get from config
		case 7: // PON 7
			baseOID = u.cfg.OltCfg.BaseOID1              // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon7.OnuIDNameOID // ONU ID Name OID variable get from config
		case 8: // PON 8
			baseOID = u.cfg.OltCfg.BaseOID1              // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon8.OnuIDNameOID // ONU ID Name OID variable get from config
		default: // Invalid PON ID
			return nil, errors.New("invalid PON ID") // Return error
		}
	default: // Invalid GTGO ID
		return nil, errors.New("invalid GTGO ID") // Return error
	}

	//Redis Key
	redisKey := "gtgo_" + strconv.Itoa(gtGoID) + "_pon_" + strconv.Itoa(ponID) + "_empty_onu_id"

	//Try to get data from Redis using GetOnuIDCtx method with context and Redis key as parameter
	cachedOnuData, err := u.redisRepository.GetOnuIDCtx(ctx, redisKey)
	if err == nil && cachedOnuData != nil {
		// If data exist in Redis, then return data from Redis
		return cachedOnuData, nil
	}

	// Perform SNMP Walk to get ONU ID and ONU Name
	snmpOID := baseOID + onuIDNameOID        // SNMP OID variable
	emptyOnuIDList := make([]model.OnuID, 0) // Create a slice of ONU ID

	// Perform SNMP Walk to get ONU ID and Name using snmpRepository Walk method with timeout context parameter
	err = u.snmpRepository.Walk(snmpOID, func(pdu gosnmp.SnmpPDU) error {
		idOnuID := utils.ExtractIDOnuID(pdu.Name) // Extract ONU ID from SNMP PDU Name

		// Append ONU information to the emptyOnuIDList
		emptyOnuIDList = append(emptyOnuIDList, model.OnuID{
			GTGO: gtGoID,  // Set GTGO ID to ONU onuInfo struct GTGO field
			PON:  ponID,   // Set PON ID to ONU onuInfo  struct PON field
			ID:   idOnuID, // Set ONU ID (extracted from SNMP PDU) to onuInfo variable (ONU ID)
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Create a map to store numbers to be deleted
	numbersToRemove := make(map[int]bool)

	// Loop through emptyOnuIDList to get the numbers to be deleted
	for _, onuInfo := range emptyOnuIDList {
		numbersToRemove[onuInfo.ID] = true
	}

	// Create a new slice to hold the gtgo_id, pon_id and onu_id data without the numbers to be deleted
	emptyOnuIDList = emptyOnuIDList[:0]

	// Loop through 128 numbers to get the numbers to be deleted
	for i := 1; i <= 128; i++ {
		if _, ok := numbersToRemove[i]; !ok {
			emptyOnuIDList = append(emptyOnuIDList, model.OnuID{
				GTGO: gtGoID, // Set GTGO ID to ONU onuInfo struct GTGO field
				PON:  ponID,  // Set PON ID to ONU onuInfo  struct PON field
				ID:   i,      // Number 1-128 that is not in the numbers to be deleted
			})
		}
	}

	// Sort by ID ascending
	sort.Slice(emptyOnuIDList, func(i, j int) bool {
		return emptyOnuIDList[i].ID < emptyOnuIDList[j].ID
	})

	// Set data to Redis using SetOnuIDCtx method with context, Redis key and data as parameter
	err = u.redisRepository.SetOnuIDCtx(ctx, redisKey, 300, emptyOnuIDList)
	if err != nil {
		return nil, err
	}

	return emptyOnuIDList, nil
}

func (u *onuUsecase) GetByGtGoIDAndPonIDWithPagination(
	ctx context.Context, gtGoID, ponID, pageIndex,
	pageSize int,
) (
	[]model.ONUInfoPerGTGO,
	int,
) {

	ctx, cancel := context.WithTimeout(ctx, time.Second*30) // Create context with timeout 30 seconds
	defer cancel()                                          // Cancel context when function is done

	var (
		baseOID            string // Base OID variable
		onuIDNameOID       string // ONU ID Name OID variable
		onuTypeOID         string // ONU Type OID variable
		onuSerialNumberOID string // ONU Serial Number OID variable
		onuRxPowerOID      string // ONU RX Power OID variable
		onuStatusOID       string // ONU Status OID variable
	)

	// Determine base OID and other OID based on GTGO ID and PON ID
	switch gtGoID {
	case 0: // GTGO 0
		switch ponID {
		case 1: // PON 1
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon1.OnuIDNameOID             // ONU ID Name OID variable get from config
			onuTypeOID = u.cfg.Board1Pon1.OnuTypeOID                 // ONU Type OID variable get from config
			onuSerialNumberOID = u.cfg.Board1Pon1.OnuSerialNumberOID // ONU Serial Number OID variable get from config
			onuRxPowerOID = u.cfg.Board1Pon1.OnuRxPowerOID           // ONU RX Power OID variable get from config
			onuStatusOID = u.cfg.Board1Pon1.OnuStatusOID             // ONU Status OID variable get from config
		case 2: // PON 2
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon2.OnuIDNameOID             // ONU ID Name OID variable get from config
			onuTypeOID = u.cfg.Board1Pon2.OnuTypeOID                 // ONU Type OID variable get from config
			onuSerialNumberOID = u.cfg.Board1Pon2.OnuSerialNumberOID // ONU Serial Number OID variable get from config
			onuRxPowerOID = u.cfg.Board1Pon2.OnuRxPowerOID           // ONU RX Power OID variable get from config
			onuStatusOID = u.cfg.Board1Pon2.OnuStatusOID             // ONU Status OID variable get from config
		case 3: // PON 3
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon3.OnuIDNameOID             // ONU ID Name OID variable get from config
			onuTypeOID = u.cfg.Board1Pon3.OnuTypeOID                 // ONU Type OID variable get from config
			onuSerialNumberOID = u.cfg.Board1Pon3.OnuSerialNumberOID // ONU Serial Number OID variable get from config
			onuRxPowerOID = u.cfg.Board1Pon3.OnuRxPowerOID           // ONU RX Power OID variable get from config
			onuStatusOID = u.cfg.Board1Pon3.OnuStatusOID             // ONU Status OID variable get from config
		case 4: // PON 4
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon4.OnuIDNameOID             // ONU ID Name OID variable get from config
			onuTypeOID = u.cfg.Board1Pon4.OnuTypeOID                 // ONU Type OID variable get from config
			onuSerialNumberOID = u.cfg.Board1Pon4.OnuSerialNumberOID // ONU Serial Number OID variable get from config
			onuRxPowerOID = u.cfg.Board1Pon4.OnuRxPowerOID           // ONU RX Power OID variable get from config
			onuStatusOID = u.cfg.Board1Pon4.OnuStatusOID             // ONU Status OID variable get from config
		case 5: // PON 5
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon5.OnuIDNameOID             // ONU ID Name OID variable get from config
			onuTypeOID = u.cfg.Board1Pon5.OnuTypeOID                 // ONU Type OID variable get from config
			onuSerialNumberOID = u.cfg.Board1Pon5.OnuSerialNumberOID // ONU Serial Number OID variable get from config
			onuRxPowerOID = u.cfg.Board1Pon5.OnuRxPowerOID           // ONU RX Power OID variable get from config
			onuStatusOID = u.cfg.Board1Pon5.OnuStatusOID             // ONU Status OID variable get from config
		case 6: // PON 6
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon6.OnuIDNameOID             // ONU ID Name OID variable get from config
			onuTypeOID = u.cfg.Board1Pon6.OnuTypeOID                 // ONU Type OID variable get from config
			onuSerialNumberOID = u.cfg.Board1Pon6.OnuSerialNumberOID // ONU Serial Number OID variable get from config
			onuRxPowerOID = u.cfg.Board1Pon6.OnuRxPowerOID           // ONU RX Power OID variable get from config
			onuStatusOID = u.cfg.Board1Pon6.OnuStatusOID             // ONU Status OID variable get from config
		case 7: // PON 7
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon7.OnuIDNameOID             // ONU ID Name OID variable get from config
			onuTypeOID = u.cfg.Board1Pon7.OnuTypeOID                 // ONU Type OID variable get from config
			onuSerialNumberOID = u.cfg.Board1Pon7.OnuSerialNumberOID // ONU Serial Number OID variable get from config
			onuRxPowerOID = u.cfg.Board1Pon7.OnuRxPowerOID           // ONU RX Power OID variable get from config
			onuStatusOID = u.cfg.Board1Pon7.OnuStatusOID             // ONU Status OID variable get from config
		case 8: // PON 8
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon8.OnuIDNameOID             // ONU ID Name OID variable get from config
			onuTypeOID = u.cfg.Board1Pon8.OnuTypeOID                 // ONU Type OID variable get from config
			onuSerialNumberOID = u.cfg.Board1Pon8.OnuSerialNumberOID // ONU Serial Number OID variable get from config
			onuRxPowerOID = u.cfg.Board1Pon8.OnuRxPowerOID           // ONU RX Power OID variable get from config
			onuStatusOID = u.cfg.Board1Pon8.OnuStatusOID             // ONU Status OID variable get from config
		default: // Invalid PON ID
			return nil, 0 // Return error
		}
	case 1: // GTGO 1
		switch ponID {
		case 1: // PON 1
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon1.OnuIDNameOID             // ONU ID Name OID variable get from config
			onuTypeOID = u.cfg.Board1Pon1.OnuTypeOID                 // ONU Type OID variable get from config
			onuSerialNumberOID = u.cfg.Board1Pon1.OnuSerialNumberOID // ONU Serial Number OID variable get from config
			onuRxPowerOID = u.cfg.Board1Pon1.OnuRxPowerOID           // ONU RX Power OID variable get from config
			onuStatusOID = u.cfg.Board1Pon1.OnuStatusOID             // ONU Status OID variable get from config
		case 2: // PON 2
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon2.OnuIDNameOID             // ONU ID Name OID variable get from config
			onuTypeOID = u.cfg.Board1Pon2.OnuTypeOID                 // ONU Type OID variable get from config
			onuSerialNumberOID = u.cfg.Board1Pon2.OnuSerialNumberOID // ONU Serial Number OID variable get from config
			onuRxPowerOID = u.cfg.Board1Pon2.OnuRxPowerOID           // ONU RX Power OID variable get from config
			onuStatusOID = u.cfg.Board1Pon2.OnuStatusOID             // ONU Status OID variable get from config
		case 3: // PON 3
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon3.OnuIDNameOID             // ONU ID Name OID variable get from config
			onuTypeOID = u.cfg.Board1Pon3.OnuTypeOID                 // ONU Type OID variable get from config
			onuSerialNumberOID = u.cfg.Board1Pon3.OnuSerialNumberOID // ONU Serial Number OID variable get from config
			onuRxPowerOID = u.cfg.Board1Pon3.OnuRxPowerOID           // ONU RX Power OID variable get from config
			onuStatusOID = u.cfg.Board1Pon3.OnuStatusOID             // ONU Status OID variable get from config
		case 4: // PON 4
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon4.OnuIDNameOID             // ONU ID Name OID variable get from config
			onuTypeOID = u.cfg.Board1Pon4.OnuTypeOID                 // ONU Type OID variable get from config
			onuSerialNumberOID = u.cfg.Board1Pon4.OnuSerialNumberOID // ONU Serial Number OID variable get from config
			onuRxPowerOID = u.cfg.Board1Pon4.OnuRxPowerOID           // ONU RX Power OID variable get from config
			onuStatusOID = u.cfg.Board1Pon4.OnuStatusOID             // ONU Status OID variable get from config
		case 5: // PON 5
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon5.OnuIDNameOID             // ONU ID Name OID variable get from config
			onuTypeOID = u.cfg.Board1Pon5.OnuTypeOID                 // ONU Type OID variable get from config
			onuSerialNumberOID = u.cfg.Board1Pon5.OnuSerialNumberOID // ONU Serial Number OID variable get from config
			onuRxPowerOID = u.cfg.Board1Pon5.OnuRxPowerOID           // ONU RX Power OID variable get from config
			onuStatusOID = u.cfg.Board1Pon5.OnuStatusOID             // ONU Status OID variable get from config
		case 6: // PON 6
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon6.OnuIDNameOID             // ONU ID Name OID variable get from config
			onuTypeOID = u.cfg.Board1Pon6.OnuTypeOID                 // ONU Type OID variable get from config
			onuSerialNumberOID = u.cfg.Board1Pon6.OnuSerialNumberOID // ONU Serial Number OID variable get from config
			onuRxPowerOID = u.cfg.Board1Pon6.OnuRxPowerOID           // ONU RX Power OID variable get from config
			onuStatusOID = u.cfg.Board1Pon6.OnuStatusOID             // ONU Status OID variable get from config
		case 7: // PON 7
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon7.OnuIDNameOID             // ONU ID Name OID variable get from config
			onuTypeOID = u.cfg.Board1Pon7.OnuTypeOID                 // ONU Type OID variable get from config
			onuSerialNumberOID = u.cfg.Board1Pon7.OnuSerialNumberOID // ONU Serial Number OID variable get from config
			onuRxPowerOID = u.cfg.Board1Pon7.OnuRxPowerOID           // ONU RX Power OID variable get from config
			onuStatusOID = u.cfg.Board1Pon7.OnuStatusOID             // ONU Status OID variable get from config
		case 8: // PON 8
			baseOID = u.cfg.OltCfg.BaseOID1                          // Base OID variable get from config
			onuIDNameOID = u.cfg.Board1Pon8.OnuIDNameOID             // ONU ID Name OID variable get from config
			onuTypeOID = u.cfg.Board1Pon8.OnuTypeOID                 // ONU Type OID variable get from config
			onuSerialNumberOID = u.cfg.Board1Pon8.OnuSerialNumberOID // ONU Serial Number OID variable get from config
			onuRxPowerOID = u.cfg.Board1Pon8.OnuRxPowerOID           // ONU RX Power OID variable get from config
			onuStatusOID = u.cfg.Board1Pon8.OnuStatusOID             // Invalid PON ID
		default: // Invalid PON ID
			return nil, 0 // Return error
		}
	default: // Invalid GTGO ID
		return nil, 0 // Return error
	}

	fmt.Println("gtgoID: ", gtGoID)
	fmt.Println("ponID:", ponID)
	fmt.Println("pageIndex:", pageIndex)
	fmt.Println("pageSize:", pageSize)

	var onuIDList []model.OnuID // Create slice to store ONU ID list

	// Redis Key
	redisKey := "gtgo_" + strconv.Itoa(gtGoID) + "_pon_" + strconv.Itoa(ponID)
	fmt.Println("redisKey:", redisKey)

	// Try to get data from Redis using GetOnuIDCtx method with context and Redis key as parameter
	cachedOnuData, err := u.redisRepository.GetOnuIDCtx(ctx, redisKey)
	if err == nil && cachedOnuData != nil {
		// If data exist in Redis, then return data from Redis
		onuIDList = cachedOnuData
	}

	// If data not exist in Redis, then get data from SNMP
	if len(onuIDList) == 0 {
		// Perform SNMP Walk to get ONU ID and ONU Name based on GTGO ID and PON ID using snmpRepository Walk method
		// with context and OID as parameter
		err = u.snmpRepository.Walk(baseOID+onuIDNameOID, func(pdu gosnmp.SnmpPDU) error {
			onuID := model.OnuID{
				ID: utils.ExtractIDOnuID(pdu.Name), // Set ONU ID to ONU onuInfo struct ID field
			}
			onuIDList = append(onuIDList, onuID) // Append ONU onuInfo struct to ONU information list
			return nil
		})

		if err != nil {
			return nil, 0 // Return error if error is not nil
		}

		// Sort ONU ID list based on ONU ID ascending
		sort.Slice(onuIDList, func(i, j int) bool {
			return onuIDList[i].ID < onuIDList[j].ID
		})

		// Set data to Redis
		err = u.redisRepository.SetOnuIDCtx(ctx, redisKey, 300, onuIDList)

		if err != nil {
			return nil, 0
		}
	} else {
		fmt.Println("data from redis")
	}

	// print onuIDList
	fmt.Println(onuIDList)

	var onuInformationList []model.ONUInfoPerGTGO // Create slice to store ONU informationList
	var count int

	snmpDataMap := make(map[string]gosnmp.SnmpPDU) // Create map to store SNMP data

	/*
		Perform SNMP Walk to get ONU ID and ONU Name
		based on GTGO ID and PON ID using snmpRepository Walk method
		with context and OID as parameter
	*/
	err = u.snmpRepository.Walk(baseOID+onuIDNameOID, func(pdu gosnmp.SnmpPDU) error {
		// Store SNMP data to map with ONU ID as key and PDU as value to be used later
		snmpDataMap[utils.ExtractONUID(pdu.Name)] = pdu
		return nil
	})

	if err != nil {
		return nil, 0 // Return error if error is not nil
	}

	fmt.Println(gtGoID)
	fmt.Println(ponID)
	fmt.Println(pageIndex)
	fmt.Println(pageSize)

	// Menghitung indeks item pertama yang akan diambil
	startIndex := pageIndex * pageSize

	// Menghitung indeks item terakhir yang akan diambil
	endIndex := startIndex + pageSize

	var currentIndex int

	//fmt.Println(snmpDataMap)

	for _, pdu := range snmpDataMap {
		if currentIndex >= startIndex {
			onuInfo := model.ONUInfoPerGTGO{
				GTGO: gtGoID,
				PON:  ponID,
				ID:   utils.ExtractIDOnuID(pdu.Name),
				Name: utils.ExtractName(pdu.Value),
			}

			// Get ONU Type based on ONU ID and ONU Type OID and store it to ONU onuInfo struct
			onuType, err := u.getONUType(ctx, onuTypeOID, strconv.Itoa(onuInfo.ID))
			if err == nil {
				onuInfo.OnuType = onuType // Set ONU Type to ONU onuInfo struct OnuType field
			}

			// Get ONU Serial Number based on ONU ID and ONU Serial Number OID and store it to ONU onuInfo struct
			onuSerialNumber, err := u.getSerialNumber(ctx, onuSerialNumberOID, strconv.Itoa(onuInfo.ID))
			if err == nil {
				onuInfo.SerialNumber = onuSerialNumber // Set ONU Serial Number to ONU onuInfo struct SerialNumber field
			}

			// Get ONU RX Power based on ONU ID and ONU RX Power OID and store it to ONU onuInfo struct
			onuRXPower, err := u.getRxPower(ctx, onuRxPowerOID, strconv.Itoa(onuInfo.ID))
			if err == nil {
				onuInfo.RXPower = onuRXPower // Set ONU RX Power to ONU onuInfo struct RXPower field
			}

			// Get ONU Status based on ONU ID and ONU Status OID and store it to ONU onuInfo struct
			onuStatus, err := u.getStatus(ctx, onuStatusOID, strconv.Itoa(onuInfo.ID))
			if err == nil {
				onuInfo.Status = onuStatus // Set ONU Status to ONU onuInfo struct Status field
			}

			onuInformationList = append(onuInformationList, onuInfo)
		}

		currentIndex++

		if currentIndex >= endIndex {
			break
		}
	}

	// Sort ONU information list based on ONU ID ascending
	sort.Slice(onuInformationList, func(i, j int) bool {
		return onuInformationList[i].ID < onuInformationList[j].ID
	})

	//fmt.Println(onuInformationList)

	if startIndex >= count {
		return nil, 0
	}
	if endIndex > count {
		endIndex = count
	}

	paginatedData := onuInformationList[startIndex:endIndex]

	fmt.Println(paginatedData)

	return paginatedData, len(paginatedData)
}
