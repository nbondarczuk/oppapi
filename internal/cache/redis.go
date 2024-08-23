package cache

import (
	"oppapi/internal/logging"
	"log/slog"
	"strconv"

	"github.com/go-redis/redis"
)

var (
	Address  string
	Password string
	DB       int
	Client   *redis.Client
)

type Redis struct {
	Client *redis.Client
}

// InitWithRedis is a singe initialisation point for the package state.
// All subsequent usage of it are done with its state variables.
func InitWithRedis(address, password, db string) error {
	Address = address
	Password = password
	var err error
	DB, err = strconv.Atoi(db)
	if err != nil {
		return err
	}
	// The connection is allocated upon init.
	client := redis.NewClient(&redis.Options{
		Addr:     Address,
		Password: Password,
		DB:       DB,
	})
	// It is checked.
	if err := client.Ping().Err(); err != nil {
		return err
	}
	logging.Logger.Debug("Success pinging Redis DB", slog.String("address", address))
	return nil
}

// WithRedis is using the module instnal state with the connection created within init.
func WithRedis() (*Redis, error) {
	return &Redis{
			Client: Client,
		},
		nil
}
