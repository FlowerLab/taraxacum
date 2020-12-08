package main

import (
	"github.com/FlowerLab/blackdatura"
	"github.com/FlowerLab/taraxacum"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var log *zap.Logger

func init() {
	taraxacum.LoadConf()
	blackdatura.Init("debug", taraxacum.Conf.Debug, blackdatura.DefaultLumberjack())
	log = blackdatura.New()
}

func main() {
	log.Info("file server start", zap.Time("start time", time.Now()))

	r := taraxacum.Router()
	server := &http.Server{
		Addr:    taraxacum.Conf.Addr,
		Handler: r,
	}
	handleSignal(server)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("server listen error", zap.Error(err))
	}
}

// handleSignal handles system signal for graceful shutdown.
func handleSignal(server *http.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		sig := <-c
		log.Info("signal notify",
			zap.Any("syscall", sig),
			zap.Time("time", time.Now()),
		)

		if err := server.Close(); err != nil {
			log.Error("server close error", zap.Error(err))
		}
		os.Exit(0)
	}()
}
