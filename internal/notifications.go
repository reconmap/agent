package internal

import (
	"context"
	"time"
)

func broadcastNotifications(app *App) {
	for {
		var ctx = context.Background()
		result, err := app.redisConn.BRPop(ctx, 2*time.Second, "notifications:queue").Result()
		if err == nil && result != nil {
			broadcast(result[1])
		}
	}
}
