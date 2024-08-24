package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	input := `application:
  name: oppapi
server:
  http:
    address: localhost2
    port: 80902
log:
  level: DEBUG2
  format: text2
repsitory:
  dbname: mongo
  url: mongodb://localhost:27017
auth:
  x_api_key: R6bSXS4pfo7bnI0zIdMqiA==
`
	makeTestConfigFile(t, input)
	defer cleanupTestConfigFile(t)

	err := Init()
	assert.NoError(t, err)
	assert.Equal(t, "go-gin-server2", options.GetString("application.name"))
	assert.Equal(t, "localhost2", options.GetString("server.http.address"))
	assert.Equal(t, 80902, options.GetInt("server.http.port"))
	assert.Equal(t, "DEBUG2", options.GetString("log.level"))
	assert.Equal(t, "text2", options.GetString("log.format"))
	assert.Equal(t, "mongo", options.GetString("repository.dbname"))
	assert.Equal(t, "mongodb://localhost:27017", options.GetString("repository.url"))
	assert.Equal(t, "R6bSXS4pfo7bnI0zIdMqiA==", options.GetString("auth.x_api_key.url"))
}
