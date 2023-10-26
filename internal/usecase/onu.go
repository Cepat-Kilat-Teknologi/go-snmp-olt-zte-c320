package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/gosnmp/gosnmp"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/config"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/internal/model"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/internal/repository"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/internal/utils"
	"github.com/rs/zerolog/log"
	"sort"
	"strconv"
)

type OnuUseCaseInterface interface {
	GetByBoardIDAndPonID(ctx context.Context, boardID, ponID int) ([]model.ONUInfoPerBoard, error)
	GetByBoardIDPonIDAndOnuID(ctx context.Context, boardID, ponID, onuID int) (model.ONUCustomerInfo, error)
	GetEmptyOnuID(ctx context.Context, boardID, ponID int) ([]model.OnuID, error)
	UpdateEmptyOnuID(ctx context.Context, boardID, ponID int) error
	GetByBoardIDAndPonIDWithPagination(ctx context.Context, boardID, ponID, page, pageSize int) (
		[]model.ONUInfoPerBoard, int,
	)
}

type onuUsecase struct {
	snmpRepository  repository.SnmpRepositoryInterface
	redisRepository repository.OnuRedisRepositoryInterface
	cfg             *config.Config
}

func NewOnuUsecase(
	snmpRepository repository.SnmpRepositoryInterface, redisRepository repository.OnuRedisRepositoryInterface,
	cfg *config.Config,
) OnuUseCaseInterface {
	return &onuUsecase{
		snmpRepository:  snmpRepository,
		redisRepository: redisRepository,
		cfg:             cfg,
	}
}

