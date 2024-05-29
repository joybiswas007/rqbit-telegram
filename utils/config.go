package utils

import (
	"fmt"
	"os"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/joybiswas007/rqbit-telegram/client"
)

var (
	config *client.Config
	once   sync.Once
)

// LoadConfig initializes the configuration
func LoadConfig(filePath string) error {
	var err error
	once.Do(func() {
		var cfg client.Config
		if _, statErr := os.Stat(filePath); os.IsNotExist(statErr) {
			err = fmt.Errorf("config file does not exist: %s", filePath)
			return
		}

		if _, decodeErr := toml.DecodeFile(filePath, &cfg); decodeErr != nil {
			err = fmt.Errorf("error decoding config file: %v", decodeErr)
			return
		}
		config = &cfg
	})
	return err
}

// GetConfig returns the loaded configuration
func GetConfig() *client.Config {
	return config
}
