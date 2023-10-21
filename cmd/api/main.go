package main

import (
	"context"
	"fmt"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/app"
	"log"
)

func main() {
	server := app.New()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		err := server.Start(ctx)
		if err != nil {
			log.Fatalf("Failed to start app: %v", err)
		}
	}()

	// Handle OS signals or other triggers for graceful shutdown if needed.
	// For example, you can use a signal package like os/signal to capture
	// and handle interrupt signals (e.g., Ctrl+C).

	// Wait for a signal to gracefully stop the application.
	select {
	case <-ctx.Done():
		// Application was gracefully stopped
		fmt.Println("Application has been gracefully stopped.")
	}
}
