package main

import (
	"github.com/rmeulen/go-fileserver/fileserver"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

const portEnvName = "PORT"
const defaultPort = 8080

const fileRootEnvName = "FILE_ROOT"
const defaultFileRoot = "./"

func main() {

	port := determinePort()
	fileRoot := determineFileRoot()

	fileServer := fileserver.CreateHandler(fileRoot)

	log.Printf("Listening on port: %d", port)
	log.Printf("Using file root: %s", fileRoot)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), fileServer))

}

func determinePort() int {
	portEnvValue := os.Getenv(portEnvName)

	if portEnvValue == "" {
		log.Printf("No value found for environment variable %s.", portEnvName)
		log.Printf("Using default port number: %d", defaultPort)
		return defaultPort
	}

	port, err := strconv.Atoi(portEnvValue)

	if err != nil {
		log.Printf("Invalid value for environment variable %s", portEnvName)
		log.Printf("Defaulting to port number: %d", defaultPort)
		return defaultPort
	}

	log.Printf("Valid environment variable %s found: %d", portEnvName, port)
	return port
}

func determineFileRoot() string {
	fileRoot := os.Getenv(fileRootEnvName)

	if fileRoot == "" {
		log.Printf("No value found for environment variable %s", fileRootEnvName)
		log.Printf("Defaulting to file root: %s", defaultFileRoot)
		fileRoot = defaultFileRoot
	}

	log.Printf("Valid environment variable %s found: %s", fileRootEnvName, fileRoot)
	return fileRoot
}
