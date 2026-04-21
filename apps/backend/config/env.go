package config

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LoadAppEnv() error {
	return LoadDotEnv(".env", "apps/backend/.env", "../.env", "../../.env")
}

func LoadDotEnv(paths ...string) error {
	for _, path := range paths {
		if err := loadDotEnvFile(path); err != nil {
			return err
		}
	}

	return nil
}

func loadDotEnvFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		return fmt.Errorf("open env file %s: %w", path, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		line = strings.TrimPrefix(line, "export ")
		key, value, found := strings.Cut(line, "=")
		if !found {
			return fmt.Errorf("parse env file %s line %d: missing '='", path, lineNumber)
		}

		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)
		if key == "" {
			return fmt.Errorf("parse env file %s line %d: empty key", path, lineNumber)
		}

		if unquoted, err := strconv.Unquote(value); err == nil {
			value = unquoted
		}

		if _, exists := os.LookupEnv(key); !exists {
			if err := os.Setenv(key, value); err != nil {
				return fmt.Errorf("set env %s: %w", key, err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("read env file %s: %w", path, err)
	}

	return nil
}
