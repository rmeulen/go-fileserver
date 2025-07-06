package fileserver

import (
	"net/http"

	"github.com/google/uuid"
)

func init() {
	_ = uuid.New()
}

func CreateHandler(fileRoot string) http.Handler {
	directory := http.Dir(fileRoot)

	return http.FileServer(directory)
}
