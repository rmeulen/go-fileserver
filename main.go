package main

import (
	"github.com/rmeulen/go-fileserver/fileserver"
	"github.com/spf13/viper"
	"fmt"
	"log"
	"net/http"
)

const portEnvName = "PORT"
const defaultPort = 8080

const fileRootEnvName = "FILE_ROOT"
const defaultFileRoot = "./"

// Config holds the application configuration
type Config struct {
	Port     int    `mapstructure:"port"`
	FileRoot string `mapstructure:"file_root"`
}

// initConfig initializes the configuration using Viper
func initConfig() *Config {
	viper.SetDefault("port", defaultPort)
	viper.SetDefault("file_root", defaultFileRoot)
	
	// Bind environment variables
	viper.BindEnv("port", portEnvName)
	viper.BindEnv("file_root", fileRootEnvName)
	
	// Enable reading from config files
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.fileserver")
	viper.AddConfigPath("/etc/fileserver")
	
	// Read config file if it exists
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf("No config file found, using defaults and environment variables")
		} else {
			log.Printf("Error reading config file: %v", err)
		}
	} else {
		log.Printf("Using config file: %s", viper.ConfigFileUsed())
	}
	
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Printf("Error unmarshaling config: %v", err)
		// Fall back to defaults
		config.Port = defaultPort
		config.FileRoot = defaultFileRoot
	}
	
	return &config
}

func main() {
	config := initConfig()
	
	port := config.Port
	fileRoot := config.FileRoot
	
	// Log the values similar to the original implementation
	log.Printf("Valid configuration found - Port: %d", port)
	log.Printf("Valid configuration found - File Root: %s", fileRoot)

	fileServer := fileserver.CreateHandler(fileRoot)

	log.Printf("Listening on port: %d", port)
	log.Printf("Using file root: %s", fileRoot)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), fileServer))
}

func determinePort() int {
	// This function is now deprecated in favor of Viper configuration
	// but kept for backward compatibility if needed
	port := viper.GetInt("port")
	
	if port == 0 {
		log.Printf("No value found for port configuration.")
		log.Printf("Using default port number: %d", defaultPort)
		return defaultPort
	}
	
	log.Printf("Valid port configuration found: %d", port)
	return port
}

func determineFileRoot() string {
	// This function is now deprecated in favor of Viper configuration
	// but kept for backward compatibility if needed
	fileRoot := viper.GetString("file_root")
	
	if fileRoot == "" {
		log.Printf("No value found for file_root configuration")
		log.Printf("Defaulting to file root: %s", defaultFileRoot)
		fileRoot = defaultFileRoot
	}
	
	log.Printf("Valid file_root configuration found: %s", fileRoot)
	return fileRoot
}
