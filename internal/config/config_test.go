package config

import "testing"

func TestLoadDefaults(t *testing.T) {
	for _, key := range []string{
		"HTTP_ADDR", "GRPC_ADDR", "TCP_ADDR", "UDP_ADDR",
		"DB_PATH", "CATALOG_DB_PATH", "JWT_SECRET", "LOG_LEVEL",
	} {
		t.Setenv(key, "")
	}

	cfg, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	defaults := map[string][2]string{
		"HTTPAddr":      {cfg.HTTPAddr, ":8080"},
		"GRPCAddr":      {cfg.GRPCAddr, ":9090"},
		"TCPAddr":       {cfg.TCPAddr, ":9000"},
		"UDPAddr":       {cfg.UDPAddr, ":9001"},
		"DBPath":        {cfg.DBPath, "mangahub.db"},
		"CatalogDBPath": {cfg.CatalogDBPath, "catalog.db"},
		"JWTSecret":     {cfg.JWTSecret, "dev-secret-change-me"},
		"LogLevel":      {cfg.LogLevel, "info"},
	}
	for field, values := range defaults {
		if values[0] != values[1] {
			t.Errorf("%s = %q, want %q", field, values[0], values[1])
		}
	}
	if !cfg.UsingDefaultSecret() {
		t.Fatal("expected default secret")
	}
}

func TestLoadFromEnv(t *testing.T) {
	values := map[string]string{
		"HTTP_ADDR":       ":9999",
		"GRPC_ADDR":       ":9998",
		"TCP_ADDR":        ":9997",
		"UDP_ADDR":        ":9996",
		"DB_PATH":         "custom.db",
		"CATALOG_DB_PATH": "catalog-custom.db",
		"JWT_SECRET":      "abc",
		"LOG_LEVEL":       "debug",
	}
	for key, value := range values {
		t.Setenv(key, value)
	}

	cfg, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	got := map[string]string{
		"HTTP_ADDR":       cfg.HTTPAddr,
		"GRPC_ADDR":       cfg.GRPCAddr,
		"TCP_ADDR":        cfg.TCPAddr,
		"UDP_ADDR":        cfg.UDPAddr,
		"DB_PATH":         cfg.DBPath,
		"CATALOG_DB_PATH": cfg.CatalogDBPath,
		"JWT_SECRET":      cfg.JWTSecret,
		"LOG_LEVEL":       cfg.LogLevel,
	}
	for key, want := range values {
		if got[key] != want {
			t.Errorf("%s = %q, want %q", key, got[key], want)
		}
	}
	if cfg.UsingDefaultSecret() {
		t.Fatal("secret should not be default")
	}
}
