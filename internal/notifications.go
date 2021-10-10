package internal

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"time"
)

func broadcastNotifications(app *App) {
	for {
		log.Debug("searching for notifications...")
		ctx := context.Background()
		result, err := app.redisConn.BRPop(ctx, 2*time.Second, "notifications:queue").Result()
		if err == nil && result != nil {
			broadcast(result[1])
		}
	}
}
