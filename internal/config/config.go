package config

import (
	"github.com/spf13/viper"
)

var (
	options  *Options
	Path     = DefaultConfigPathName
	FileName = DefaultConfigFileName
)

// Options is a viper embedding.
type Options struct {
	*viper.Viper
}

// Init loads configuration first using defaults, then from a config file.
func Init() error {
	// load config from env variables
	options = &Options{viper.New()}

	// Set defaults for all env application settings
	initDefaults()

	// Bind viper names with env variables.
	if err := bindEnvVars(); err != nil {
		return err
	}

	// Use config file to override defaults.
	if err := loadConfigFromFile(); err != nil {

		return err
	}

	return nil
}

func initDefaults() {
	// application settings
	options.Viper.SetDefault("application.name", DefaultApplicationName)

	// server settings
	options.Viper.SetDefault("server.http.address", DefaultServerHTTPAddress)
	options.Viper.SetDefault("server.http.port", DefaultServerHTTPPort)

	// logging settings
	options.Viper.SetDefault("log.level", DefaultLogLevel)
	options.Viper.SetDefault("log.format", DefaultLogFormat)

	// backend
	options.Viper.SetDefault("repository.dbname", DefaultRepositoryDBName)
	options.Viper.SetDefault("repository.url", DefaultRepositoryURL)
}

func bindEnvVars() error {
	var err error
	err = options.Viper.BindEnv("application.name", "APPLICATION_NAME")
	if err != nil {
		return err
	}
	err = options.Viper.BindEnv("server.http.address", "SERVER_HTTP_ADDRESS")
	if err != nil {
		return err
	}
	err = options.Viper.BindEnv("server.http.port", "SERVER_HTTP_PORT")
	if err != nil {
		return err
	}
	err = options.Viper.BindEnv("log.level", "LOG_LEVEL")
	if err != nil {
		return err
	}
	err = options.Viper.BindEnv("log.format", "LOG_FORMAT")
	if err != nil {
		return err
	}
	err = options.Viper.BindEnv("repository.dbname", "REPOSITORY_DBNAME")
	if err != nil {
		return err
	}
	err = options.Viper.BindEnv("repository.url", "REPOSITORY_URL")
	if err != nil {
		return err
	}
	err = options.Viper.BindEnv("cache.redis.address", "CACHE_REDIS_ADDRESS")
	if err != nil {
		return err
	}
	err = options.Viper.BindEnv("cache.redis.password", "CACHE_REDIS_PASSWORD")
	if err != nil {
		return err
	}
	err = options.Viper.BindEnv("cache.redis.db", "CACHE_REDIS_DB")
	if err != nil {
		return err
	}
	return nil
}

func loadConfigFromFile() error {
	options.Viper.AddConfigPath(Path)
	options.Viper.SetConfigName(FileName)
	options.Viper.SetConfigType("yaml")
	if err := options.Viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
