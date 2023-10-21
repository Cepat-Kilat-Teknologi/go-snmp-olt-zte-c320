package graceful

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Shutdown(ctx context.Context, server *http.Server) error {
	ch := make(chan error, 1)

	go func() {
		err := server.ListenAndServe()
		if err != nil && !errors.Is(http.ErrServerClosed, err) {
			ch <- fmt.Errorf("failed to start server: %v", err)
		}
		close(ch)
	}()

	// Create a channel to capture OS signals (e.g., SIGINT or SIGTERM).
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		if err := server.Shutdown(timeoutCtx); err != nil {
			log.Printf("Failed to gracefully shut down the server: %v", err)
		}
	case sig := <-signalCh:
		log.Printf("Received signal: %v. Shutting down gracefully...", sig)

		// Inisialisasi konteks dengan timeout
		shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Printf("Failed to gracefully shut down the server: %v", err)
		}
	}

	return nil
}
