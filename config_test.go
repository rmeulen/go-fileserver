package main

import (
	"os"
	"testing"
)

func TestInitConfig(t *testing.T) {
	// Test default values
	config := initConfig()
	
	if config.Port != defaultPort {
		t.Errorf("Expected port %d, got %d", defaultPort, config.Port)
	}
	
	if config.FileRoot != defaultFileRoot {
		t.Errorf("Expected file root %s, got %s", defaultFileRoot, config.FileRoot)
	}
}

func TestInitConfigWithEnvVars(t *testing.T) {
	// Set environment variables
	os.Setenv("PORT", "9090")
	os.Setenv("FILE_ROOT", "/tmp/test")
	defer func() {
		os.Unsetenv("PORT")
		os.Unsetenv("FILE_ROOT")
	}()
	
	config := initConfig()
	
	if config.Port != 9090 {
		t.Errorf("Expected port %d, got %d", 9090, config.Port)
	}
	
	if config.FileRoot != "/tmp/test" {
		t.Errorf("Expected file root %s, got %s", "/tmp/test", config.FileRoot)
	}
}