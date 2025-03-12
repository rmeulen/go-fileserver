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

func main() {

	viper.AutomaticEnv()
	viper.SetDefault(portEnvName, defaultPort)
	viper.SetDefault(fileRootEnvName, defaultFileRoot)

	port := viper.GetInt(portEnvName)
	fileRoot := viper.GetString(fileRootEnvName)

	fileServer := fileserver.CreateHandler(fileRoot)

	// Log the port and file root
	log.Printf("Listening on port: %d", port)
	log.Printf("Using file root: %s", fileRoot)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), fileServer))

}