func (u *onuUsecase) getOltConfig(boardID, ponID int) (*model.OltConfig, error) {
	// Determine base OID and other OID based on Board ID and PON ID
	switch boardID {
	case 1: // Board 1
		switch ponID {
		case 1: // PON 1
			return &model.OltConfig{
				BaseOID:            u.cfg.OltCfg.BaseOID1,
				OnuIDNameOID:       u.cfg.Board1Pon1.OnuIDNameOID,
				OnuTypeOID:         u.cfg.Board1Pon1.OnuTypeOID,
				OnuSerialNumberOID: u.cfg.Board1Pon1.OnuSerialNumberOID,
				OnuRxPowerOID:      u.cfg.Board1Pon1.OnuRxPowerOID,
				OnuTxPowerOID:      u.cfg.Board1Pon1.OnuTxPowerOID,
				OnuStatusOID:       u.cfg.Board1Pon1.OnuStatusOID,
				OnuIPAddressOID:    u.cfg.Board1Pon1.OnuIPAddressOID,
				OnuDescriptionOID:  u.cfg.Board1Pon1.OnuDescriptionOID,
			}, nil
		case 2: // PON 2
			return &model.OltConfig{
				BaseOID:            u.cfg.OltCfg.BaseOID1,
				OnuIDNameOID:       u.cfg.Board1Pon2.OnuIDNameOID,
				OnuTypeOID:         u.cfg.Board1Pon2.OnuTypeOID,
				OnuSerialNumberOID: u.cfg.Board1Pon2.OnuSerialNumberOID,
				OnuRxPowerOID:      u.cfg.Board1Pon2.OnuRxPowerOID,
				OnuTxPowerOID:      u.cfg.Board1Pon2.OnuTxPowerOID,
				OnuStatusOID:       u.cfg.Board1Pon2.OnuStatusOID,
				OnuIPAddressOID:    u.cfg.Board1Pon2.OnuIPAddressOID,
				OnuDescriptionOID:  u.cfg.Board1Pon2.OnuDescriptionOID,
			}, nil
		case 3: // PON 3
			return &model.OltConfig{
				BaseOID:            u.cfg.OltCfg.BaseOID1,
				OnuIDNameOID:       u.cfg.Board1Pon3.OnuIDNameOID,
				OnuTypeOID:         u.cfg.Board1Pon3.OnuTypeOID,
				OnuSerialNumberOID: u.cfg.Board1Pon3.OnuSerialNumberOID,
				OnuRxPowerOID:      u.cfg.Board1Pon3.OnuRxPowerOID,
				OnuTxPowerOID:      u.cfg.Board1Pon3.OnuTxPowerOID,
				OnuStatusOID:       u.cfg.Board1Pon3.OnuStatusOID,
				OnuIPAddressOID:    u.cfg.Board1Pon3.OnuIPAddressOID,
				OnuDescriptionOID:  u.cfg.Board1Pon3.OnuDescriptionOID,
			}, nil
		case 4: // PON 4
			return &model.OltConfig{
				BaseOID:            u.cfg.OltCfg.BaseOID1,
				OnuIDNameOID:       u.cfg.Board1Pon4.OnuIDNameOID,
				OnuTypeOID:         u.cfg.Board1Pon4.OnuTypeOID,
				OnuSerialNumberOID: u.cfg.Board1Pon4.OnuSerialNumberOID,
				OnuRxPowerOID:      u.cfg.Board1Pon4.OnuRxPowerOID,
				OnuTxPowerOID:      u.cfg.Board1Pon4.OnuTxPowerOID,
				OnuStatusOID:       u.cfg.Board1Pon4.OnuStatusOID,
				OnuIPAddressOID:    u.cfg.Board1Pon4.OnuIPAddressOID,
				OnuDescriptionOID:  u.cfg.Board1Pon4.OnuDescriptionOID,
			}, nil
		case 5: // PON 5
			return &model.OltConfig{
				BaseOID:            u.cfg.OltCfg.BaseOID1,
				OnuIDNameOID:       u.cfg.Board1Pon5.OnuIDNameOID,
				OnuTypeOID:         u.cfg.Board1Pon5.OnuTypeOID,
				OnuSerialNumberOID: u.cfg.Board1Pon5.OnuSerialNumberOID,
				OnuRxPowerOID:      u.cfg.Board1Pon5.OnuRxPowerOID,
				OnuTxPowerOID:      u.cfg.Board1Pon5.OnuTxPowerOID,
				OnuStatusOID:       u.cfg.Board1Pon5.OnuStatusOID,
				OnuIPAddressOID:    u.cfg.Board1Pon5.OnuIPAddressOID,
				OnuDescriptionOID:  u.cfg.Board1Pon5.OnuDescriptionOID,
			}, nil
		case 6: // PON 6
			return &model.OltConfig{
				BaseOID:            u.cfg.OltCfg.BaseOID1,
				OnuIDNameOID:       u.cfg.Board1Pon6.OnuIDNameOID,
				OnuTypeOID:         u.cfg.Board1Pon6.OnuTypeOID,
				OnuSerialNumberOID: u.cfg.Board1Pon6.OnuSerialNumberOID,
				OnuRxPowerOID:      u.cfg.Board1Pon6.OnuRxPowerOID,
				OnuTxPowerOID:      u.cfg.Board1Pon6.OnuTxPowerOID,
				OnuStatusOID:       u.cfg.Board1Pon6.OnuStatusOID,
				OnuIPAddressOID:    u.cfg.Board1Pon6.OnuIPAddressOID,
				OnuDescriptionOID:  u.cfg.Board1Pon6.OnuDescriptionOID,
			}, nil
		case 7: // PON 7
			return &model.OltConfig{
				BaseOID:            u.cfg.OltCfg.BaseOID1,
				OnuIDNameOID:       u.cfg.Board1Pon7.OnuIDNameOID,
				OnuTypeOID:         u.cfg.Board1Pon7.OnuTypeOID,
				OnuSerialNumberOID: u.cfg.Board1Pon7.OnuSerialNumberOID,
				OnuRxPowerOID:      u.cfg.Board1Pon7.OnuRxPowerOID,
				OnuTxPowerOID:      u.cfg.Board1Pon7.OnuTxPowerOID,
				OnuStatusOID:       u.cfg.Board1Pon7.OnuStatusOID,
				OnuIPAddressOID:    u.cfg.Board1Pon7.OnuIPAddressOID,
				OnuDescriptionOID:  u.cfg.Board1Pon7.OnuDescriptionOID,
			}, nil
		case 8: // PON 8
			return &model.OltConfig{
				BaseOID:            u.cfg.OltCfg.BaseOID1,
				OnuIDNameOID:       u.cfg.Board1Pon8.OnuIDNameOID,
				OnuTypeOID:         u.cfg.Board1Pon8.OnuTypeOID,
				OnuSerialNumberOID: u.cfg.Board1Pon8.OnuSerialNumberOID,
				OnuRxPowerOID:      u.cfg.Board1Pon8.OnuRxPowerOID,
				OnuTxPowerOID:      u.cfg.Board1Pon8.OnuTxPowerOID,
				OnuStatusOID:       u.cfg.Board1Pon8.OnuStatusOID,
				OnuIPAddressOID:    u.cfg.Board1Pon8.OnuIPAddressOID,
				OnuDescriptionOID:  u.cfg.Board1Pon8.OnuDescriptionOID,
			}, nil
		default: // Invalid PON ID
			log.Error().Msg("Invalid PON ID")        // Log error message to logger
			return nil, errors.New("invalid PON ID") // Return error
		}
	case 2: // Board 2
		switch ponID {
		case 1: // PON 1
			return &model.OltConfig{
				BaseOID:            u.cfg.OltCfg.BaseOID1,
				OnuIDNameOID:       u.cfg.Board2Pon1.OnuIDNameOID,
				OnuTypeOID:         u.cfg.Board2Pon1.OnuTypeOID,
				OnuSerialNumberOID: u.cfg.Board2Pon1.OnuSerialNumberOID,
				OnuRxPowerOID:      u.cfg.Board2Pon1.OnuRxPowerOID,
				OnuTxPowerOID:      u.cfg.Board2Pon1.OnuTxPowerOID,
				OnuStatusOID:       u.cfg.Board2Pon1.OnuStatusOID,
				OnuIPAddressOID:    u.cfg.Board2Pon1.OnuIPAddressOID,
				OnuDescriptionOID:  u.cfg.Board2Pon1.OnuDescriptionOID,
			}, nil
		case 2: // PON 2
			return &model.OltConfig{
				BaseOID:            u.cfg.OltCfg.BaseOID1,
				OnuIDNameOID:       u.cfg.Board2Pon2.OnuIDNameOID,
				OnuTypeOID:         u.cfg.Board2Pon2.OnuTypeOID,
				OnuSerialNumberOID: u.cfg.Board2Pon2.OnuSerialNumberOID,
				OnuRxPowerOID:      u.cfg.Board2Pon2.OnuRxPowerOID,
				OnuTxPowerOID:      u.cfg.Board2Pon2.OnuTxPowerOID,
				OnuStatusOID:       u.cfg.Board2Pon2.OnuStatusOID,
				OnuIPAddressOID:    u.cfg.Board2Pon2.OnuIPAddressOID,
				OnuDescriptionOID:  u.cfg.Board2Pon2.OnuDescriptionOID,
			}, nil
		case 3: // PON 3
			return &model.OltConfig{
				BaseOID:            u.cfg.OltCfg.BaseOID1,
				OnuIDNameOID:       u.cfg.Board2Pon3.OnuIDNameOID,
				OnuTypeOID:         u.cfg.Board2Pon3.OnuTypeOID,
				OnuSerialNumberOID: u.cfg.Board2Pon3.OnuSerialNumberOID,
				OnuRxPowerOID:      u.cfg.Board2Pon3.OnuRxPowerOID,
				OnuTxPowerOID:      u.cfg.Board2Pon3.OnuTxPowerOID,
				OnuStatusOID:       u.cfg.Board2Pon3.OnuStatusOID,
				OnuIPAddressOID:    u.cfg.Board2Pon3.OnuIPAddressOID,
				OnuDescriptionOID:  u.cfg.Board2Pon3.OnuDescriptionOID,
			}, nil
		case 4: // PON 4
			return &model.OltConfig{
				BaseOID:            u.cfg.OltCfg.BaseOID1,
				OnuIDNameOID:       u.cfg.Board2Pon4.OnuIDNameOID,
				OnuTypeOID:         u.cfg.Board2Pon4.OnuTypeOID,
				OnuSerialNumberOID: u.cfg.Board2Pon4.OnuSerialNumberOID,
				OnuRxPowerOID:      u.cfg.Board2Pon4.OnuRxPowerOID,
				OnuTxPowerOID:      u.cfg.Board2Pon4.OnuTxPowerOID,
				OnuStatusOID:       u.cfg.Board2Pon4.OnuStatusOID,
				OnuIPAddressOID:    u.cfg.Board2Pon4.OnuIPAddressOID,
				OnuDescriptionOID:  u.cfg.Board2Pon4.OnuDescriptionOID,
			}, nil
		case 5: // PON 5
			return &model.OltConfig{
				BaseOID:            u.cfg.OltCfg.BaseOID1,
				OnuIDNameOID:       u.cfg.Board2Pon5.OnuIDNameOID,
				OnuTypeOID:         u.cfg.Board2Pon5.OnuTypeOID,
				OnuSerialNumberOID: u.cfg.Board2Pon5.OnuSerialNumberOID,
				OnuRxPowerOID:      u.cfg.Board2Pon5.OnuRxPowerOID,
				OnuTxPowerOID:      u.cfg.Board2Pon5.OnuTxPowerOID,
				OnuStatusOID:       u.cfg.Board2Pon5.OnuStatusOID,
				OnuIPAddressOID:    u.cfg.Board2Pon5.OnuIPAddressOID,
				OnuDescriptionOID:  u.cfg.Board2Pon5.OnuDescriptionOID,
			}, nil
		case 6: // PON 6
			return &model.OltConfig{
				BaseOID:            u.cfg.OltCfg.BaseOID1,
				OnuIDNameOID:       u.cfg.Board2Pon6.OnuIDNameOID,
				OnuTypeOID:         u.cfg.Board2Pon6.OnuTypeOID,
				OnuSerialNumberOID: u.cfg.Board2Pon6.OnuSerialNumberOID,
				OnuRxPowerOID:      u.cfg.Board2Pon6.OnuRxPowerOID,
				OnuTxPowerOID:      u.cfg.Board2Pon6.OnuTxPowerOID,
				OnuStatusOID:       u.cfg.Board2Pon6.OnuStatusOID,
				OnuIPAddressOID:    u.cfg.Board2Pon6.OnuIPAddressOID,
				OnuDescriptionOID:  u.cfg.Board2Pon6.OnuDescriptionOID,
			}, nil
		case 7: // PON 7
			return &model.OltConfig{
				BaseOID:            u.cfg.OltCfg.BaseOID1,
				OnuIDNameOID:       u.cfg.Board2Pon7.OnuIDNameOID,
				OnuTypeOID:         u.cfg.Board2Pon7.OnuTypeOID,
				OnuSerialNumberOID: u.cfg.Board2Pon7.OnuSerialNumberOID,
				OnuRxPowerOID:      u.cfg.Board2Pon7.OnuRxPowerOID,
				OnuTxPowerOID:      u.cfg.Board2Pon7.OnuTxPowerOID,
				OnuStatusOID:       u.cfg.Board2Pon7.OnuStatusOID,
				OnuIPAddressOID:    u.cfg.Board2Pon7.OnuIPAddressOID,
				OnuDescriptionOID:  u.cfg.Board2Pon7.OnuDescriptionOID,
			}, nil
		case 8: // PON 8
			return &model.OltConfig{
				BaseOID:            u.cfg.OltCfg.BaseOID1,
				OnuIDNameOID:       u.cfg.Board2Pon8.OnuIDNameOID,
				OnuTypeOID:         u.cfg.Board2Pon8.OnuTypeOID,
				OnuSerialNumberOID: u.cfg.Board2Pon8.OnuSerialNumberOID,
				OnuRxPowerOID:      u.cfg.Board2Pon8.OnuRxPowerOID,
				OnuTxPowerOID:      u.cfg.Board2Pon8.OnuTxPowerOID,
				OnuStatusOID:       u.cfg.Board2Pon8.OnuStatusOID,
				OnuIPAddressOID:    u.cfg.Board2Pon8.OnuIPAddressOID,
				OnuDescriptionOID:  u.cfg.Board2Pon8.OnuDescriptionOID,
			}, nil
		default: // Invalid PON ID
			log.Error().Msg("Invalid PON ID")        // Log error message to logger
			return nil, errors.New("invalid PON ID") // Return error
		}
	default: // Invalid Board ID
		log.Error().Msg("Invalid Board ID")        // Log error message to logger
		return nil, errors.New("invalid Board ID") // Return error
	}
}

