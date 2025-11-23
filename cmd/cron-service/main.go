package main

import (
	"context"
	"log"

	"github.com/Mariia230800/redis-data-race-demo/internal/app"
)

func main() {

	ctx := context.Background()

	err := app.Run(ctx)
	if err != nil {
		log.Fatalf("cron-service start failed: %v", err)
	}

}
