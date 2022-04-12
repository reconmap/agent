package internal

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

const (
	redisTimeout = 2 * time.Second
)

func broadcastNotifications(app *App) {
	for {
		log.Debug("searching for notifications...")
		ctx := context.Background()
		result, err := app.redisConn.BRPop(ctx, redisTimeout, "notifications:queue").Result()
		if err != redis.Nil {
			log.Error(err)
		} else if result != nil {
			broadcast(result[1])
		}
	}
}