func (u *onuUsecase) GetByBoardIDAndPonID(ctx context.Context, boardID, ponID int) ([]model.ONUInfoPerBoard, error) {

	ctx, cancel := context.WithCancel(context.Background()) // Create context with cancel function to cancel context
	defer cancel()                                          // Defer cancel function to be called later

	// Log info message to logger
	log.Info().Msg("Get All ONU Information from Board ID: " + strconv.Itoa(boardID) + " and PON ID: " + strconv.Itoa(
		ponID))

	// Get OLT config based on Board ID and PON ID
	oltConfig, err := u.getOltConfig(boardID, ponID)
	if err != nil {
		log.Error().Msg("Failed to get OLT Config: " + err.Error()) // Log error message to logger
		return nil, err                                             // Return error if error is not nil
	}

	// Redis Key
	redisKey := "board_" + strconv.Itoa(boardID) + "_pon_" + strconv.Itoa(ponID)

	// Try to get data from Redis using GetONUInfoList method with context and Redis key as parameter
	cachedOnuData, err := u.redisRepository.GetONUInfoList(ctx, redisKey)
	if err == nil && cachedOnuData != nil {
		log.Info().Msg("Get ONU Information from Redis with Key: " + redisKey) // Log info message to logger
		return cachedOnuData, nil                                              // Return cached data if error is nil and cached data is not nil
	}

	var onuInformationList []model.ONUInfoPerBoard // Create slice to store ONU informationList

	snmpDataMap := make(map[string]gosnmp.SnmpPDU) // Create map to store SNMP data

	/*
		Perform SNMP Walk to get ONU ID and ONU Name
		based on Board ID and PON ID using snmpRepository Walk method
		with context and OID as parameter
	*/

	log.Info().Msg("Get All ONU Information from SNMP Walk Board ID: " + strconv.Itoa(
		boardID) + " and PON ID: " + strconv.Itoa(ponID)) // Log info message to logger

	err = u.snmpRepository.BulkWalk(oltConfig.BaseOID+oltConfig.OnuIDNameOID, func(pdu gosnmp.SnmpPDU) error {
		// Store SNMP data to map with ONU ID as key and PDU as value to be used later
		snmpDataMap[utils.ExtractONUID(pdu.Name)] = pdu // Extract ONU ID from SNMP PDU Name and use it as key in map
		return nil                                      // Return nil error
	})

	if err != nil {
		return nil, err
	}

	/*
		Loop through SNMP data map to get ONU information based on ONU ID and ONU Name stored in map before and store
		it to slice of ONU information list to be returned later to caller function as response data
	*/
	for _, pdu := range snmpDataMap {
		onuInfo := model.ONUInfoPerBoard{
			Board: boardID,                        // Set Board ID to ONU onuInfo struct Board field
			PON:   ponID,                          // Set PON ID to ONU onuInfo  struct PON field
			ID:    utils.ExtractIDOnuID(pdu.Name), // Set ONU ID to ONU onuInfo struct ID field
			Name:  utils.ExtractName(pdu.Value),   // Set ONU Name to ONU onuInfo struct Name field
		}

		// Get ONU Type based on ONU ID and ONU Type OID and store it to ONU onuInfo struct
		onuType, err := u.getONUType(ctx, oltConfig.OnuTypeOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.OnuType = onuType // Set ONU Type to ONU onuInfo struct OnuType field
		}

		// Get ONU Serial Number based on ONU ID and ONU Serial Number OID and store it to ONU onuInfo struct
		onuSerialNumber, err := u.getSerialNumber(ctx, oltConfig.OnuSerialNumberOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.SerialNumber = onuSerialNumber // Set ONU Serial Number to ONU onuInfo struct SerialNumber field
		}

		// Get ONU RX Power based on ONU ID and ONU RX Power OID and store it to ONU onuInfo struct
		onuRXPower, err := u.getRxPower(ctx, oltConfig.OnuRxPowerOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.RXPower = onuRXPower // Set ONU RX Power to ONU onuInfo struct RXPower field
		}

		// Get ONU Status based on ONU ID and ONU Status OID and store it to ONU onuInfo struct
		onuStatus, err := u.getStatus(ctx, oltConfig.OnuStatusOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.Status = onuStatus // Set ONU Status to ONU onuInfo struct Status field
		}

		onuInformationList = append(onuInformationList, onuInfo) // Append ONU onuInfo struct to ONU information list
	}

	// Sort ONU information list based on ONU ID ascending
	sort.Slice(onuInformationList, func(i, j int) bool {
		return onuInformationList[i].ID < onuInformationList[j].ID
	})

	// Save ONU information list to Redis 12 hours
	err = u.redisRepository.SaveONUInfoList(ctx, redisKey, 300, onuInformationList)

	log.Info().Msg("Save ONU Information to Redis with Key: " + redisKey) // Log info message to logger

	if err != nil {
		log.Error().Msg("Failed to save ONU Information to Redis: " + err.Error()) // Log error message to logger
		return nil, err                                                            // Return error if error is not nil
	}

	return onuInformationList, nil // Return ONU information list and nil error
}

