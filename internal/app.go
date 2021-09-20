package internal

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type App struct {
	redisConn *redis.Client
	muxRouter *mux.Router
}

func (app *App) Initialise() {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	app.redisConn = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	app.muxRouter = mux.NewRouter()
	app.muxRouter.HandleFunc("/term", handleWebsocket)
	app.muxRouter.HandleFunc("/notifications", handleNotifications)

	log.SetLevel(log.DebugLevel)
}

func (app *App) Run() {
	log.Info("Reconmap agent")
	log.Warn("Warning, this is an experimental function that has not been secured")

	var listen = flag.String("listen", "127.0.0.1:2020", "Host:port to listen on")
	flag.Parse()

	go broadcastNotifications(app)

	if err := http.ListenAndServe(*listen, app.muxRouter); err != nil {
		log.WithError(err).Fatal("Something went wrong with the webserver")
	}
}
