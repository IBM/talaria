package utils

import (
	"log"
	"log/slog"
	"os"
	"strings"
)

type LBKProfile int

const (
	Localdev LBKProfile = iota
	Dev
	Prod
	Unknown
)

func GetProfile() LBKProfile {
	switch os.Getenv("LKB_PROFILE") {
	case "localdev":
		return Localdev
	case "dev":
		return Dev
	case "prod":
		return Prod
	default:
		return Unknown
	}
}

func GetLogLevel() slog.Level {
	switch strings.ToLower(os.Getenv("LOG_LEVEL")) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		log.Println("no log level set or value is invalid, setting default WARN level")
		return slog.LevelWarn

	}
}

// GetEnvVar retrieves an environment variable or returns a default value if it is not set.
//
// Params:
//
// lookUpVar (string): The name of the environment variable to retrieve.
// defaultVal (string): The default value to return if the environment variable is not set.
//
// Returns:
//
// string: The value of the environment variable or the default value if it is not set.
func GetEnvVar(lookUpVar, defaultVal string) string {
	val, ok := os.LookupEnv(lookUpVar)
	if !ok {
		return defaultVal
	} else {
		return val
	}
}
