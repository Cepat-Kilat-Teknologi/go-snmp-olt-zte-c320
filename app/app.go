package app

import (
	"context"
	"fmt"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/config"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/internal/handler"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/internal/repository"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/internal/usecase"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/internal/utils"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/pkg/graceful"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/pkg/redis"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/pkg/snmp"
	rds "github.com/redis/go-redis/v9"
	"log"
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

	// Get config path
	configPath := utils.GetConfigPath(os.Getenv("config"))

	// Load config
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// Initialize Redis client
	redisClient := redis.NewRedisClient(cfg)

	// Close Redis client
	defer func(redisClient *rds.Client) {
		err := redisClient.Close()
		if err != nil {
			log.Printf("Failed to close Redis client: %v", err)
		}
	}(redisClient)

	// Initialize SNMP connection
	snmpConn, err := snmp.SetupSnmpConnection(cfg)
	if err != nil {
		return fmt.Errorf("failed to set up SNMP connection: %w", err)
	}

	// Close SNMP connection
	defer func() {
		if err := snmpConn.Conn.Close(); err != nil {
			log.Printf("Failed to close SNMP connection: %v", err)
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
	log.Printf("Starting server at %s", addr)

	// Graceful shutdown
	return graceful.Shutdown(ctx, server)
}
