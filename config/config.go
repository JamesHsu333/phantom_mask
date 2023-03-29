package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/JamesHsu333/kdan/pkg/database/postgres"
	"github.com/JamesHsu333/kdan/pkg/logger"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "Search microservice config path")
}

// App config struct
type Config struct {
	Server      ServerConfig
	Postgres    postgres.Config
	HealthCheck HealthCheck
	Logger      logger.Config
}

// Server config struct
type ServerConfig struct {
	Name              string
	Description       string
	Port              string
	HttpPort          string
	PprofPort         string
	Mode              string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	SSL               bool
	CtxDefaultTimeout time.Duration
	Debug             bool
	MaxConnectionIdle time.Duration
	Timeout           time.Duration
	MaxConnectionAge  time.Duration
	Time              time.Duration
}

// Metrics config
type HealthCheck struct {
	URL         string
	ServiceName string
	Path        string
}

func InitConfig() (*Config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		log.Fatalf("getConfigPath: %v", err)
	}

	cfgFile, err := loadConfig(configPath)
	if err != nil {
		log.Fatalf("loadConfig: %v", err)
	}

	cfg, err := parseConfig(cfgFile)
	if err != nil {
		log.Fatalf("parseConfig: %v", err)
	}

	return cfg, nil
}

// Load config file from given path
func loadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

// Parse config file
func parseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}

// get config path for local or docker
func getConfigPath() (string, error) {
	if configPath == "" {
		configPathFromEnv := os.Getenv("config")
		if configPathFromEnv != "" {
			configPath = configPathFromEnv
		} else {
			getwd, err := os.Getwd()
			if err != nil {
				return "", errors.Wrap(err, "os.Getwd")
			}
			configPath = fmt.Sprintf("%s/config/config.yaml", getwd)
		}
	}
	if configPath == "docker" {
		return "./config/config-docker", nil
	}
	return "./config/config-local", nil
}
