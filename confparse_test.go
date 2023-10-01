package confparse

import (
	"testing"

	"gopkg.in/yaml.v3"
	"gotest.tools/v3/assert"
)

type Config struct {
	Server struct {
		Host    string `yaml:"host" required:"true"`
		Port    int    `yaml:"port" required:"true"`
		Enabled bool   `yaml:"enabled" required:"true"`
	} `yaml:"server" required:"true"`
}

func TestComplete(t *testing.T) {
	configPath := "config.yml"
	config := Config{}
	LoadConfig(configPath, &config)
}

func TestParseCompleteConfig(t *testing.T) {
	config := `
server:
  host: localhost
  port: 43000
  enabled: false
`
	cfg := &Config{}
	err := yaml.Unmarshal([]byte(config), cfg)
	if err != nil {
		assert.Error(t, err, "")
	}

	err = ParseConfig(cfg)
	if err != nil {
		assert.Error(t, err, "")
	}

}

func TestParseIncompleteConfig(t *testing.T) {
	config := ""
	cfg := &Config{}
	err := yaml.Unmarshal([]byte(config), cfg)
	if err != nil {
		assert.Error(t, err, "")
	}

	err = ParseConfig(cfg)
	if err != nil {
		assert.Error(t, err, "config field Host requires a value")
	}
}
