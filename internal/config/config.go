package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

// Config represents the application configuration
type Config struct {
	App    AppConfig    `yaml:"app"`
	Server ServerConfig `yaml:"server"`
	Logger LoggerConfig `yaml:"logging"`
	ASPI   ASPIConfig   `yaml:"aspi"`
}

// AppConfig represents application-level configuration
type AppConfig struct {
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
	Environment string `yaml:"environment"`
}

// ServerConfig represents HTTP server configuration
type ServerConfig struct {
	Host         string        `yaml:"host"`
	Port         string        `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
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

// Load loads configuration from YAML file with environment variable substitution
func Load(configPath string) (*Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Substitute environment variables
	configData := expandEnvVars(string(data))

	var config Config
	if err := yaml.Unmarshal([]byte(configData), &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Validate required fields
	if err := config.validate(); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return &config, nil
}

// expandEnvVars replaces ${VAR_NAME:default_value} patterns with environment variables
func expandEnvVars(data string) string {
	// Simple environment variable expansion for ${VAR:default} format
	result := data
	start := 0
	for {
		begin := strings.Index(result[start:], "${")
		if begin == -1 {
			break
		}
		begin += start

		end := strings.Index(result[begin:], "}")
		if end == -1 {
			break
		}
		end += begin

		// Extract variable specification: VAR_NAME:default_value
		varSpec := result[begin+2 : end]
		parts := strings.SplitN(varSpec, ":", 2)
		varName := parts[0]
		defaultValue := ""
		if len(parts) > 1 {
			defaultValue = parts[1]
		}

		// Get environment variable value or use default
		envValue := os.Getenv(varName)
		if envValue == "" {
			envValue = defaultValue
		}

		// Replace the variable placeholder
		result = result[:begin] + envValue + result[end+1:]
		start = begin + len(envValue)
	}

	return result
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

// GetServerAddress returns the full server address
func (c *Config) GetServerAddress() string {
	return c.Server.Host + ":" + c.Server.Port
}

// IsDevelopment returns true if running in development environment
func (c *Config) IsDevelopment() bool {
	return c.App.Environment == "development"
}

// IsProduction returns true if running in production environment
func (c *Config) IsProduction() bool {
	return c.App.Environment == "production"
}