func (u *onuUsecase) GetByBoardIDPonIDAndOnuID(ctx context.Context, boardID, ponID, onuID int) (
	model.ONUCustomerInfo, error,
) {
	// Create context with timeout 30 seconds
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Cancel context when function is done

	// Get OLT config based on Board ID and PON ID
	oltConfig, err := u.getOltConfig(boardID, ponID)
	if err != nil {
		log.Error().Msg("Failed to get OLT Config: " + err.Error()) // Log error message to logger
		return model.ONUCustomerInfo{}, err                         // Return error if error is not nil
	}

	// Create a slice of ONUCustomerInfo
	var onuInformationList model.ONUCustomerInfo

	// Create a map to store SNMP Walk results
	snmpDataMap := make(map[string]gosnmp.SnmpPDU)

	log.Info().Msg("Get Detail ONU Information with SNMP Walk from Board ID: " + strconv.Itoa(
		boardID) + " PON ID: " + strconv.Itoa(
		ponID) + " ONU ID: " + strconv.Itoa(onuID))

	// Perform SNMP Walk to get ONU ID and Name using snmpRepository Walk method with timeout context parameter
	err = u.snmpRepository.Walk(oltConfig.BaseOID+oltConfig.OnuIDNameOID+"."+strconv.Itoa(onuID),
		func(pdu gosnmp.SnmpPDU) error {
			// Save SNMP Walk result in map with ID as key and Name as value (extracted from SNMP PDU)
			snmpDataMap[utils.ExtractONUID(pdu.Name)] = pdu // Extract ONU ID from SNMP PDU Name and use it as key in map
			return nil
		})

	if err != nil {
		log.Error().Msg("Failed to walk OID: " + err.Error())            // Log error message to logger
		return model.ONUCustomerInfo{}, errors.New("failed to walk OID") // Return error
	}

	/*
		Loop through SNMP data map to get ONU information based on ONU ID and ONU Name stored in map before and store
		it to slice of ONU information list to be returned later to caller function as response data
	*/
	for _, pdu := range snmpDataMap {
		onuInfo := model.ONUCustomerInfo{
			Board: boardID,                        // Set Board ID to ONU onuInfo struct Board field
			PON:   ponID,                          // Set PON ID to ONU onuInfo  struct PON field
			ID:    utils.ExtractIDOnuID(pdu.Name), // Set ONU ID (extracted from SNMP PDU) to onuInfo variable (ONU ID)
			Name:  utils.ExtractName(pdu.Value),   // Set ONU Name (extracted from SNMP PDU) to onuInfo variable (ONU Name)
		}

		// Get Data ONU Type from SNMP Walk using getONUType method
		onuType, err := u.getONUType(ctx, oltConfig.OnuTypeOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.OnuType = onuType // Set ONU Type from SNMP Walk result if no error to onuInfo variable (ONU Type)
		}

		// Get Data ONU Serial Number from SNMP Walk using getSerialNumber method
		onuSerialNumber, err := u.getSerialNumber(ctx, oltConfig.OnuSerialNumberOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.SerialNumber = onuSerialNumber // Set ONU Serial Number from SNMP Walk result to onuInfo variable (ONU Serial Number)
		}

		// Get Data ONU RX Power from SNMP Walk using getRxPower method
		onuRXPower, err := u.getRxPower(ctx, oltConfig.OnuRxPowerOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.RXPower = onuRXPower // Set ONU RX Power from SNMP Walk result to onuInfo variable (ONU RX Power)
		}

		// Get Data ONU TX Power from SNMP Walk using getTxPower method
		onuTXPower, err := u.getTxPower(ctx, oltConfig.OnuTxPowerOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.TXPower = onuTXPower // Set ONU TX Power from SNMP Walk result to onuInfo variable (ONU TX Power)
		}

		// Get Data ONU Status from SNMP Walk using getStatus method
		onuStatus, err := u.getStatus(ctx, oltConfig.OnuStatusOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.Status = onuStatus // Set ONU Status from SNMP Walk result to onuInfo variable (ONU Status)
		}

		// Get Data ONU IP Address from SNMP Walk using getIPAddress method
		onuIPAddress, err := u.getIPAddress(ctx, oltConfig.OnuIPAddressOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.IPAddress = onuIPAddress // Set ONU IP Address from SNMP Walk result to onuInfo variable (ONU IP Address)
		}

		// Get Data ONU Description from SNMP Walk using getDescription method
		onuDescription, err := u.getDescription(ctx, oltConfig.OnuDescriptionOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.Description = onuDescription // Set ONU Description from SNMP Walk result to onuInfo variable (ONU Description)
		}

		onuInformationList = onuInfo // Append onuInfo variable to onuInformationList slice
	}

	return onuInformationList, nil
}

