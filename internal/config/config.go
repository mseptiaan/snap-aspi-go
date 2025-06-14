package config

import (
	"fmt"
	"time"
)

// Config represents the application configuration
type Config struct {
	App    AppConfig    `yaml:"app"`
	ASPI   ASPIConfig   `yaml:"aspi"`
	Logger LoggerConfig `yaml:"logging"`
}

// AppConfig represents application-level configuration
type AppConfig struct {
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
	Environment string `yaml:"environment"`
}

// LoggerConfig represents logging configuration
type LoggerConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
	Output string `yaml:"output"`
}

// ASPIConfig represents ASPI API configuration
type ASPIConfig struct {
	BaseURL        string          `yaml:"base_url"`
	ClientID       string          `yaml:"client_id"`
	ClientSecret   string          `yaml:"client_secret"`
	PrivateKeyPath string          `yaml:"private_key_path"`
	PublicKeyPath  string          `yaml:"public_key_path"`
	Endpoints      EndpointsConfig `yaml:"endpoints"`
	Timeouts       TimeoutsConfig  `yaml:"timeouts"`
	Retry          RetryConfig     `yaml:"retry"`
}

// EndpointsConfig represents API endpoint configuration
type EndpointsConfig struct {
	B2BToken   string `yaml:"b2b_token"`
	B2B2CToken string `yaml:"b2b2c_token"`
}

// TimeoutsConfig represents timeout configuration
type TimeoutsConfig struct {
	ConnectTimeout time.Duration `yaml:"connect_timeout"`
	RequestTimeout time.Duration `yaml:"request_timeout"`
}

// RetryConfig represents retry configuration
type RetryConfig struct {
	MaxAttempts     int           `yaml:"max_attempts"`
	BackoffDuration time.Duration `yaml:"backoff_duration"`
}

// validate validates the configuration
func (c *Config) validate() error {
	if c.ASPI.ClientID == "" {
		return fmt.Errorf("ASPI client ID is required")
	}
	if c.ASPI.ClientSecret == "" {
		return fmt.Errorf("ASPI client secret is required")
	}
	if c.ASPI.BaseURL == "" {
		return fmt.Errorf("ASPI base URL is required")
	}
	return nil
}

// IsDevelopment returns true if running in development environment
func (c *Config) IsDevelopment() bool {
	return c.App.Environment == "development"
}

// IsProduction returns true if running in production environment
func (c *Config) IsProduction() bool {
	return c.App.Environment == "production"
}