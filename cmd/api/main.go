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
	defer cancel()                                          // Cancel context when the main function is finished

	// Start application server in a goroutine
	go func() {
		err := server.Start(ctx) // Start the application server
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to start server") // Log error message
			cancel()                                           // Cancel context if an error occurred
		}
	}()

	// Create a channel to wait for a signal to stop the application
	stopSignal := make(chan struct{})

	// You can replace the select statement with a simple channel receive
	<-stopSignal

	// Log that the application is stopping
	log.Info().Msg("Application is stopping")
}
