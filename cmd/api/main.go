package main

import (
	"context"
	"github.com/sumitroajiprabowo/go-snmp-olt-c320/app"
	"log"
)

func main() {
	server := app.New()
	err := server.Start(context.TODO())
	if err != nil {
		log.Println("Failed to start app:", err)
	}
}