func (u *onuUsecase) getONUType(ctx context.Context, OnuTypeOID, onuID string) (string, error) {

	var onuType string // Variable to store ONU Type

	ctx, cancel := context.WithCancel(context.Background()) // Create context with timeout 30 seconds
	defer cancel()                                          // Cancel context when function is done

	baseOID := u.cfg.OltCfg.BaseOID2 // Base OID variable get from config

	// Perform SNMP Get to get ONU Type using snmpRepository Get method with timeout context parameter
	oids := []string{baseOID + OnuTypeOID + "." + onuID}
	result, err := u.snmpRepository.Get(oids)
	if err != nil {
		log.Error().Msg("Failed to perform SNMP Get to get ONU Type: " + err.Error()) // Log error message to logger
		return "", errors.New("failed to perform SNMP Get")                           // Return error
	}

	// Check if the result contains the expected OID
	if len(result.Variables) > 0 {
		onuType = utils.ExtractName(result.Variables[0].Value) // Extract ONU Type from the result
	} else {
		log.Error().Msg("Failed to get ONU Type: No variables in the response")
		return "", errors.New("no variables in the response")
	}

	return onuType, nil // Return ONU Type
}

func (u *onuUsecase) getSerialNumber(ctx context.Context, OnuSerialNumberOID, onuID string) (string, error) {

	var onuSerialNumber string // Variable to store ONU Serial Number

	ctx, cancel := context.WithCancel(context.Background()) // Create context with timeout 30 seconds
	defer cancel()                                          // Cancel context when function is done

	baseOID := u.cfg.OltCfg.BaseOID1 // Base OID variable get from config

	// Perform SNMP Get to get ONU Serial Number using snmpRepository Get method with timeout context parameter
	oids := []string{baseOID + OnuSerialNumberOID + "." + onuID}
	result, err := u.snmpRepository.Get(oids)
	if err != nil {
		log.Error().Msg("Failed to perform SNMP Get for serial number: " + err.Error()) // Log error message to logger
		return "", errors.New("failed to perform SNMP Get")                             // Return error
	}

	// Check if the result contains the expected OID
	if len(result.Variables) > 0 {
		onuSerialNumber = utils.ExtractSerialNumber(result.Variables[0].Value) // Extract ONU Serial Number from the result
	} else {
		log.Error().Msg("Failed to get ONU Serial Number: No variables in the response")
		return "", errors.New("no variables in the response")
	}

	return onuSerialNumber, nil // Return ONU Serial Number
}

func (u *onuUsecase) getRxPower(ctx context.Context, OnuRxPowerOID, onuID string) (string, error) {

	var onuRxPower string // Variable to store ONU RX Power

	ctx, cancel := context.WithCancel(context.Background()) // Create context with timeout 30 seconds
	defer cancel()                                          // Cancel context when function is done

	baseOID := u.cfg.OltCfg.BaseOID1 // Base OID variable get from config

	// Perform SNMP Get to get ONU RX Power using snmpRepository Get method with timeout context parameter
	oids := []string{baseOID + OnuRxPowerOID + "." + onuID + ".1"}
	result, err := u.snmpRepository.Get(oids)
	if err != nil {
		log.Error().Msg("Failed to perform SNMP Get for RX Power: " + err.Error()) // Log error message to logger
		return "", errors.New("failed to perform SNMP Get")                        // Return error
	}

	// Check if the result contains the expected OID
	if len(result.Variables) > 0 {
		onuRxPower, _ = utils.ConvertAndMultiply(result.Variables[0].Value) // Extract ONU RX Power from the result
	} else {
		log.Error().Msg("Failed to get ONU RX Power: No variables in the response")
		return "", errors.New("no variables in the response")
	}

	return onuRxPower, nil // Return ONU RX Power
}

func (u *onuUsecase) getTxPower(ctx context.Context, OnuTxPowerOID, onuID string) (string, error) {

	var onuTxPower string // Variable to store ONU TX Power

	ctx, cancel := context.WithCancel(context.Background()) // Create context with timeout 30 seconds
	defer cancel()                                          // Cancel context when function is done

	baseOID := u.cfg.OltCfg.BaseOID2 // Base OID variable get from config

	// Perform SNMP Get to get ONU TX Power using snmpRepository Get method with timeout context parameter
	oids := []string{baseOID + OnuTxPowerOID + "." + onuID + ".1"}
	result, err := u.snmpRepository.Get(oids)
	if err != nil {
		log.Error().Msg("Failed to perform SNMP Get for TX Power: " + err.Error()) // Log error message to logger
		return "", errors.New("failed to perform SNMP Get")                        // Return error
	}

	// Check if the result contains the expected OID
	if len(result.Variables) > 0 {
		onuTxPower, _ = utils.ConvertAndMultiply(result.Variables[0].Value) // Extract ONU TX Power from the result
	} else {
		log.Error().Msg("Failed to get ONU TX Power: No variables in the response")
		return "", errors.New("no variables in the response")
	}

	return onuTxPower, nil // Return ONU TX Power
}

func (u *onuUsecase) getStatus(ctx context.Context, OnuStatusOID, onuID string) (string, error) {

	var onuStatus string // Variable to store ONU Status

	ctx, cancel := context.WithCancel(context.Background()) // Create context with timeout 30 seconds
	defer cancel()                                          // Cancel context when function is done

	baseOID := u.cfg.OltCfg.BaseOID1 // Base OID variable get from config

	// Perform SNMP Get to get ONU Status using snmpRepository Get method with timeout context parameter
	oids := []string{baseOID + OnuStatusOID + "." + onuID}
	result, err := u.snmpRepository.Get(oids)
	if err != nil {
		log.Error().Msg("Failed to perform SNMP Get for status: " + err.Error()) // Log error message to logger
		return "", errors.New("failed to perform SNMP Get")                      // Return error
	}

	// Check if the result contains the expected OID
	if len(result.Variables) > 0 {
		onuStatus = utils.ExtractAndGetStatus(result.Variables[0].Value) // Extract ONU Status from the result
	} else {
		log.Error().Msg("Failed to get ONU Status: No variables in the response")
		return "", errors.New("no variables in the response")
	}

	return onuStatus, nil // Return ONU Status
}

