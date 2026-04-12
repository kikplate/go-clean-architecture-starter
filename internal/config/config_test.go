package config

import (
	"os"
	"testing"
	"time"
)

func TestLoad_MissingDatabaseURL(t *testing.T) {
	t.Setenv("DATABASE_URL", "")
	t.Setenv("HTTP_ADDR", "")
	t.Setenv("LOG_LEVEL", "")
	t.Setenv("REQUEST_TIMEOUT_MS", "")
	t.Setenv("SHUTDOWN_TIMEOUT_MS", "")
	_, err := Load()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestLoad_OK(t *testing.T) {
	t.Setenv("DATABASE_URL", "postgres://u:p@localhost:5432/db?sslmode=disable")
	t.Setenv("HTTP_ADDR", ":9090")
	t.Setenv("LOG_LEVEL", "debug")
	t.Setenv("REQUEST_TIMEOUT_MS", "5000")
	t.Setenv("SHUTDOWN_TIMEOUT_MS", "7000")
	cfg, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	if cfg.DatabaseURL != os.Getenv("DATABASE_URL") {
		t.Fatalf("database url mismatch")
	}
	if cfg.HTTPAddr != ":9090" {
		t.Fatalf("http addr: %s", cfg.HTTPAddr)
	}
	if cfg.LogLevel != "debug" {
		t.Fatalf("log level: %s", cfg.LogLevel)
	}
	if cfg.RequestTimeout != 5*time.Second {
		t.Fatalf("request timeout: %s", cfg.RequestTimeout)
	}
	if cfg.ShutdownTimeout != 7*time.Second {
		t.Fatalf("shutdown timeout: %s", cfg.ShutdownTimeout)
	}
}

func TestLoad_DefaultDurationsOnInvalid(t *testing.T) {
	t.Setenv("DATABASE_URL", "postgres://u:p@localhost:5432/db?sslmode=disable")
	t.Setenv("REQUEST_TIMEOUT_MS", "not-a-number")
	t.Setenv("SHUTDOWN_TIMEOUT_MS", "-1")
	cfg, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	if cfg.RequestTimeout != 15*time.Second {
		t.Fatalf("request timeout: %s", cfg.RequestTimeout)
	}
	if cfg.ShutdownTimeout != 20*time.Second {
		t.Fatalf("shutdown timeout: %s", cfg.ShutdownTimeout)
	}
}
