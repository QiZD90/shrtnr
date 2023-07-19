package main

import (
	"log"
	"net/http"
	"os"

	"github.com/QiZD90/shrtnr/config"
	"github.com/QiZD90/shrtnr/internal/mux"
	"github.com/QiZD90/shrtnr/internal/repository"
	"github.com/QiZD90/shrtnr/internal/service"
)

func main() {
	errorLog := log.New(os.Stderr, "[ERROR] ", log.Ltime|log.Llongfile)
	infoLog := log.New(os.Stdout, "[INFO] ", log.Ltime)

	// Parse config
	cfg, err := config.Parse()
	if err != nil {
		errorLog.Fatal(err)
	}

	// Connect to repo
	repo := repository.NewRedisRepository(cfg.Redis.Addr, cfg.Redis.Password, cfg.Redis.Expiration)

	// Configure service
	service := &service.ShrtnrService{
		ErrorLog:   errorLog,
		InfoLog:    infoLog,
		URLPrefix:  cfg.URLPrefix,
		Repository: repo,
	}

	// Start the server
	server := http.Server{
		Addr:         cfg.Server.Addr,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,

		ErrorLog: errorLog,

		Handler: mux.Get(service),
	}

	if err := server.ListenAndServe(); err != nil {
		errorLog.Fatal(err)
	}
}
