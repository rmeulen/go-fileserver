package fileserver

import (
	"net/http"
)

func CreateHandler(fileRoot string) http.Handler {
	directory := http.Dir(fileRoot)

	return http.FileServer(directory)
}
