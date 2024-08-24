package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGettersWithStringRV(t *testing.T) {
	tests := []struct {
		label    string
		getter   func() string
		expected string
	}{
		{
			label:    "ApplicationName",
			getter:   ApplicationName,
			expected: "oppapi3",
		},
		{
			label:    "ServerHTTPAddress",
			getter:   ServerHTTPAddress,
			expected: "localhost3",
		},
		{
			label:    "LogLevel",
			getter:   LogLevel,
			expected: "DEBUG3",
		},
		{
			label:    "LogFormat",
			getter:   LogFormat,
			expected: "text3",
		},
		{
			label:    "ServerHTTPPort",
			getter:   ServerHTTPPort,
			expected: "80903",
		},
		{
			label:    "RepositoryDBName",
			getter:   RepositoryDBName,
			expected: "mongo",
		},
		{
			label:    "RepositoryURL",
			getter:   RepositoryURL,
			expected: "mongodb://localhost:27017",
		},
		{
			label:    "AutXAPIKey",
			getter:   AuthXAPIKEy,
			expected: "something",
		},
	}
	input := `application:
  name: oppapi
server:
  http:
    address: localhost3
    port: 80903
log:
  level: DEBUG3
  format: text3
repsitory:
  dbname: mongo
  url: mongodb://localhost:27017
auth:
  x_api_key: something
`
	makeTestConfigFile(t, input)
	defer cleanupTestConfigFile(t)

	err := Init()
	assert.NoError(t, err)

	for _, td := range tests {
		t.Run(td.label, func(t *testing.T) {
			result := td.getter()
			assert.Equal(t, td.expected, result)
		})
	}
}
