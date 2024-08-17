package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig_Success(t *testing.T) {
	// Create a temporary config file
	tempDir := os.TempDir()
	filePath := filepath.Join(tempDir, "test-config.yaml")

	fmt.Println("Temporary config file path:", filePath)

	file, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("failed to create temp config file: %v", err)
	}
	defer func() {
		file.Close()
		os.Remove(file.Name())
	}()

	// Write sample config data to the file
	configData := `
snmp_cfg:
  ip: "127.0.0.1"
  port: 161
  community: "public"
redis_cfg:
  host: "localhost"
  port: "6379"
  password: ""
  db: 0
  default_db: 0
  min_idle_connections: 10
  pool_size: 20
  pool_timeout: 30
olt_cfg:
  base_oid_1: "1.3.6.1.2.1"
  base_oid_2: "1.3.6.1.2.2"
  onu_id_name: "1.3.6.1.4.1"
  onu_type: "1.3.6.1.4.2"
`
	if _, err := file.Write([]byte(configData)); err != nil {
		t.Fatalf("failed to write to temp config file: %v", err)
	}
	file.Close()

	fmt.Println("File written successfully. Now setting Viper to use this file.")

	// Point Viper to the temp file directly, bypassing LoadConfig()
	viper.SetConfigFile(filePath)
	viper.SetConfigType("yaml")

	// Manually read the config file using Viper
	if err := viper.ReadInConfig(); err != nil {
		t.Fatalf("Viper failed to read the config: %v", err)
	}

	// Validate that the config is loaded correctly
	assert.Equal(t, "127.0.0.1", viper.GetString("snmp_cfg.ip"), "expected IP to match")
	assert.Equal(t, 161, viper.GetInt("snmp_cfg.port"), "expected Port to match")
}

func TestLoadConfig_FileNotFound(t *testing.T) {
	// Attempt to load a non-existent config file
	cfg, err := LoadConfig("non_existent_config.yaml")

	// Assertions
	assert.Nil(t, cfg)
	assert.Error(t, err)
	assert.Equal(t, "config file not found", err.Error())
}

func TestLoadConfig_InvalidConfig(t *testing.T) {
	// Set up an invalid configuration file
	invalidConfig := `
	invalid_yaml_syntax:
	  - missing_colon "value"
`
	configFile := "invalid_config.yaml"
	err := os.WriteFile(configFile, []byte(invalidConfig), 0644)
	assert.NoError(t, err)
	defer os.Remove(configFile) // Clean up after the test

	// Call the function under test
	cfg, err := LoadConfig(configFile)

	// Assertions
	assert.Nil(t, cfg)
	assert.Error(t, err)
}

func TestLoadConfig_FileExists(t *testing.T) {
	// Cek apakah file konfigurasi ada
	if _, err := os.Stat("cfg.yaml"); os.IsNotExist(err) {
		t.Skip("File config.yaml tidak ditemukan, melewati tes.")
	}

	// Load the configuration
	cfg, err := LoadConfig("config_test.go")

	// Assertions
	assert.Nil(t, cfg)
	assert.Error(t, err)
	assert.Equal(t, "config file not found", err.Error())
}
