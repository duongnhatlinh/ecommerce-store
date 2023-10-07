package app

import (
	"context"
	"ecommercestore/internal/conf"
	"ecommercestore/internal/database"
	"ecommercestore/internal/routes"
	"errors"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start(conf conf.Config) {
	database.ConnectionDatabase(conf)
	router := routes.SetRoute()

	server := http.Server{
		Addr:    conf.Host + ":" + conf.Port,
		Handler: router,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Error().Err(err).Msg("Server ListenAndServe error")
		}
	}()

	timeWait := 15 * time.Second
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), timeWait)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server forced to shutdown")
	}
	close(quit)
	log.Info().Msg("Server exiting.")
}
