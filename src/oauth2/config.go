package oauth2

import (
	"os"
	"log"
)

const (
	defaultPort = "9000"
	defaultStaticFilesLocation = "./static"
)

var (
	config *ServerConfig
)

func init() {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = defaultPort
	}
	log.Printf("Server Port: %s", port)

	frontendStaticFiles := os.Getenv("FRONTEND_SOURCE_FILES")
	if frontendStaticFiles == "" {
		frontendStaticFiles = defaultStaticFilesLocation
	}
	log.Printf("Static Files: %s", frontendStaticFiles)

	config = &ServerConfig{
		Port: port,
		FrontendStaticFiles: frontendStaticFiles,
	}
}
