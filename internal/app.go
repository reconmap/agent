package internal

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// App contains properties needed for agent
// to connect to redis and http router.
type App struct {
	redisConn *redis.Client
	muxRouter *mux.Router
}

// NewApp returns a App struct that has intialized a redis client and http router.
func NewApp() App {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/term", handleWebsocket)
	muxRouter.HandleFunc("/notifications", handleNotifications)

	return App{
		muxRouter: muxRouter,
	}
}

func (app *App) connectRedis() *error {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	ctx := context.Background()

	redisConn := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	if _, err := redisConn.Ping(ctx).Result(); err != nil {
		return &err
	}

	app.redisConn = redisConn

	return nil
}

// Run starts the agent.
func (app *App) Run() *error {
	log.Info("Reconmap agent")

	listen := flag.String("listen", ":5520", "Host:port to listen on")
	flag.Parse()

	err := app.connectRedis()
	if err != nil {
		errorFormatted := fmt.Errorf("Unable to connect to redis (%v)", *err)
		return &errorFormatted
	}

	go broadcastNotifications(app)

	if err := http.ListenAndServe(*listen, app.muxRouter); err != nil {
		log.WithError(err).Fatal("Something went wrong with the webserver")
	}

	return nil
}
