package config

import (
	"os"
	"strconv"
	"strings"
)

var Conf Config

type Config struct {
	DB        DBConfig
	DebugMode string
	SecretKey string
}

type DBConfig struct {
	DbDriver   string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPass     string
	DbName     string
	DbMaxConns int `default:"10"`
	DbMaxIdle  int `default:"5"`
}

func New() *Config {
	Conf = Config{
		DB: DBConfig{
			DbDriver:   getEnv("DB_DRIVER", "sqlite"),
			DbHost:     getEnv("DB_HOST", "./api.db"),
			DbMaxConns: getEnvAsInt("DB_MAX_CONNECT", 10),
			DbMaxIdle:  getEnvAsInt("DB_MAX_IDLE", 5),
		},
		DebugMode: getEnv("DEBUG_MODE", "debug"),
		SecretKey: getEnv("SECRET_KEY", "qwertyuiop"),
	}
	return &Conf
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
