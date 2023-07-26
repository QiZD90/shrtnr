package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/QiZD90/shrtnr/config"
	v1 "github.com/QiZD90/shrtnr/internal/controller/http/v1"
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
	s := &service.ShrtnrService{
		ErrorLog:   errorLog,
		InfoLog:    infoLog,
		URLPrefix:  cfg.Service.URLPrefix,
		Repository: repo,
	}

	// Create mux
	mux := v1.NewMux(s)

	// Start the server
	server := http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,

		ErrorLog: errorLog,

		Handler: mux,
	}

	infoLog.Printf("Listening at %s:%s", cfg.Server.Host, cfg.Server.Port)
	if err := server.ListenAndServe(); err != nil {
		errorLog.Fatal(err)
	}
}
