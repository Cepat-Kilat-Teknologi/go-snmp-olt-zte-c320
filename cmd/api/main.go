package main

import (
	"context"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/app"
	"github.com/rs/zerolog/log"
)

func main() {
	server := app.New()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		err := server.Start(ctx)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to start server")
			cancel()
		}
	}()

	// Handle OS signals or other triggers for graceful shutdown if needed.
	// For example, you can use a signal package like os/signal to capture
	// and handle interrupt signals (e.g., Ctrl+C).

	// Wait for a signal to gracefully stop the application.
	select {
	case <-ctx.Done():
		// Application was gracefully stopped or an error occurred
		log.Error().Err(ctx.Err()).Msg("Application was gracefully stopped or an error occurred")
	}
}
