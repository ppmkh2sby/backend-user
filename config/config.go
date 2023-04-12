package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

// YAMLConfigLoader is the loader of YAML file configuration
type YAMLConfigLoader struct {
	fileLocation string
}

// NewYAMLConfigLoader return the YAMLConfigLoader
func NewYamlConfigLoader(fileLocation string) *YAMLConfigLoader {
	return &YAMLConfigLoader{
		fileLocation: fileLocation,
	}
}

// ServiceConfig stores the whole configuration for service
type ServiceConfig struct {
	ServiceData ServiceDataConfig `yaml:"service_data"`
	SourceData  SourceDataConfig  `yaml:"source_data"`
	LoginRule   LoginRuleConfig   `yaml:"login_rule"`
}

// ServiceDataConfig contains the service data configuration
type ServiceDataConfig struct {
	Address  string `yaml:"address"`
	Database string `yaml:"database"`
}

// SourceDataConfig contains the source data configuration
type SourceDataConfig struct {
	PostgresDBServer   string `yaml:"postgres_db_server"`
	PostgresDBPort     int    `yaml:"postgres_db_port"`
	PostgresDBName     string `yaml:"postgres_db_name"`
	PostgresDBUsername string `yaml:"postgres_db_username"`
	PostgresDBPassword string `yaml:"postgres_db_password"`
	PostgresDBTimeout  int    `yaml:"postgres_db_timeout"`
}

// LoginRuleConfig contains rule for singing in.
// MaxAttempt to determine max attempt limit for wrong password.
// SuspensionDuration to determine how long user will be suspended if reaches login attempt limitation in minute
type LoginRuleConfig struct {
	MaxAttempt         int    `yaml:"max_attempt"`
	SuspensionDuration string `yaml:"suspension_duration"`
}

// getRawConfig parse the configuration from YAML file
func getRawConfig(fileLocation string) (*ServiceConfig, error) {
	configByte, err := os.ReadFile(fileLocation)
	if err != nil {
		return nil, err
	}
	config := &ServiceConfig{}
	err = yaml.Unmarshal(configByte, config)
	if err != nil {
		return nil, err
	}
	return config, err
}

// GetServiceConfig get raw config from YAML file Location
func (c *YAMLConfigLoader) GetServiceConfig() (*ServiceConfig, error) {
	config, err := getRawConfig(c.fileLocation)
	if err != nil {
		return nil, fmt.Errorf("unable to get raw config content: %v", err)
	}
	return config, nil
}