func (u *onuUsecase) getIPAddress(ctx context.Context, OnuIPAddressOID, onuID string) (string, error) {

	var onuIPAddress string // Variable to store ONU IP Address

	ctx, cancel := context.WithCancel(context.Background()) // Create context with timeout 30 seconds
	defer cancel()                                          // Cancel context when function is done

	baseOID := u.cfg.OltCfg.BaseOID2 // Base OID variable get from config

	// Perform SNMP Get to get ONU IP Address using snmpRepository Get method with timeout context parameter
	oids := []string{baseOID + OnuIPAddressOID + "." + onuID + ".1"}
	result, err := u.snmpRepository.Get(oids)
	if err != nil {
		log.Error().Msg("Failed to perform SNMP Get for IP Address: " + err.Error()) // Log error message to logger
		return "", errors.New("failed to perform SNMP Get")                          // Return error
	}

	// Check if the result contains the expected OID
	if len(result.Variables) > 0 {
		onuIPAddress = utils.ExtractName(result.Variables[0].Value) // Extract ONU IP Address from the result
	} else {
		log.Error().Msg("Failed to get ONU IP Address: No variables in the response")
		return "", errors.New("no variables in the response")
	}

	return onuIPAddress, nil // Return ONU IP Address
}

func (u *onuUsecase) getDescription(ctx context.Context, OnuDescriptionOID, onuID string) (string, error) {

	var onuDescription string // Variable to store ONU Description

	ctx, cancel := context.WithCancel(context.Background()) // Create context with timeout 30 seconds
	defer cancel()                                          // Cancel context when function is done

	baseOID := u.cfg.OltCfg.BaseOID1 // Base OID variable get from config

	// Perform SNMP Get to get ONU Description using snmpRepository Get method with timeout context parameter
	oids := []string{baseOID + OnuDescriptionOID + "." + onuID}
	result, err := u.snmpRepository.Get(oids)
	if err != nil {
		log.Error().Msg("Failed to perform SNMP Get for description: " + err.Error()) // Log error message to logger
		return "", errors.New("failed to perform SNMP Get")                           // Return error
	}

	// Check if the result contains the expected OID
	if len(result.Variables) > 0 {
		onuDescription = utils.ExtractName(result.Variables[0].Value) // Extract ONU Description from the result
	} else {
		log.Error().Msg("Failed to get ONU Description: No variables in the response")
		return "", errors.New("no variables in the response")
	}

	return onuDescription, nil // Return ONU Description
}

