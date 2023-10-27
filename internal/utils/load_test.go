package utils

import "testing"

func TestGetConfigPath(t *testing.T) {
	// Kasus uji untuk configPath "development"
	configPath := "development"
	expectedPath := "./config/config-dev"
	result := GetConfigPath(configPath)
	if result != expectedPath {
		t.Errorf("For configPath %s, got %s, expected %s", configPath, result, expectedPath)
	}

	// Kasus uji untuk configPath "heroku"
	configPath = "heroku"
	expectedPath = "./config/config-heroku"
	result = GetConfigPath(configPath)
	if result != expectedPath {
		t.Errorf("For configPath %s, got %s, expected %s", configPath, result, expectedPath)
	}

	// Kasus uji untuk configPath "production"
	configPath = "production"
	expectedPath = "./config/config-prod"
	result = GetConfigPath(configPath)
	if result != expectedPath {
		t.Errorf("For configPath %s, got %s, expected %s", configPath, result, expectedPath)
	}

	// Kasus uji default
	configPath = "unknown"
	expectedPath = "./config/cfg"
	result = GetConfigPath(configPath)
	if result != expectedPath {
		t.Errorf("For configPath %s, got %s, expected %s", configPath, result, expectedPath)
	}
}
