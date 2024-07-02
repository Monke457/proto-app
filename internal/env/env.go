package env

import (
	"log/slog"
	"os"
	"strings"
)

func init() {
	data, err := os.ReadFile(".env")
	if err != nil {
		slog.Warn("Failed to read .env file")
	}

	vars := parse(data)
	failed := make([]string, 0)

	for key, val := range vars {
		if _, ok := os.LookupEnv(key); ok {
			continue
		}
		err := os.Setenv(key, val)
		if err != nil {
			failed = append(failed, key)
		}
	}

	if len(failed) > 0 {
		slog.Warn("Failed to set environment variables", "Variables", failed)
	}
}

func Get(key string) string {
	value, ok := os.LookupEnv(key); 
	if !ok {
		slog.Warn("Environment variable does not exist", "Key", key)
	}
	return value 
}

func parse(data []byte) map[string]string {
	lines := strings.Split(string(data), "\n")
	vars := make(map[string]string)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || line[0] == '#' {
			continue
		}

		key, val, f := strings.Cut(line, "=")
		if !f {
			continue
		}

		key = strings.Trim(key, "\" ")
		if strings.ContainsRune(val, '#') {
			val, _, _ = strings.Cut(val, "#")
		}
		val = strings.Trim(val, "\" ")

		vars[key] = val
	}

	return vars
}
