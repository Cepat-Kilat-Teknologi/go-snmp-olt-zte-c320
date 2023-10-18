package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/sumitroajiprabowo/go-snmp-olt-c320/config"
	"github.com/sumitroajiprabowo/go-snmp-olt-c320/internal/handler"
	snmp2 "github.com/sumitroajiprabowo/go-snmp-olt-c320/internal/repository/snmp"
	"github.com/sumitroajiprabowo/go-snmp-olt-c320/internal/usecase"
	"github.com/sumitroajiprabowo/go-snmp-olt-c320/pkg/snmp"
	"github.com/sumitroajiprabowo/go-snmp-olt-c320/pkg/utils"
	"log"
	"net/http"
	"os"
	"time"
)

type App struct {
	router http.Handler
}

func New() *App {
	return &App{}
}

func (a *App) Start(ctx context.Context) error {
	configPath := utils.GetConfigPath(os.Getenv("config"))
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	snmpConn, err := snmp.SetupSnmpConnection(cfg)
	if err != nil {
		return fmt.Errorf("failed to set up SNMP connection: %w", err)
	}
	defer func() {
		if err := snmpConn.Conn.Close(); err != nil {
			log.Printf("Failed to close SNMP connection: %v", err)
		}
	}()

	// Initialize repository
	snmpRepo := snmp2.NewPonRepository(snmpConn)

	// Initialize usecase
	onuUsecase := usecase.NewOnuUsecase(snmpRepo, cfg)

	// Initialize handler
	onuHandler := handler.NewOnuHandler(onuUsecase)

	// Initialize router
	a.router = loadRoutes(onuHandler)

	fmt.Printf("Starting server at %s:%d\n", cfg.ServerCfg.Host, cfg.ServerCfg.Port)

	addr := fmt.Sprintf("%s:%d", cfg.ServerCfg.Host, cfg.ServerCfg.Port)
	server := &http.Server{
		Addr:    addr,
		Handler: a.router,
	}

	ch := make(chan error, 1)

	go func() {
		err = server.ListenAndServe()
		if err != nil && !errors.Is(http.ErrServerClosed, err) {
			ch <- fmt.Errorf("failed to start server: %v", err)
		}
		close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		if err := server.Shutdown(timeoutCtx); err != nil {
			log.Printf("Failed to gracefully shut down the server: %v", err)
		}
	}

	return nil
}
