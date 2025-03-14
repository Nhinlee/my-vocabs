package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/getsops/sops/v3/decrypt"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type Env string

const (
	LOCAL Env = "local"
	PROD  Env = "prod"
)

type Config struct {
	Env               Env      `mapstructure:"env" yaml:"env"`
	HttpServerAddress string   `mapstructure:"http_server_address" yaml:"http_server_address"`
	AllowOrigins      []string `mapstructure:"allow_origins" yaml:"allow_origins"`

	Secrets Secrets
}

type Secrets struct {
	DBSource string `mapstructure:"db_source" yaml:"db_source"`
	GSAKey   string `mapstructure:"gsa_key" yaml:"gsa_key"`
	GSAEmail string `mapstructure:"gsa_email" yaml:"gsa_email"`
}

func LoadConfigAndSecrets() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("cfg")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	err = config.loadSecrets()
	return &config, err
}

func (c *Config) loadSecrets() error {
	encryptedPath := "."
	filePath := filepath.Join(encryptedPath, "secret.yaml")

	var decryptedData []byte
	var err error
	if c.Env == LOCAL {
		decryptedData, err = os.ReadFile(filePath)
	} else {
		decryptedData, err = decrypt.File(filePath, "yaml")
	}
	if err != nil {
		return fmt.Errorf("failed to read encrypted file: %w", err)
	}

	var secrets Secrets
	if err := yaml.Unmarshal(decryptedData, &secrets); err != nil {
		return fmt.Errorf("failed to unmarshal decrypted secrets: %w", err)
	}

	c.Secrets = secrets

	return nil
}
