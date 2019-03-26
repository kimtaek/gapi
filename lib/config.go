package lib

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Logs struct {
	System string `yaml:"system"`
}

type AppConfig struct {
	AppSecret        string `yaml:"app_secret"`
	JwtSecret        string `yaml:"jwt_secret"`
	JwtTTL           int    `yaml:"jwt_ttl"`
	JWtTRT           int    `yaml:"jwt_trt"`
	SlackIncomingUrl string `yaml:"slack_incoming_url"`
}

type AppConfigs struct {
	AppMode string `yaml:"mode"`
	Logs    Logs
	Debug   AppConfig
	Test    AppConfig
	Release AppConfig
}

type RedisConfig struct {
	Host               string `yaml:"host"`
	Port               string `yaml:"port"`
	Password           string `yaml:"password"`
	MaxIdleConnections int    `yaml:"max_idle_connections"`
	MaxOpenConnections int    `yaml:"max_open_connections"`
}

type MysqlConfig struct {
	Host               string `yaml:"host"`
	Port               string `yaml:"port"`
	Username           string `yaml:"username"`
	Password           string `yaml:"password"`
	Database           string `yaml:"database"`
	MaxIdleConnections int    `yaml:"max_idle_connections"`
	MaxOpenConnections int    `yaml:"max_open_connections"`
}

type DatabaseConfig struct {
	RedisConfig RedisConfig `yaml:"redis"`
	MysqlConfig MysqlConfig `yaml:"mysql"`
}

type DatabaseConfigs struct {
	Debug   DatabaseConfig
	Test    DatabaseConfig
	Release DatabaseConfig
}

type Config struct {
	AppMode        string
	Logs           Logs
	AppConfig      AppConfig
	DatabaseConfig DatabaseConfig
}

var config Config

func InitConfigs() *Config {
	var err error
	var appConfigs AppConfigs
	var databaseConfigs DatabaseConfigs
	appConfigFile, _ := ioutil.ReadFile("config/app.yaml")
	databaseConfigFile, _ := ioutil.ReadFile("config/database.yaml")
	err = yaml.Unmarshal(appConfigFile, &appConfigs)
	if err != nil {
		log.Fatalf("appConfig Unmarshal: %v", err)
	}
	err = yaml.Unmarshal(databaseConfigFile, &databaseConfigs)
	if err != nil {
		log.Fatalf("databaseConfig Unmarshal: %v", err)
	}
	config.AppMode = os.Getenv("APP_MODE")
	config.Logs = appConfigs.Logs
	if config.AppMode == "" {
		config.AppMode = appConfigs.AppMode
	}
	switch config.AppMode {
	case "debug":
		config.AppConfig = appConfigs.Debug
		config.DatabaseConfig = databaseConfigs.Debug
		break
	case "test":
		config.AppConfig = appConfigs.Test
		config.DatabaseConfig = databaseConfigs.Test
		break
	case "release":
		config.AppConfig = appConfigs.Release
		config.DatabaseConfig = databaseConfigs.Release
		break
	default:
		log.Println("Undefined APP_MODE")
		os.Exit(1)
	}
	return &config
}

func GetConfigs() *Config {
	return &config
}
