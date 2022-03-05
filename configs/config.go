package configs

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	DB     DbConfig
	SERVER ServerConfig
	AUTH   AuthConfig
}

func New() *Config {
	return &Config{
		DB: DbConfig{
			DbHost:     getEnv("DB_HOST", "localhost"),
			DbPort:     uint(getEnvAsInt("DB_PORT", 5432)),
			DbName:     getEnv("DB_NAME", "postgres"),
			DbUsername: getEnv("DB_USERNAME", "postgres"),
			DbPassword: getEnv("DB_PASSWORD", "postgres"),
			DbLink:     getEnv("DB_LINK", "postgres://postgres:postgres@localhost:5432/started_kit"),
		},
		SERVER: ServerConfig{
			ServerPort:  uint(getEnvAsInt("SERVER_PORT", 3000)),
			ServerDebug: getEnvAsBool("DEBUG", true),
		},
		AUTH: AuthConfig{
			JwtAccessKey:  getEnv("JWT_ACCESS_KEY", ""),
			JwtAccessTTL:  uint(getEnvAsInt("JWT_ACCESS_TTL", 1)),
			JwtRefreshKey: getEnv("JWT_REFRESH_KEY", ""),
			JwtRefreshTTL: uint(getEnvAsInt("JWT_REFRESH_TTL", 1)),
		},
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsSlice(key string, defaultValue []string, sep string) []string {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}
	value := strings.Split(valueStr, sep)
	return value
}
