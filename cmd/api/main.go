package main

import (
	"context"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/app"
	"github.com/rs/zerolog/log"
)

func main() {

	// Initialize application
	server := app.New()                                     // Create a new instance of application
	ctx, cancel := context.WithCancel(context.Background()) // Create a new context with cancel function
	defer cancel()                                          // Cancel context when main function is finished

	// Start application server in a goroutine
	go func() {
		err := server.Start(ctx) // Start application server
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to start server") // Log error message
			cancel()                                           // Cancel context if error occurred
		}
	}()

	// Handle OS signals or other triggers for graceful shutdown if needed.
	// For example, you can use a signal package like os/signal to capture
	// and handle interrupt signals (e.g., Ctrl+C).

	// Wait for a signal to gracefully stop the application.
	select {
	case <-ctx.Done(): // Application was gracefully stopped or an error occurred
		log.Error().Err(ctx.Err()).Msg("Application was gracefully stopped or an error occurred")
	}
}
