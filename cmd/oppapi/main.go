package main

import (
	"log/slog"
	"os"

	"oppapi/internal/config"
	"oppapi/internal/logging"
	"oppapi/internal/server"
)

var version string

func showStartupInfo() {
	logging.Logger.Info("Config", slog.String("application.name", config.ApplicationName()))
	logging.Logger.Info("Config", slog.String("server.http.address", config.ServerHTTPAddress()))
	logging.Logger.Info("Config", slog.String("server.http.port", config.ServerHTTPPort()))
	logging.Logger.Info("Config", slog.String("log.level", config.LogLevel()))
	logging.Logger.Info("Config", slog.String("log.format", config.LogFormat()))
	logging.Logger.Info("Config", slog.String("repository.dbname", config.RepositoryDBName()))
	logging.Logger.Info("Config", slog.String("repository.url", config.RepositoryURL()))
	logging.Logger.Info("Config", slog.String("cache.redis.address", config.CacheRedisAddress()))
	logging.Logger.Info("Config", slog.String("cache.redis.password", config.CacheRedisPassword()))
	logging.Logger.Info("Config", slog.String("cache.redis.db", config.CacheRedisDB()))
}

func main() {
	if err := config.Init(); err != nil {
		slog.Error("Invalid config, exiting",
			slog.String("err",
				err.Error()))
		os.Exit(1)
	}

	// Start logger now as it may require to change log format.
	if err := logging.Init(version, config.LogLevel(), config.LogFormat()); err != nil {
		slog.Error("Error initializing logger, exiting",
			slog.String("err",
				err.Error()))
		os.Exit(2)
	}

	showStartupInfo()

	// Start web service API.
	s, err := server.New(version)
	if err != nil {
		slog.Error("Error creating server, exiting",
			slog.String("err",
				err.Error()))
		os.Exit(3)
	}

	// Run the server listener.
	err = s.Run()
	if err != nil {
		slog.Error("Error running server, exiting",
			slog.String("err",
				err.Error()))
		os.Exit(4)
	}
}
