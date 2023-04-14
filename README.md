# confparse

Provide the path to a yaml config file, and a struct describing the structure of the config to validate the yaml structure.

# Usage

```go
package main

import (
	"log"
    "fmt"

	"github.com/UnionJoin/confparse/v2"
)

type Config struct {
	Server struct {
		Required bool
		Host     string `yaml:"host" required:"true"`
		Port     int `yaml:"port" required:"true"`
	} `yaml:"server" required:"true"`
}

func Run(config Config) error {
	fmt.Printf("starting server, listening on %s:%d", config.Server.Host, config.Server.Port)
	return nil
}

func main() {

	configPath := "config.yml"
	config := Config{}

	err := confparse.LoadConfig(configPath, &config)
	if err != nil {
		log.Fatal(err)
	}

	Run(config)
}
```
