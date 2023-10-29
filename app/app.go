package app

import (
	"context"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/config"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/internal/handler"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/internal/repository"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/internal/usecase"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/internal/utils"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/pkg/graceful"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/pkg/redis"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/pkg/snmp"
	rds "github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

type App struct {
	router http.Handler
}

func New() *App {
	return &App{}
}

func (a *App) Start(ctx context.Context) error {

	// Get config path from APP_ENV environment variable
	configPath := utils.GetConfigPath(os.Getenv("APP_ENV"))

	// Load configuration file from config path
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Error().Err(err).Msg("Failed to load config")
	}

	// Initialize Redis client
	redisClient := redis.NewRedisClient(cfg)

	// Check Redis connection
	err = redisClient.Ping(ctx).Err()
	if err != nil {
		log.Error().Err(err).Msg("Failed to ping Redis server")
	} else {
		log.Info().Msg("Redis server successfully connected")
	}

	// Close Redis client
	defer func(redisClient *rds.Client) {
		err := redisClient.Close()
		if err != nil {
			log.Error().Err(err).Msg("Failed to close Redis client")
		}
	}(redisClient)

	// Initialize SNMP connection
	snmpConn, err := snmp.SetupSnmpConnection(ctx, cfg)
	if err != nil {
		log.Error().Err(err).Msg("Failed to setup SNMP connection")
	}

	// Check SNMP connection
	/*
		if SNMP Connection with wrong credentials in SNMP v3, return error is nil
		if SNMP Connection with wrong Port in SNMP v2 v2c, return error is nil
		if SNMP Connection with wrong community v2 v2c, return error is nil

		Connect creates and opens a socket. Because UDP is a connectionless protocol,
		you won't know if the remote host is responding until you send packets.
		Neither will you know if the host is regularly disappearing and reappearing.
	*/

	if snmpConn.Connect() != nil {
		log.Error().Err(err).Msg("Failed to connect to SNMP server")
	} else {
		log.Info().Msg("SNMP server successfully connected")
	}

	// Close SNMP connection after application shutdown
	defer func() {
		if err := snmpConn.Conn.Close(); err != nil {
			log.Error().Err(err).Msg("Failed to close SNMP connection")
		}
	}()

	// Initialize repository
	snmpRepo := repository.NewPonRepository(snmpConn)
	redisRepo := repository.NewOnuRedisRepo(redisClient)

	// Initialize usecase
	onuUsecase := usecase.NewOnuUsecase(snmpRepo, redisRepo, cfg)

	// Initialize handler
	onuHandler := handler.NewOnuHandler(onuUsecase)

	// Initialize router
	a.router = loadRoutes(onuHandler)

	// Start server
	addr := "8081"
	server := &http.Server{
		Addr:    ":" + addr,
		Handler: a.router,
	}

	// Start server at given address
	log.Info().Msgf("Application started at %s", addr)

	// Graceful shutdown
	return graceful.Shutdown(ctx, server)
}