func (u *onuUsecase) GetEmptyOnuID(ctx context.Context, boardID, ponID int) ([]model.OnuID, error) {
	ctx, cancel := context.WithCancel(context.Background()) // Create context with timeout 30 seconds
	defer cancel()

	// Get OLT config based on Board ID and PON ID
	oltConfig, err := u.getOltConfig(boardID, ponID)
	if err != nil {
		log.Error().Msg("Failed to get OLT Config for Get Empty ONU ID: " + err.Error()) // Log error message to logger
		return nil, err                                                                  // Return error if error is not nil
	}

	//Redis Key
	redisKey := "board_" + strconv.Itoa(boardID) + "_pon_" + strconv.Itoa(ponID) + "_empty_onu_id"

	//Try to get data from Redis using GetOnuIDCtx method with context and Redis key as parameter
	cachedOnuData, err := u.redisRepository.GetOnuIDCtx(ctx, redisKey)
	if err == nil && cachedOnuData != nil {
		log.Info().Msg("Get Empty ONU ID from Redis with Key: " + redisKey) // Log info message to logger
		// If data exist in Redis, then return data from Redis
		return cachedOnuData, nil
	}

	// Perform SNMP Walk to get ONU ID and ONU Name
	snmpOID := oltConfig.BaseOID + oltConfig.OnuIDNameOID // SNMP OID variable
	emptyOnuIDList := make([]model.OnuID, 0)              // Create a slice of ONU ID

	log.Info().Msg("Get Empty ONU ID with SNMP Walk from Board ID: " + strconv.Itoa(
		boardID) + " and PON ID: " + strconv.Itoa(ponID)) // Log info message to logger

	// Perform SNMP BulkWalk to get ONU ID and Name using snmpRepository BulkWalk method with timeout context parameter
	err = u.snmpRepository.BulkWalk(snmpOID, func(pdu gosnmp.SnmpPDU) error {
		idOnuID := utils.ExtractIDOnuID(pdu.Name) // Extract ONU ID from SNMP PDU Name

		// Append ONU information to the emptyOnuIDList
		emptyOnuIDList = append(emptyOnuIDList, model.OnuID{
			Board: boardID, // Set Board ID to ONU onuInfo struct Board field
			PON:   ponID,   // Set PON ID to ONU onuInfo  struct PON field
			ID:    idOnuID, // Set ONU ID (extracted from SNMP PDU) to onuInfo variable (ONU ID)
		})

		return nil
	})

	if err != nil {
		log.Error().Msg("Failed to perform SNMP Walk get empty ONU ID: " + err.Error()) // Log error message to logger
		return nil, err
	}

	// Create a map to store numbers to be deleted
	numbersToRemove := make(map[int]bool)

	// Loop through emptyOnuIDList to get the numbers to be deleted
	for _, onuInfo := range emptyOnuIDList {
		numbersToRemove[onuInfo.ID] = true
	}

	// Create a new slice to hold the board_id, pon_id and onu_id data without the numbers to be deleted
	emptyOnuIDList = emptyOnuIDList[:0]

	// Loop through 128 numbers to get the numbers to be deleted
	for i := 1; i <= 128; i++ {
		if _, ok := numbersToRemove[i]; !ok {
			emptyOnuIDList = append(emptyOnuIDList, model.OnuID{
				Board: boardID, // Set Board ID to ONU onuInfo struct Board field
				PON:   ponID,   // Set PON ID to ONU onuInfo  struct PON field
				ID:    i,       // Number 1-128 that is not in the numbers to be deleted
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
		log.Error().Msg("Failed to set data to Redis: " + err.Error()) // Log error message to logger
		return nil, err
	}

	log.Info().Msg("Save Empty ONU ID to Redis with Key: " + redisKey) // Log info message to logger

	return emptyOnuIDList, nil
}

func (u *onuUsecase) UpdateEmptyOnuID(ctx context.Context, boardID, ponID int) error {
	ctx, cancel := context.WithCancel(context.Background()) // Create context with timeout 30 seconds
	defer cancel()

	// Get OLT config based on Board ID and PON ID
	oltConfig, err := u.getOltConfig(boardID, ponID)
	if err != nil {
		log.Error().Msg("Failed to get OLT Config: " + err.Error()) // Log error message to logger
		return err                                                  // Return error if error is not nil
	}

	// Perform SNMP Walk to get ONU ID and ONU Name
	snmpOID := oltConfig.BaseOID + oltConfig.OnuIDNameOID // SNMP OID variable
	emptyOnuIDList := make([]model.OnuID, 0)              // Create a slice of ONU ID

	log.Info().Msg("Get Empty ONU ID with SNMP Walk from Board ID: " + strconv.Itoa(
		boardID) + " and PON ID: " + strconv.
		Itoa(ponID)) // Log info message to logger

	// Perform SNMP BulkWalk to get ONU ID and Name using snmpRepository BulkWalk method with timeout context parameter
	err = u.snmpRepository.BulkWalk(snmpOID, func(pdu gosnmp.SnmpPDU) error {
		idOnuID := utils.ExtractIDOnuID(pdu.Name) // Extract ONU ID from SNMP PDU Name

		// Append ONU information to the emptyOnuIDList
		emptyOnuIDList = append(emptyOnuIDList, model.OnuID{
			Board: boardID, // Set Board ID to ONU onuInfo struct Board field
			PON:   ponID,   // Set PON ID to ONU onuInfo  struct PON field
			ID:    idOnuID, // Set ONU ID (extracted from SNMP PDU) to onuInfo variable (ONU ID)
		})

		return nil
	})

	if err != nil {
		return errors.New("failed to perform SNMP Walk")
	}

	// Create a map to store numbers to be deleted
	numbersToRemove := make(map[int]bool)

	// Loop through emptyOnuIDList to get the numbers to be deleted
	for _, onuInfo := range emptyOnuIDList {
		numbersToRemove[onuInfo.ID] = true
	}

	// Create a new slice to hold the board_id, pon_id and onu_id data without the numbers to be deleted
	emptyOnuIDList = emptyOnuIDList[:0]

	// Loop through 128 numbers to get the numbers to be deleted
	for i := 1; i <= 128; i++ {
		if _, ok := numbersToRemove[i]; !ok {
			emptyOnuIDList = append(emptyOnuIDList, model.OnuID{
				Board: boardID, // Set Board ID to ONU onuInfo struct Board field
				PON:   ponID,   // Set PON ID to ONU onuInfo  struct PON field
				ID:    i,       // Number 1-128 that is not in the numbers to be deleted
			})
		}
	}

	// Sort by ID ascending
	sort.Slice(emptyOnuIDList, func(i, j int) bool {
		return emptyOnuIDList[i].ID < emptyOnuIDList[j].ID
	})

	//Redis Key
	redisKey := "board_" + strconv.Itoa(boardID) + "_pon_" + strconv.Itoa(ponID) + "_empty_onu_id"

	// Set data to Redis using SetOnuIDCtx method with context, Redis key and data as parameter
	err = u.redisRepository.SetOnuIDCtx(ctx, redisKey, 300, emptyOnuIDList)
	if err != nil {
		log.Error().Msg("Failed to set data to Redis: " + err.Error()) // Log error message to logger
		return errors.New("failed to set data to Redis")
	}

	log.Info().Msg("Save Update Empty ONU ID to Redis with Key: " + redisKey) // Log info message to logger

	return nil
}

func (u *onuUsecase) GetByBoardIDAndPonIDWithPagination(
	ctx context.Context, boardID, ponID, pageIndex,
	pageSize int,
) (
	[]model.ONUInfoPerBoard,
	int,
) {

	ctx, cancel := context.WithCancel(context.Background()) // Create context with timeout 30 seconds
	defer cancel()                                          // Cancel context when function is done

	// Get OLT config based on Board ID and PON ID
	oltConfig, err := u.getOltConfig(boardID, ponID)
	if err != nil {
		return nil, 0 // Return error if error is not nil
	}

	fmt.Println("boardID: ", boardID)
	fmt.Println("ponID:", ponID)
	fmt.Println("pageIndex:", pageIndex)
	fmt.Println("pageSize:", pageSize)

	var onlyOnuIDList []model.OnuOnlyID // Create slice to store ONU ID list

	//// Redis Key
	//redisKey := "board_" + strconv.Itoa(boardID) + "_pon_" + strconv.Itoa(ponID)
	//
	//// Try to get data from Redis using GetOnuIDCtx method with context and Redis key as parameter
	//cachedOnuData, err := u.redisRepository.GetOnlyOnuIDCtx(ctx, redisKey)
	//if err == nil && cachedOnuData != nil {
	//	// If data exist in Redis, then return data from Redis
	//	onlyOnuIDList = cachedOnuData
	//}

	snmpOID := oltConfig.BaseOID + oltConfig.OnuIDNameOID // SNMP OID variable

	// If data not exist in Redis, then get data from SNMP
	if len(onlyOnuIDList) == 0 {
		// Perform SNMP Walk to get ONU ID and ONU Name based on Board ID and PON ID using snmpRepository Walk method
		// with context and OID as parameter
		err := u.snmpRepository.Walk(snmpOID, func(pdu gosnmp.SnmpPDU) error {
			// Append ONU information to the onlyOnuIDList
			onlyOnuIDList = append(onlyOnuIDList, model.OnuOnlyID{
				ID: utils.ExtractIDOnuID(pdu.Name), // Extract ONU ID from SNMP PDU Name
			})
			return nil
		})

		if err != nil {
			return nil, 0 // Return error if error is not nil
		}

		// Sort ONU ID list based on ONU ID ascending
		sort.Slice(onlyOnuIDList, func(i, j int) bool {
			return onlyOnuIDList[i].ID < onlyOnuIDList[j].ID
		})

		//// Set data to Redis
		//err = u.redisRepository.SaveOnlyOnuIDCtx(ctx, redisKey, 300, onlyOnuIDList)
		//
		//if err != nil {
		//	return nil, 0
		//}
	} else {
		fmt.Println("data from redis")
	}

	// Menghitung indeks item pertama yang akan diambil
	startIndex := pageIndex * pageSize

	// Menghitung indeks item terakhir yang akan diambil
	endIndex := startIndex + pageSize

	// Jika indeks item terakhir yang akan diambil lebih besar dari jumlah item, maka indeks item terakhir yang akan diambil adalah jumlah item
	if endIndex > len(onlyOnuIDList) {
		endIndex = len(onlyOnuIDList)
	}

	// Ambil Data ONU ID yang akan ditampilkan berdasarkan indeks item pertama dan indeks item terakhir dari data onlyOnuIDList
	onlyOnuIDList = onlyOnuIDList[startIndex:endIndex]

	// print onuIDList
	fmt.Println("onlyOnuIDListPages", onlyOnuIDList)

	var onuInformationList []model.ONUInfoPerBoard // Create slice to store ONU informationList
	var count int

	var currentIndex int

	// Loop through onlyOnuIDList to get ONU information based on ONU ID
	for _, onuInfo := range onlyOnuIDList {
		if currentIndex >= startIndex {

			// Perform SNMP Walk to get ONU Type based on Board ID, PON ID and ONU ID using snmpRepository Walk method
			// with context and OID as parameter
			err := u.snmpRepository.Walk(oltConfig.BaseOID+oltConfig.OnuTypeOID+"."+strconv.Itoa(onuInfo.ID),
				func(pdu gosnmp.SnmpPDU) error {
					// Append ONU information to the onuInformationList
					onuInformationList = append(onuInformationList, model.ONUInfoPerBoard{
						Board: boardID, // Set Board ID to ONU onuInfo struct Board field
						PON:   ponID,   // Set PON ID to ONU onuInfo  struct PON field
						ID:    onuInfo.ID,
					})
					return nil
				})

			if err != nil {
				return nil, 0 // Return error if error is not nil
			}

			// Perform SNMP Walk to get ONU Serial Number based on Board ID,
			//PON ID and ONU ID using snmpRepository Walk method
			// with context and OID as parameter
			err = u.snmpRepository.Walk(oltConfig.BaseOID+oltConfig.OnuSerialNumberOID+"."+strconv.Itoa(onuInfo.ID),
				func(pdu gosnmp.SnmpPDU) error {
					// Set ONU Serial Number to ONU onuInfo struct SerialNumber field
					onuInformationList[currentIndex].SerialNumber = utils.ExtractName(pdu.Value)
					return nil
				})

			if err != nil {
				return nil, 0 // Return error if error is not nil
			}

			// Perform SNMP Walk to get ONU RX Power based on Board ID,
			//PON ID and ONU ID using snmpRepository Walk method
			// with context and OID as parameter
			err = u.snmpRepository.Walk(oltConfig.BaseOID+oltConfig.OnuRxPowerOID+"."+strconv.Itoa(onuInfo.ID),
				func(pdu gosnmp.SnmpPDU) error {
					// Set ONU RX Power to ONU onuInfo struct RXPower field
					onuInformationList[currentIndex].RXPower = utils.ExtractName(pdu.Value)
					return nil
				})

			if err != nil {
				return nil, 0 // Return error if error is not nil
			}

			// Perform SNMP Walk to get ONU Status based on Board ID, PON ID and ONU ID using snmpRepository Walk method
			// with context and OID as parameter
			err = u.snmpRepository.Walk(oltConfig.BaseOID+oltConfig.OnuStatusOID+"."+strconv.Itoa(onuInfo.ID),
				func(pdu gosnmp.SnmpPDU) error {
					// Set ONU Status to ONU onuInfo struct Status field
					onuInformationList[currentIndex].Status = utils.ExtractName(pdu.Value)
					return nil
				})

			if err != nil {
				return nil, 0 // Return error if error is not nil
			}
		}

		currentIndex++

		if currentIndex >= endIndex {
			break
		}
	}

	// Get Count from onlyOnuIDList
	count = len(onlyOnuIDList)

	// Sort ONU information list based on ONU ID ascending
	sort.Slice(onuInformationList, func(i, j int) bool {
		return onuInformationList[i].ID < onuInformationList[j].ID
	})

	// Jika indeks item pertama lebih besar dari jumlah data, maka kembalikan data kosong
	if startIndex >= count {
		return nil, 0
	}

	// Jika indeks item terakhir yang akan diambil melebihi jumlah data, maka atur indeks item terakhir menjadi jumlah data
	if endIndex > count {
		endIndex = count
	}

	// Ambil data ONU yang sesuai dengan halaman yang diminta
	paginatedData := onuInformationList[startIndex:endIndex]

	fmt.Println(paginatedData)

	// Kembalikan data halaman bersama dengan jumlah total data yang tersedia
	return paginatedData, count
}

func (u *onuUsecase) GetEmptyOnuIDQueue(ctx context.Context, boardID, ponID int) ([]model.OnuID, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Generate the Redis key
	redisKey := fmt.Sprintf("board_%d_pon_%d_empty_onu_id", boardID, ponID)

	// Try to get data from Redis
	cachedOnuData, err := u.redisRepository.GetOnuIDCtx(ctx, redisKey)
	if err == nil && cachedOnuData != nil {
		return cachedOnuData, nil
	}

	// Get OLT config based on Board ID and PON ID
	oltConfig, err := u.getOltConfig(boardID, ponID)
	if err != nil {
		return nil, err
	}

	// Perform SNMP Walk to get ONU ID and ONU Name
	snmpOID := oltConfig.BaseOID + oltConfig.OnuIDNameOID
	emptyOnuIDList := make([]model.OnuID, 0)

	// Perform SNMP Walk using snmpRepository Walk method with timeout context parameter
	err = u.snmpRepository.Walk(snmpOID, func(pdu gosnmp.SnmpPDU) error {
		idOnuID := utils.ExtractIDOnuID(pdu.Name)
		emptyOnuIDList = append(emptyOnuIDList, model.OnuID{
			Board: boardID,
			PON:   ponID,
			ID:    idOnuID,
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

	// Create a new slice to hold the board_id, pon_id, and onu_id data without the numbers to be deleted
	var filteredEmptyOnuIDList []model.OnuID

	// Loop through 128 numbers to get the numbers to be deleted
	for i := 1; i <= 128; i++ {
		if _, ok := numbersToRemove[i]; !ok {
			filteredEmptyOnuIDList = append(filteredEmptyOnuIDList, model.OnuID{
				Board: boardID,
				PON:   ponID,
				ID:    i,
			})
		}
	}

	// Set data to Redis
	err = u.redisRepository.SetOnuIDCtx(ctx, redisKey, 300, filteredEmptyOnuIDList)
	if err != nil {
		return nil, err
	}

	// Sort by ID ascending
	sort.Slice(filteredEmptyOnuIDList, func(i, j int) bool {
		return filteredEmptyOnuIDList[i].ID < filteredEmptyOnuIDList[j].ID
	})

	return filteredEmptyOnuIDList, nil
}
