package config

import (
	_ "embed"
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var AppConfig Config

const (
	defaultCfgFile = "agent-config-default.yml"
)

type Config struct {
	ModelConfig ModelConfig
}

type ModelConfig struct {
	Model string
}

func InitConfig(cfgFile string, cfgPath string) (Config, error) {
	filePath := filepath.Join(cfgPath, cfgFile)
	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		filePath = filepath.Join(cfgPath, defaultCfgFile)
	}
	viper.SetConfigFile(filePath)

	// Searches for config file in given paths and read it
	if err := viper.ReadInConfig(); err != nil {
		cfgNotfound := viper.ConfigFileNotFoundError{}
		if errors.As(err, &cfgNotfound) {
			log.Fatalf("File Not found, %s", err)
		} else {
			log.Fatalf("Error reading config file, %s", err)
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	log.Printf("init config with file: %v", viper.ConfigFileUsed())
	return config, nil
}
