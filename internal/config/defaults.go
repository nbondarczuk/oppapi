package config

const (
	// Hardcoded initial file name, mauy be changed in test
	DefaultConfigFileName = "config.yaml"
	DefaultConfigPathName = "config"

	// Default option values
	DefaultApplicationName   = "oppapi"
	DefaultServerHTTPAddress = "localhost"
	DefaultServerHTTPPort    = "8000"
	DefaultLogLevel          = "DEBUG"
	DefaultLogFormat         = "text"

	// Log levels
	LogLevelInfo  = "INFO"
	LogLevelDebug = "DEBUG"

	// repository - mongo
	DefaultRepositoryDBName = "mongo"
	DefaultRepositoryURL    = "mongodb://localhost:27017"
)
