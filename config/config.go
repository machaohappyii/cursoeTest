package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server   ServerConfig   `yaml:"server" json:"server"`
	Database DatabaseConfig `yaml:"database" json:"database"`
}

type ServerConfig struct {
	Port string `yaml:"port" json:"port"`
	Mode string `yaml:"mode" json:"mode"`
}

type DatabaseConfig struct {
	Driver   string `yaml:"driver" json:"driver"`
	Host     string `yaml:"host" json:"host"`
	Port     string `yaml:"port" json:"port"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	Name     string `yaml:"name" json:"name"`
	Charset  string `yaml:"charset" json:"charset"`
}

func LoadConfig() *Config {
	config := &Config{}
	
	// Try to load from YAML file first
	if err := loadFromYAML(config); err != nil {
		log.Printf("Failed to load config from YAML, using defaults with env override: %v", err)
		// Fallback to default values with environment variable override
		loadDefaultConfig(config)
	}
	
	// Always allow environment variables to override YAML config
	overrideWithEnv(config)
	
	return config
}

func loadFromYAML(config *Config) error {
	configFile := "config.yaml"
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return fmt.Errorf("config file %s not found", configFile)
	}
	
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("failed to read config file: %v", err)
	}
	
	if err := yaml.Unmarshal(data, config); err != nil {
		return fmt.Errorf("failed to parse YAML config: %v", err)
	}
	
	return nil
}

func loadDefaultConfig(config *Config) {
	config.Server = ServerConfig{
		Port: "9002",
		Mode: "debug",
	}
	config.Database = DatabaseConfig{
		Driver:   "mysql",
		Host:     "localhost",
		Port:     "3306",
		Username: "root",
		Password: "123456",
		Name:     "test",
		Charset:  "utf8mb4",
	}
}

func overrideWithEnv(config *Config) {
	if port := os.Getenv("PORT"); port != "" {
		config.Server.Port = port
	}
	if mode := os.Getenv("GIN_MODE"); mode != "" {
		config.Server.Mode = mode
	}
	if driver := os.Getenv("DB_DRIVER"); driver != "" {
		config.Database.Driver = driver
	}
	if host := os.Getenv("DB_HOST"); host != "" {
		config.Database.Host = host
	}
	if port := os.Getenv("DB_PORT"); port != "" {
		config.Database.Port = port
	}
	if username := os.Getenv("DB_USERNAME"); username != "" {
		config.Database.Username = username
	}
	if password := os.Getenv("DB_PASSWORD"); password != "" {
		config.Database.Password = password
	}
	if name := os.Getenv("DB_NAME"); name != "" {
		config.Database.Name = name
	}
	if charset := os.Getenv("DB_CHARSET"); charset != "" {
		config.Database.Charset = charset
	}
}

func (c *Config) GetDatabaseDSN() string {
	switch c.Database.Driver {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
			c.Database.Username,
			c.Database.Password,
			c.Database.Host,
			c.Database.Port,
			c.Database.Name,
			c.Database.Charset,
		)
	case "sqlite3":
		return c.Database.Name
	default:
		return c.Database.Name
	}
}

func (c *Config) GetServerAddr() string {
	return fmt.Sprintf(":%s", c.Server.Port)
}
