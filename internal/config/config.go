package config

import "os"

const defaultJWTSecret = "dev-secret-change-me"

type Config struct {
	HTTPAddr           string
	GRPCAddr           string
	TCPAddr            string
	UDPAddr            string
	DBPath             string
	CatalogDBPath      string
	JWTSecret          string
	LogLevel           string
	usingDefaultSecret bool
}

func Load() (Config, error) {
	secret := envOrDefault("JWT_SECRET", defaultJWTSecret)
	return Config{
		HTTPAddr:           envOrDefault("HTTP_ADDR", ":8080"),
		GRPCAddr:           envOrDefault("GRPC_ADDR", ":9090"),
		TCPAddr:            envOrDefault("TCP_ADDR", ":9000"),
		UDPAddr:            envOrDefault("UDP_ADDR", ":9001"),
		DBPath:             envOrDefault("DB_PATH", "mangahub.db"),
		CatalogDBPath:      envOrDefault("CATALOG_DB_PATH", "catalog.db"),
		JWTSecret:          secret,
		LogLevel:           envOrDefault("LOG_LEVEL", "info"),
		usingDefaultSecret: secret == defaultJWTSecret,
	}, nil
}

func (c Config) UsingDefaultSecret() bool {
	return c.usingDefaultSecret
}

func envOrDefault(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return fallback
}
