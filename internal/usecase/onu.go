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
	"time"
)

type OnuUseCaseInterface interface {
	GetByBoardIDAndPonID(ctx context.Context, boardID, ponID int) ([]model.ONUInfoPerBoard, error)
	GetByBoardIDPonIDAndOnuID(boardID, ponID, onuID int) (model.ONUCustomerInfo, error)
	GetEmptyOnuID(ctx context.Context, boardID, ponID int) ([]model.OnuID, error)
	GetOnuIDAndSerialNumber(boardID, ponID int) ([]model.OnuSerialNumber, error)
	UpdateEmptyOnuID(ctx context.Context, boardID, ponID int) error
	GetByBoardIDAndPonIDWithPagination(boardID, ponID, page, pageSize int) (
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

// getOltInfo is a function to get OLT information
func (u *onuUsecase) getOltConfig(boardID, ponID int) (*model.OltConfig, error) {
	cfg, err := u.getBoardConfig(boardID, ponID)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}
	return cfg, nil
}

// getBoardConfig is a function to get board configuration
func (u *onuUsecase) getBoardConfig(boardID, ponID int) (*model.OltConfig, error) {
	switch boardID {
	case 1:
		return u.getBoard1Config(ponID), nil
	case 2:
		return u.getBoard2Config(ponID), nil
	default:
		return nil, errors.New("invalid Board ID")
	}
}

// getBoard1Config is a function to get board 1 configuration
func (u *onuUsecase) getBoard1Config(ponID int) *model.OltConfig {
	// Define the configuration for Board 1
	switch ponID {
	case 1:
		return &model.OltConfig{
			BaseOID:                   u.cfg.OltCfg.BaseOID1,
			OnuIDNameOID:              u.cfg.Board1Pon1.OnuIDNameOID,
			OnuTypeOID:                u.cfg.Board1Pon1.OnuTypeOID,
			OnuSerialNumberOID:        u.cfg.Board1Pon1.OnuSerialNumberOID,
			OnuRxPowerOID:             u.cfg.Board1Pon1.OnuRxPowerOID,
			OnuTxPowerOID:             u.cfg.Board1Pon1.OnuTxPowerOID,
			OnuStatusOID:              u.cfg.Board1Pon1.OnuStatusOID,
			OnuIPAddressOID:           u.cfg.Board1Pon1.OnuIPAddressOID,
			OnuDescriptionOID:         u.cfg.Board1Pon1.OnuDescriptionOID,
			OnuLastOnlineOID:          u.cfg.Board1Pon1.OnuLastOnlineOID,
			OnuLastOfflineOID:         u.cfg.Board1Pon1.OnuLastOfflineOID,
			OnuLastOfflineReasonOID:   u.cfg.Board1Pon1.OnuLastOfflineReasonOID,
			OnuGponOpticalDistanceOID: u.cfg.Board1Pon1.OnuGponOpticalDistanceOID,
		}
	case 2:
		return &model.OltConfig{
			BaseOID:                   u.cfg.OltCfg.BaseOID1,
			OnuIDNameOID:              u.cfg.Board1Pon2.OnuIDNameOID,
			OnuTypeOID:                u.cfg.Board1Pon2.OnuTypeOID,
			OnuSerialNumberOID:        u.cfg.Board1Pon2.OnuSerialNumberOID,
			OnuRxPowerOID:             u.cfg.Board1Pon2.OnuRxPowerOID,
			OnuTxPowerOID:             u.cfg.Board1Pon2.OnuTxPowerOID,
			OnuStatusOID:              u.cfg.Board1Pon2.OnuStatusOID,
			OnuIPAddressOID:           u.cfg.Board1Pon2.OnuIPAddressOID,
			OnuDescriptionOID:         u.cfg.Board1Pon2.OnuDescriptionOID,
			OnuLastOnlineOID:          u.cfg.Board1Pon2.OnuLastOnlineOID,
			OnuLastOfflineOID:         u.cfg.Board1Pon2.OnuLastOfflineOID,
			OnuLastOfflineReasonOID:   u.cfg.Board1Pon2.OnuLastOfflineReasonOID,
			OnuGponOpticalDistanceOID: u.cfg.Board1Pon2.OnuGponOpticalDistanceOID,
		}
	case 3: // PON 3
		return &model.OltConfig{
			BaseOID:                   u.cfg.OltCfg.BaseOID1,
			OnuIDNameOID:              u.cfg.Board1Pon3.OnuIDNameOID,
			OnuTypeOID:                u.cfg.Board1Pon3.OnuTypeOID,
			OnuSerialNumberOID:        u.cfg.Board1Pon3.OnuSerialNumberOID,
			OnuRxPowerOID:             u.cfg.Board1Pon3.OnuRxPowerOID,
			OnuTxPowerOID:             u.cfg.Board1Pon3.OnuTxPowerOID,
			OnuStatusOID:              u.cfg.Board1Pon3.OnuStatusOID,
			OnuIPAddressOID:           u.cfg.Board1Pon3.OnuIPAddressOID,
			OnuDescriptionOID:         u.cfg.Board1Pon3.OnuDescriptionOID,
			OnuLastOnlineOID:          u.cfg.Board1Pon3.OnuLastOnlineOID,
			OnuLastOfflineOID:         u.cfg.Board1Pon3.OnuLastOfflineOID,
			OnuLastOfflineReasonOID:   u.cfg.Board1Pon3.OnuLastOfflineReasonOID,
			OnuGponOpticalDistanceOID: u.cfg.Board1Pon3.OnuGponOpticalDistanceOID,
		}
	case 4: // PON 4
		return &model.OltConfig{
			BaseOID:                   u.cfg.OltCfg.BaseOID1,
			OnuIDNameOID:              u.cfg.Board1Pon4.OnuIDNameOID,
			OnuTypeOID:                u.cfg.Board1Pon4.OnuTypeOID,
			OnuSerialNumberOID:        u.cfg.Board1Pon4.OnuSerialNumberOID,
			OnuRxPowerOID:             u.cfg.Board1Pon4.OnuRxPowerOID,
			OnuTxPowerOID:             u.cfg.Board1Pon4.OnuTxPowerOID,
			OnuStatusOID:              u.cfg.Board1Pon4.OnuStatusOID,
			OnuIPAddressOID:           u.cfg.Board1Pon4.OnuIPAddressOID,
			OnuDescriptionOID:         u.cfg.Board1Pon4.OnuDescriptionOID,
			OnuLastOnlineOID:          u.cfg.Board1Pon4.OnuLastOnlineOID,
			OnuLastOfflineOID:         u.cfg.Board1Pon4.OnuLastOfflineOID,
			OnuLastOfflineReasonOID:   u.cfg.Board1Pon4.OnuLastOfflineReasonOID,
			OnuGponOpticalDistanceOID: u.cfg.Board1Pon4.OnuGponOpticalDistanceOID,
		}
	case 5: // PON 5
		return &model.OltConfig{
			BaseOID:                   u.cfg.OltCfg.BaseOID1,
			OnuIDNameOID:              u.cfg.Board1Pon5.OnuIDNameOID,
			OnuTypeOID:                u.cfg.Board1Pon5.OnuTypeOID,
			OnuSerialNumberOID:        u.cfg.Board1Pon5.OnuSerialNumberOID,
			OnuRxPowerOID:             u.cfg.Board1Pon5.OnuRxPowerOID,
			OnuTxPowerOID:             u.cfg.Board1Pon5.OnuTxPowerOID,
			OnuStatusOID:              u.cfg.Board1Pon5.OnuStatusOID,
			OnuIPAddressOID:           u.cfg.Board1Pon5.OnuIPAddressOID,
			OnuDescriptionOID:         u.cfg.Board1Pon5.OnuDescriptionOID,
			OnuLastOnlineOID:          u.cfg.Board1Pon5.OnuLastOnlineOID,
			OnuLastOfflineOID:         u.cfg.Board1Pon5.OnuLastOfflineOID,
			OnuLastOfflineReasonOID:   u.cfg.Board1Pon5.OnuLastOfflineReasonOID,
			OnuGponOpticalDistanceOID: u.cfg.Board1Pon5.OnuGponOpticalDistanceOID,
		}
	case 6: // PON 6
		return &model.OltConfig{
			BaseOID:                   u.cfg.OltCfg.BaseOID1,
			OnuIDNameOID:              u.cfg.Board1Pon6.OnuIDNameOID,
			OnuTypeOID:                u.cfg.Board1Pon6.OnuTypeOID,
			OnuSerialNumberOID:        u.cfg.Board1Pon6.OnuSerialNumberOID,
			OnuRxPowerOID:             u.cfg.Board1Pon6.OnuRxPowerOID,
			OnuTxPowerOID:             u.cfg.Board1Pon6.OnuTxPowerOID,
			OnuStatusOID:              u.cfg.Board1Pon6.OnuStatusOID,
			OnuIPAddressOID:           u.cfg.Board1Pon6.OnuIPAddressOID,
			OnuDescriptionOID:         u.cfg.Board1Pon6.OnuDescriptionOID,
			OnuLastOnlineOID:          u.cfg.Board1Pon6.OnuLastOnlineOID,
			OnuLastOfflineOID:         u.cfg.Board1Pon6.OnuLastOfflineOID,
			OnuLastOfflineReasonOID:   u.cfg.Board1Pon6.OnuLastOfflineReasonOID,
			OnuGponOpticalDistanceOID: u.cfg.Board1Pon6.OnuGponOpticalDistanceOID,
		}
	case 7: // PON 7
		return &model.OltConfig{
			BaseOID:                   u.cfg.OltCfg.BaseOID1,
			OnuIDNameOID:              u.cfg.Board1Pon7.OnuIDNameOID,
			OnuTypeOID:                u.cfg.Board1Pon7.OnuTypeOID,
			OnuSerialNumberOID:        u.cfg.Board1Pon7.OnuSerialNumberOID,
			OnuRxPowerOID:             u.cfg.Board1Pon7.OnuRxPowerOID,
			OnuTxPowerOID:             u.cfg.Board1Pon7.OnuTxPowerOID,
			OnuStatusOID:              u.cfg.Board1Pon7.OnuStatusOID,
			OnuIPAddressOID:           u.cfg.Board1Pon7.OnuIPAddressOID,
			OnuDescriptionOID:         u.cfg.Board1Pon7.OnuDescriptionOID,
			OnuLastOnlineOID:          u.cfg.Board1Pon7.OnuLastOnlineOID,
			OnuLastOfflineOID:         u.cfg.Board1Pon7.OnuLastOfflineOID,
			OnuLastOfflineReasonOID:   u.cfg.Board1Pon7.OnuLastOfflineReasonOID,
			OnuGponOpticalDistanceOID: u.cfg.Board1Pon7.OnuGponOpticalDistanceOID,
		}
	case 8: // PON 8
		return &model.OltConfig{
			BaseOID:                   u.cfg.OltCfg.BaseOID1,
			OnuIDNameOID:              u.cfg.Board1Pon8.OnuIDNameOID,
			OnuTypeOID:                u.cfg.Board1Pon8.OnuTypeOID,
			OnuSerialNumberOID:        u.cfg.Board1Pon8.OnuSerialNumberOID,
			OnuRxPowerOID:             u.cfg.Board1Pon8.OnuRxPowerOID,
			OnuTxPowerOID:             u.cfg.Board1Pon8.OnuTxPowerOID,
			OnuStatusOID:              u.cfg.Board1Pon8.OnuStatusOID,
			OnuIPAddressOID:           u.cfg.Board1Pon8.OnuIPAddressOID,
			OnuDescriptionOID:         u.cfg.Board1Pon8.OnuDescriptionOID,
			OnuLastOnlineOID:          u.cfg.Board1Pon8.OnuLastOnlineOID,
			OnuLastOfflineOID:         u.cfg.Board1Pon8.OnuLastOfflineOID,
			OnuLastOfflineReasonOID:   u.cfg.Board1Pon8.OnuLastOfflineReasonOID,
			OnuGponOpticalDistanceOID: u.cfg.Board1Pon8.OnuGponOpticalDistanceOID,
		}
	default:
		log.Error().Msg("Invalid PON ID") // Log error message
		return nil
	}
}

// getBoard2Config is a function to get board 2 configuration
func (u *onuUsecase) getBoard2Config(ponID int) *model.OltConfig {
	// Define the configuration for Board 2
	switch ponID {
	case 1: // PON 1
		return &model.OltConfig{
			BaseOID:                   u.cfg.OltCfg.BaseOID1,
			OnuIDNameOID:              u.cfg.Board2Pon1.OnuIDNameOID,
			OnuTypeOID:                u.cfg.Board2Pon1.OnuTypeOID,
			OnuSerialNumberOID:        u.cfg.Board2Pon1.OnuSerialNumberOID,
			OnuRxPowerOID:             u.cfg.Board2Pon1.OnuRxPowerOID,
			OnuTxPowerOID:             u.cfg.Board2Pon1.OnuTxPowerOID,
			OnuStatusOID:              u.cfg.Board2Pon1.OnuStatusOID,
			OnuIPAddressOID:           u.cfg.Board2Pon1.OnuIPAddressOID,
			OnuDescriptionOID:         u.cfg.Board2Pon1.OnuDescriptionOID,
			OnuLastOnlineOID:          u.cfg.Board2Pon1.OnuLastOnlineOID,
			OnuLastOfflineOID:         u.cfg.Board2Pon1.OnuLastOfflineOID,
			OnuLastOfflineReasonOID:   u.cfg.Board2Pon1.OnuLastOfflineReasonOID,
			OnuGponOpticalDistanceOID: u.cfg.Board2Pon1.OnuGponOpticalDistanceOID,
		}
	case 2: // PON 2
		return &model.OltConfig{
			BaseOID:                   u.cfg.OltCfg.BaseOID1,
			OnuIDNameOID:              u.cfg.Board2Pon2.OnuIDNameOID,
			OnuTypeOID:                u.cfg.Board2Pon2.OnuTypeOID,
			OnuSerialNumberOID:        u.cfg.Board2Pon2.OnuSerialNumberOID,
			OnuRxPowerOID:             u.cfg.Board2Pon2.OnuRxPowerOID,
			OnuTxPowerOID:             u.cfg.Board2Pon2.OnuTxPowerOID,
			OnuStatusOID:              u.cfg.Board2Pon2.OnuStatusOID,
			OnuIPAddressOID:           u.cfg.Board2Pon2.OnuIPAddressOID,
			OnuDescriptionOID:         u.cfg.Board2Pon2.OnuDescriptionOID,
			OnuLastOnlineOID:          u.cfg.Board2Pon2.OnuLastOnlineOID,
			OnuLastOfflineOID:         u.cfg.Board2Pon2.OnuLastOfflineOID,
			OnuLastOfflineReasonOID:   u.cfg.Board2Pon2.OnuLastOfflineReasonOID,
			OnuGponOpticalDistanceOID: u.cfg.Board2Pon2.OnuGponOpticalDistanceOID,
		}
	case 3: // PON 3
		return &model.OltConfig{
			BaseOID:                   u.cfg.OltCfg.BaseOID1,
			OnuIDNameOID:              u.cfg.Board2Pon3.OnuIDNameOID,
			OnuTypeOID:                u.cfg.Board2Pon3.OnuTypeOID,
			OnuSerialNumberOID:        u.cfg.Board2Pon3.OnuSerialNumberOID,
			OnuRxPowerOID:             u.cfg.Board2Pon3.OnuRxPowerOID,
			OnuTxPowerOID:             u.cfg.Board2Pon3.OnuTxPowerOID,
			OnuStatusOID:              u.cfg.Board2Pon3.OnuStatusOID,
			OnuIPAddressOID:           u.cfg.Board2Pon3.OnuIPAddressOID,
			OnuDescriptionOID:         u.cfg.Board2Pon3.OnuDescriptionOID,
			OnuLastOnlineOID:          u.cfg.Board2Pon3.OnuLastOnlineOID,
			OnuLastOfflineOID:         u.cfg.Board2Pon3.OnuLastOfflineOID,
			OnuLastOfflineReasonOID:   u.cfg.Board2Pon3.OnuLastOfflineReasonOID,
			OnuGponOpticalDistanceOID: u.cfg.Board2Pon3.OnuGponOpticalDistanceOID,
		}
	case 4: // PON 4
		return &model.OltConfig{
			BaseOID:                   u.cfg.OltCfg.BaseOID1,
			OnuIDNameOID:              u.cfg.Board2Pon4.OnuIDNameOID,
			OnuTypeOID:                u.cfg.Board2Pon4.OnuTypeOID,
			OnuSerialNumberOID:        u.cfg.Board2Pon4.OnuSerialNumberOID,
			OnuRxPowerOID:             u.cfg.Board2Pon4.OnuRxPowerOID,
			OnuTxPowerOID:             u.cfg.Board2Pon4.OnuTxPowerOID,
			OnuStatusOID:              u.cfg.Board2Pon4.OnuStatusOID,
			OnuIPAddressOID:           u.cfg.Board2Pon4.OnuIPAddressOID,
			OnuDescriptionOID:         u.cfg.Board2Pon4.OnuDescriptionOID,
			OnuLastOnlineOID:          u.cfg.Board2Pon4.OnuLastOnlineOID,
			OnuLastOfflineOID:         u.cfg.Board2Pon4.OnuLastOfflineOID,
			OnuLastOfflineReasonOID:   u.cfg.Board2Pon4.OnuLastOfflineReasonOID,
			OnuGponOpticalDistanceOID: u.cfg.Board2Pon4.OnuGponOpticalDistanceOID,
		}
	case 5: // PON 5
		return &model.OltConfig{
			BaseOID:                   u.cfg.OltCfg.BaseOID1,
			OnuIDNameOID:              u.cfg.Board2Pon5.OnuIDNameOID,
			OnuTypeOID:                u.cfg.Board2Pon5.OnuTypeOID,
			OnuSerialNumberOID:        u.cfg.Board2Pon5.OnuSerialNumberOID,
			OnuRxPowerOID:             u.cfg.Board2Pon5.OnuRxPowerOID,
			OnuTxPowerOID:             u.cfg.Board2Pon5.OnuTxPowerOID,
			OnuStatusOID:              u.cfg.Board2Pon5.OnuStatusOID,
			OnuIPAddressOID:           u.cfg.Board2Pon5.OnuIPAddressOID,
			OnuDescriptionOID:         u.cfg.Board2Pon5.OnuDescriptionOID,
			OnuLastOnlineOID:          u.cfg.Board2Pon5.OnuLastOnlineOID,
			OnuLastOfflineOID:         u.cfg.Board2Pon5.OnuLastOfflineOID,
			OnuLastOfflineReasonOID:   u.cfg.Board2Pon5.OnuLastOfflineReasonOID,
			OnuGponOpticalDistanceOID: u.cfg.Board2Pon5.OnuGponOpticalDistanceOID,
		}
	case 6: // PON 6
		return &model.OltConfig{
			BaseOID:                   u.cfg.OltCfg.BaseOID1,
			OnuIDNameOID:              u.cfg.Board2Pon6.OnuIDNameOID,
			OnuTypeOID:                u.cfg.Board2Pon6.OnuTypeOID,
			OnuSerialNumberOID:        u.cfg.Board2Pon6.OnuSerialNumberOID,
			OnuRxPowerOID:             u.cfg.Board2Pon6.OnuRxPowerOID,
			OnuTxPowerOID:             u.cfg.Board2Pon6.OnuTxPowerOID,
			OnuStatusOID:              u.cfg.Board2Pon6.OnuStatusOID,
			OnuIPAddressOID:           u.cfg.Board2Pon6.OnuIPAddressOID,
			OnuDescriptionOID:         u.cfg.Board2Pon6.OnuDescriptionOID,
			OnuLastOnlineOID:          u.cfg.Board2Pon6.OnuLastOnlineOID,
			OnuLastOfflineOID:         u.cfg.Board2Pon6.OnuLastOfflineOID,
			OnuLastOfflineReasonOID:   u.cfg.Board2Pon6.OnuLastOfflineReasonOID,
			OnuGponOpticalDistanceOID: u.cfg.Board2Pon6.OnuGponOpticalDistanceOID,
		}
	case 7: // PON 7
		return &model.OltConfig{
			BaseOID:                   u.cfg.OltCfg.BaseOID1,
			OnuIDNameOID:              u.cfg.Board2Pon7.OnuIDNameOID,
			OnuTypeOID:                u.cfg.Board2Pon7.OnuTypeOID,
			OnuSerialNumberOID:        u.cfg.Board2Pon7.OnuSerialNumberOID,
			OnuRxPowerOID:             u.cfg.Board2Pon7.OnuRxPowerOID,
			OnuTxPowerOID:             u.cfg.Board2Pon7.OnuTxPowerOID,
			OnuStatusOID:              u.cfg.Board2Pon7.OnuStatusOID,
			OnuIPAddressOID:           u.cfg.Board2Pon7.OnuIPAddressOID,
			OnuDescriptionOID:         u.cfg.Board2Pon7.OnuDescriptionOID,
			OnuLastOnlineOID:          u.cfg.Board2Pon7.OnuLastOnlineOID,
			OnuLastOfflineOID:         u.cfg.Board2Pon7.OnuLastOfflineOID,
			OnuLastOfflineReasonOID:   u.cfg.Board2Pon7.OnuLastOfflineReasonOID,
			OnuGponOpticalDistanceOID: u.cfg.Board2Pon7.OnuGponOpticalDistanceOID,
		}
	case 8: // PON 8
		return &model.OltConfig{
			BaseOID:                   u.cfg.OltCfg.BaseOID1,
			OnuIDNameOID:              u.cfg.Board2Pon8.OnuIDNameOID,
			OnuTypeOID:                u.cfg.Board2Pon8.OnuTypeOID,
			OnuSerialNumberOID:        u.cfg.Board2Pon8.OnuSerialNumberOID,
			OnuRxPowerOID:             u.cfg.Board2Pon8.OnuRxPowerOID,
			OnuTxPowerOID:             u.cfg.Board2Pon8.OnuTxPowerOID,
			OnuStatusOID:              u.cfg.Board2Pon8.OnuStatusOID,
			OnuIPAddressOID:           u.cfg.Board2Pon8.OnuIPAddressOID,
			OnuDescriptionOID:         u.cfg.Board2Pon8.OnuDescriptionOID,
			OnuLastOnlineOID:          u.cfg.Board2Pon8.OnuLastOnlineOID,
			OnuLastOfflineOID:         u.cfg.Board2Pon8.OnuLastOfflineOID,
			OnuLastOfflineReasonOID:   u.cfg.Board2Pon8.OnuLastOfflineReasonOID,
			OnuGponOpticalDistanceOID: u.cfg.Board2Pon8.OnuGponOpticalDistanceOID,
		}
	default:
		log.Error().Msg("Invalid PON ID") // Log error message
		return nil
	}
}

func (u *onuUsecase) GetByBoardIDAndPonID(ctx context.Context, boardID, ponID int) ([]model.ONUInfoPerBoard, error) {

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

	err = u.snmpRepository.Walk(oltConfig.BaseOID+oltConfig.OnuIDNameOID, func(pdu gosnmp.SnmpPDU) error {
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
		onuType, err := u.getONUType(oltConfig.OnuTypeOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.OnuType = onuType // Set ONU Type to ONU onuInfo struct OnuType field
		}

		// Get ONU Serial Number based on ONU ID and ONU Serial Number OID and store it to ONU onuInfo struct
		onuSerialNumber, err := u.getSerialNumber(oltConfig.OnuSerialNumberOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.SerialNumber = onuSerialNumber // Set ONU Serial Number to ONU onuInfo struct SerialNumber field
		}

		// Get ONU RX Power based on ONU ID and ONU RX Power OID and store it to ONU onuInfo struct
		onuRXPower, err := u.getRxPower(oltConfig.OnuRxPowerOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.RXPower = onuRXPower // Set ONU RX Power to ONU onuInfo struct RXPower field
		}

		// Get ONU Status based on ONU ID and ONU Status OID and store it to ONU onuInfo struct
		onuStatus, err := u.getStatus(oltConfig.OnuStatusOID, strconv.Itoa(onuInfo.ID))
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

func (u *onuUsecase) GetByBoardIDPonIDAndOnuID(boardID, ponID, onuID int) (
	model.ONUCustomerInfo, error,
) {

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
		onuType, err := u.getONUType(oltConfig.OnuTypeOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.OnuType = onuType // Set ONU Type from SNMP Walk result if no error to onuInfo variable (ONU Type)
		}

		// Get Data ONU Serial Number from SNMP Walk using getSerialNumber method
		onuSerialNumber, err := u.getSerialNumber(oltConfig.OnuSerialNumberOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.SerialNumber = onuSerialNumber // Set ONU Serial Number from SNMP Walk result to onuInfo variable (ONU Serial Number)
		}

		// Get Data ONU RX Power from SNMP Walk using getRxPower method
		onuRXPower, err := u.getRxPower(oltConfig.OnuRxPowerOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.RXPower = onuRXPower // Set ONU RX Power from SNMP Walk result to onuInfo variable (ONU RX Power)
		}

		// Get Data ONU TX Power from SNMP Walk using getTxPower method
		onuTXPower, err := u.getTxPower(oltConfig.OnuTxPowerOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.TXPower = onuTXPower // Set ONU TX Power from SNMP Walk result to onuInfo variable (ONU TX Power)
		}

		// Get Data ONU Status from SNMP Walk using getStatus method
		onuStatus, err := u.getStatus(oltConfig.OnuStatusOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.Status = onuStatus // Set ONU Status from SNMP Walk result to onuInfo variable (ONU Status)
		}

		// Get Data ONU IP Address from SNMP Walk using getIPAddress method
		onuIPAddress, err := u.getIPAddress(oltConfig.OnuIPAddressOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.IPAddress = onuIPAddress // Set ONU IP Address from SNMP Walk result to onuInfo variable (ONU IP Address)
		}

		// Get Data ONU Description from SNMP Walk using getDescription method
		onuDescription, err := u.getDescription(oltConfig.OnuDescriptionOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.Description = onuDescription // Set ONU Description from SNMP Walk result to onuInfo variable (ONU Description)
		}

		// Get Data ONU Last Online from SNMP Walk using getLastOnline method
		onuLastOnline, err := u.getLastOnline(oltConfig.OnuLastOnlineOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.LastOnline = onuLastOnline // Set ONU Last Online from SNMP Walk result to onuInfo variable (ONU Last Online)
		}

		// Get Data ONU Last Offline from SNMP Walk using getLastOffline method
		onuLastOffline, err := u.getLastOffline(oltConfig.OnuLastOfflineOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.LastOffline = onuLastOffline // Set ONU Last Offline from SNMP Walk result to onuInfo variable (ONU Last Offline)
		}

		// Get Data Uptime Duration from getUptimeDuration method
		onuUptimeDuration, err := u.getUptimeDuration(onuLastOnline)
		if err == nil {
			onuInfo.Uptime = onuUptimeDuration // Set ONU Uptime Duration from SNMP Walk result to onuInfo variable (ONU Uptime Duration)
		}

		// Get Data Last Down Duration from getLastDownDuration method
		onuLastDownDuration, err := u.getLastDownDuration(onuLastOffline, onuLastOnline)
		if err == nil {
			onuInfo.LastDownTimeDuration = onuLastDownDuration // Set ONU Last Down Duration from SNMP Walk result to onuInfo variable (ONU Last Down Duration)
		}

		// Get Data ONU Last Offline Reason from SNMP Walk using getLastOfflineReason method
		onuLastOfflineReason, err := u.getLastOfflineReason(oltConfig.OnuLastOfflineReasonOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.LastOfflineReason = onuLastOfflineReason // Set ONU Last Offline Reason from SNMP Walk result to onuInfo variable (ONU Last Offline Reason)
		}

		// Get Data ONU GPON Optical Distance from SNMP Walk using getGponOpticalDistance method
		onuGponOpticalDistance, err := u.getOnuGponOpticalDistance(oltConfig.OnuGponOpticalDistanceOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.GponOpticalDistance = onuGponOpticalDistance // Set ONU GPON Optical Distance from SNMP Walk result to onuInfo variable (ONU GPON Optical Distance)
		}

		onuInformationList = onuInfo // Append onuInfo variable to onuInformationList slice
	}

	return onuInformationList, nil
}

func (u *onuUsecase) GetEmptyOnuID(ctx context.Context, boardID, ponID int) ([]model.OnuID, error) {

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
	err = u.snmpRepository.Walk(snmpOID, func(pdu gosnmp.SnmpPDU) error {
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

func (u *onuUsecase) GetOnuIDAndSerialNumber(boardID, ponID int) (
	[]model.OnuSerialNumber, error,
) {

	// Get OLT config based on Board ID and PON ID
	oltConfig, err := u.getOltConfig(boardID, ponID)
	if err != nil {
		log.Error().Msg("Failed to get OLT Config: " + err.Error()) // Log error message to logger
		return nil, err                                             // Return error if error is not nil
	}

	// Perform SNMP Walk to get ONU ID
	snmpOID := oltConfig.BaseOID + oltConfig.OnuIDNameOID // SNMP OID variable
	onuIDList := make([]model.OnuID, 0)                   // Create a slice of ONU ID

	log.Info().Msg("Get ONU ID with SNMP Walk from Board ID: " + strconv.Itoa(
		boardID) + " and PON ID: " + strconv.Itoa(ponID)) // Log info message to logger

	// Perform SNMP BulkWalk to get ONU ID and Name using snmpRepository BulkWalk method with timeout context parameter
	err = u.snmpRepository.Walk(snmpOID, func(pdu gosnmp.SnmpPDU) error {
		idOnuID := utils.ExtractIDOnuID(pdu.Name) // Extract ONU ID from SNMP PDU Name
		// Append ONU information to the onuIDList
		onuIDList = append(onuIDList, model.OnuID{
			Board: boardID, // Set Board ID to ONU onuInfo struct Board field
			PON:   ponID,   // Set PON ID to ONU onuInfo  struct PON field
			ID:    idOnuID, // Set ONU ID (extracted from SNMP PDU) to onuInfo variable (ONU ID)
		})

		return nil
	})

	if err != nil {
		log.Error().Msg("Failed to perform SNMP Walk get ONU ID: " + err.Error()) // Log error message to logger
		return nil, err
	}

	// Create a slice of ONU Serial Number
	var onuSerialNumberList []model.OnuSerialNumber

	// Loop through onuIDList to get ONU Serial Number
	for _, onuInfo := range onuIDList {
		// Get Data ONU Serial Number from SNMP Walk using getSerialNumber method
		onuSerialNumber, err := u.getSerialNumber(oltConfig.OnuSerialNumberOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuSerialNumberList = append(onuSerialNumberList, model.OnuSerialNumber{
				Board:        boardID, // Set Board ID to ONU onuInfo struct Board field
				PON:          ponID,   // Set PON ID to ONU onuInfo  struct PON field
				ID:           onuInfo.ID,
				SerialNumber: onuSerialNumber, // Set ONU Serial Number from SNMP Walk result to onuInfo variable (ONU Serial Number)
			})
		}
	}

	// Sort ONU Serial Number list based on ONU ID ascending
	sort.Slice(onuSerialNumberList, func(i, j int) bool {
		return onuSerialNumberList[i].ID < onuSerialNumberList[j].ID
	})

	return onuSerialNumberList, nil // Return ONU Serial Number list and nil error

}

func (u *onuUsecase) UpdateEmptyOnuID(ctx context.Context, boardID, ponID int) error {

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
	err = u.snmpRepository.Walk(snmpOID, func(pdu gosnmp.SnmpPDU) error {
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
	boardID, ponID, pageIndex, pageSize int,
) ([]model.ONUInfoPerBoard, int) {

	// Get OLT config based on Board ID and PON ID
	oltConfig, err := u.getOltConfig(boardID, ponID)
	if err != nil {
		return nil, 0
	}

	// SNMP OID variable
	snmpOID := oltConfig.BaseOID + oltConfig.OnuIDNameOID

	var onlyOnuIDList []model.OnuOnlyID
	var count int

	// If data not exist in Redis, then get data from SNMP
	if len(onlyOnuIDList) == 0 {
		err := u.snmpRepository.Walk(snmpOID, func(pdu gosnmp.SnmpPDU) error {
			onlyOnuIDList = append(onlyOnuIDList, model.OnuOnlyID{
				ID: utils.ExtractIDOnuID(pdu.Name),
			})
			return nil
		})

		if err != nil {
			return nil, 0
		}
	} else {
		//// If data exist in Redis, then get data from Redis
		//onlyOnuIDList, err = u.redisRepository.GetOnuOnlyID(ctx, "board_"+strconv.Itoa(boardID)+"_pon_"+strconv.Itoa(ponID)+"_only_onu_id")
		//if err != nil {
		//	return nil, 0
		//}
		log.Error().Msg("Failed to get data from Redis")
	}

	count = len(onlyOnuIDList)

	// Calculate the index of the first item to be retrieved
	//startIndex := pageIndex * pageSize
	startIndex := (pageIndex - 1) * pageSize

	// Calculate the index of the last item to be retrieved
	endIndex := startIndex + pageSize

	// If the index of the last item to be retrieved is greater than the number of items, set it to the number of items
	if endIndex > len(onlyOnuIDList) {
		endIndex = len(onlyOnuIDList)
	}

	// Get ONU IDs to be displayed based on the index of the first and last items from the onlyOnuIDList data
	onlyOnuIDList = onlyOnuIDList[startIndex:endIndex]

	var onuInformationList []model.ONUInfoPerBoard

	// Loop through onlyOnuIDList to get ONU information based on ONU ID
	for _, onuID := range onlyOnuIDList {
		onuInfo := model.ONUInfoPerBoard{
			Board: boardID,  // Set Board ID to ONUInfo struct Board field
			PON:   ponID,    // Set PON ID to ONUInfo struct PON field
			ID:    onuID.ID, // Set ONU ID to ONUInfo struct ID field
		}

		// Get Name base on ONU ID and ONU Name OID and store it to ONU onuInfo struct
		onuName, err := u.getName(oltConfig.OnuIDNameOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.Name = onuName // Set ONU Name to ONU onuInfo struct Name field
		}

		// Get ONU Type based on ONU ID and ONU Type OID and store it to ONU onuInfo struct
		onuType, err := u.getONUType(oltConfig.OnuTypeOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.OnuType = onuType // Set ONU Type to ONU onuInfo struct OnuType field
		}

		// Get ONU Serial Number based on ONU ID and ONU Serial Number OID and store it to ONU onuInfo struct
		onuSerialNumber, err := u.getSerialNumber(oltConfig.OnuSerialNumberOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.SerialNumber = onuSerialNumber // Set ONU Serial Number to ONU onuInfo struct SerialNumber field
		}

		// Get ONU RX Power based on ONU ID and ONU RX Power OID and store it to ONU onuInfo struct
		onuRXPower, err := u.getRxPower(oltConfig.OnuRxPowerOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.RXPower = onuRXPower // Set ONU RX Power to ONU onuInfo struct RXPower field
		}

		// Get ONU Status based on ONU ID and ONU Status OID and store it to ONU onuInfo struct
		onuStatus, err := u.getStatus(oltConfig.OnuStatusOID, strconv.Itoa(onuInfo.ID))
		if err == nil {
			onuInfo.Status = onuStatus // Set ONU Status to ONU onuInfo struct Status field
		}

		// Append ONU information to the onuInformationList
		onuInformationList = append(onuInformationList, onuInfo)
	}

	//Sort ONU information list based on ONU ID ascending
	sort.Slice(onuInformationList, func(i, j int) bool {
		return onuInformationList[i].ID < onuInformationList[j].ID
	})

	// Return the page data along with the total number of available data
	return onuInformationList, count
}

func (u *onuUsecase) getName(OnuIDNameOID, onuID string) (string, error) {

	var onuName string // Variable to store ONU Name

	baseOID := u.cfg.OltCfg.BaseOID1 // Base OID variable get from config

	// Perform SNMP Get to get ONU Name using snmpRepository Get method with timeout context parameter
	oids := []string{baseOID + OnuIDNameOID + "." + onuID}
	result, err := u.snmpRepository.Get(oids)
	//result, err := u.snmpRepository.Get(oids)
	if err != nil {
		log.Error().Msg("Failed to perform SNMP Get for Name: " + err.Error()) // Log error message to logger
		return "", errors.New("failed to perform SNMP Get")                    // Return error
	}

	// Check if the result contains the expected OID
	if len(result.Variables) > 0 {
		onuName = utils.ExtractName(result.Variables[0].Value) // Extract ONU Name from the result
	} else {
		log.Error().Msg("Failed to get ONU Name: No variables in the response")
		return "", errors.New("no variables in the response")
	}

	return onuName, nil // Return ONU Name
}

func (u *onuUsecase) getONUType(OnuTypeOID, onuID string) (string, error) {

	var onuType string // Variable to store ONU Type

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

func (u *onuUsecase) getSerialNumber(OnuSerialNumberOID, onuID string) (string, error) {

	var onuSerialNumber string // Variable to store ONU Serial Number

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

func (u *onuUsecase) getRxPower(OnuRxPowerOID, onuID string) (string, error) {

	var onuRxPower string // Variable to store ONU RX Power

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

func (u *onuUsecase) getTxPower(OnuTxPowerOID, onuID string) (string, error) {

	var onuTxPower string // Variable to store ONU TX Power

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

func (u *onuUsecase) getStatus(OnuStatusOID, onuID string) (string, error) {

	var onuStatus string // Variable to store ONU Status

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

func (u *onuUsecase) getIPAddress(OnuIPAddressOID, onuID string) (string, error) {

	var onuIPAddress string // Variable to store ONU IP Address

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

func (u *onuUsecase) getDescription(OnuDescriptionOID, onuID string) (string, error) {

	var onuDescription string // Variable to store ONU Description

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

func (u *onuUsecase) getLastOnline(OnuLastOnlineOID, onuID string) (string, error) {

	var onuLastOnline string // Variable to store ONU Last Online

	baseOID := u.cfg.OltCfg.BaseOID1 // Base OID variable get from config

	// Perform SNMP Get to get ONU Last Online using snmpRepository Get method with timeout context parameter
	oids := []string{baseOID + OnuLastOnlineOID + "." + onuID}
	result, err := u.snmpRepository.Get(oids)
	if err != nil {
		log.Error().Msg("Failed to perform SNMP Get for last online: " + err.Error()) // Log error message to logger
		return "", errors.New("failed to perform SNMP Get")                           // Return error
	}

	// Check if the result contains the expected OID
	if len(result.Variables) > 0 {
		value := result.Variables[0].Value.([]byte) // Assuming the value is returned as a byte array (Octet String)

		// Convert the Octet String to a DateTime
		onuLastOnline, err = utils.ConvertByteArrayToDateTime(value)

		if err != nil {
			log.Error().Msg("Failed to convert byte array to DateTime: " + err.Error())
			return "", err
		}

	} else {
		log.Error().Msg("Failed to get ONU Last Online: No variables in the response")
		return "", errors.New("no variables in the response")
	}

	return onuLastOnline, nil // Return ONU Last Online as a string
}

func (u *onuUsecase) getLastOffline(OnuLastOfflineOID, onuID string) (string, error) {

	var onuLastOffline string // Variable to store ONU Last Offline

	baseOID := u.cfg.OltCfg.BaseOID1 // Base OID variable get from config

	// Perform SNMP Get to get ONU Last Offline using snmpRepository Get method with timeout context parameter
	oids := []string{baseOID + OnuLastOfflineOID + "." + onuID}
	result, err := u.snmpRepository.Get(oids)
	if err != nil {
		log.Error().Msg("Failed to perform SNMP Get for last offline: " + err.Error()) // Log error message to logger
		return "", errors.New("failed to perform SNMP Get")                            // Return error
	}

	// Check if the result contains the expected OID
	if len(result.Variables) > 0 {
		value := result.Variables[0].Value.([]byte) // Assuming the value is returned as a byte array (Octet String)

		// Convert the Octet String to a DateTime
		onuLastOffline, err = utils.ConvertByteArrayToDateTime(value)

		if err != nil {
			log.Error().Msg("Failed to convert byte array to DateTime: " + err.Error())
			return "", err
		}

	} else {
		log.Error().Msg("Failed to get ONU Last Offline: No variables in the response")
		return "", errors.New("no variables in the response")
	}

	return onuLastOffline, nil // Return ONU Last Offline as a string
}

func (u *onuUsecase) getLastOfflineReason(OnuLastOfflineReasonOID, onuID string) (string, error) {

	var onuLastOfflineReason string // Variable to store ONU Last Offline Reason

	baseOID := u.cfg.OltCfg.BaseOID1 // Base OID variable get from config

	// Perform SNMP Get to get ONU Last Offline Reason using snmpRepository Get method with timeout context parameter
	oids := []string{baseOID + OnuLastOfflineReasonOID + "." + onuID}
	result, err := u.snmpRepository.Get(oids)
	if err != nil {
		log.Error().Msg("Failed to perform SNMP Get for last offline reason: " + err.Error()) // Log error message to logger
		return "", errors.New("failed to perform SNMP Get")                                   // Return error
	}

	// Check if the result contains the expected OID
	if len(result.Variables) > 0 {
		onuLastOfflineReason = utils.ExtractLastOfflineReason(result.Variables[0].Value) // Extract ONU Last Offline Reason from the result
	} else {
		log.Error().Msg("Failed to get ONU Last Offline Reason: No variables in the response")
		return "", errors.New("no variables in the response")
	}

	return onuLastOfflineReason, nil // Return ONU Last Offline Reason
}

func (u *onuUsecase) getOnuGponOpticalDistance(OnuGponOpticalDistanceOID, onuID string) (string, error) {

	var onuGponOpticalDistance string // Variable to store ONU GPON Optical Distance

	baseOID := u.cfg.OltCfg.BaseOID1 // Base OID variable get from config

	// Perform SNMP Get to get ONU GPON Optical Distance using snmpRepository Get method with timeout context parameter
	oids := []string{baseOID + OnuGponOpticalDistanceOID + "." + onuID}
	fmt.Println(oids)
	result, err := u.snmpRepository.Get(oids)
	if err != nil {
		log.Error().Msg("Failed to perform SNMP Get for GPON Optical Distance: " + err.Error()) // Log error message to logger
		return "", errors.New("failed to perform SNMP Get")                                     // Return error
	}

	// Check if the result contains the expected OID
	if len(result.Variables) > 0 {
		onuGponOpticalDistance = utils.ExtractGponOpticalDistance(result.Variables[0].Value) // Extract ONU GPON Optical Distance from the result
	} else {
		log.Error().Msg("Failed to get ONU GPON Optical Distance: No variables in the response")
		return "", errors.New("no variables in the response")
	}

	return onuGponOpticalDistance, nil // Return ONU GPON Optical Distance
}

func (u *onuUsecase) getUptimeDuration(lastOnline string) (string, error) {

	// Get current time in UTC
	currentTime := time.Now()

	// Convert last online time to UTC
	lastOnlineTime, err := time.Parse("2006-01-02 15:04:05", lastOnline)
	if err != nil {
		log.Error().Msg("Failed to parse last online time: " + err.Error())
		return "", err
	}

	fmt.Println("This is time now: ", currentTime)
	fmt.Println("This is last online time: ", lastOnlineTime)
	fmt.Println("This is duration: ", currentTime.Sub(lastOnlineTime))

	// Calculate the duration between the last online time and the current time
	duration := currentTime.Sub(lastOnlineTime) + time.Hour*7

	// Convert the duration to a string
	uptimeDuration := utils.ConvertDurationToString(duration)

	return uptimeDuration, nil

}

// Last Down Duration
func (u *onuUsecase) getLastDownDuration(lastOffline, lastOnline string) (string, error) {

	// Convert last offline time to time
	lastOfflineTime, err := time.Parse("2006-01-02 15:04:05", lastOffline)
	if err != nil {
		log.Error().Msg("Failed to parse last offline time: " + err.Error())
		return "", err
	}

	// Convert last online time to time
	lastOnlineTime, err := time.Parse("2006-01-02 15:04:05", lastOnline)
	if err != nil {
		log.Error().Msg("Failed to parse last online time: " + err.Error())
		return "", err
	}

	// Calculate the duration between the last offline time and the last online time
	duration := lastOnlineTime.Sub(lastOfflineTime)

	// Convert the duration to a string
	lastDownDuration := utils.ConvertDurationToString(duration)

	return lastDownDuration, nil
}
