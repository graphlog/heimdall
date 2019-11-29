package main

import (
	"os"

	root "github.com/graphlog/heimdall/pkg"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	server, _ := root.InitializeApp()

	log.Info("Application is starting on http://localhost", server.Port)
	if err := fasthttp.ListenAndServe(server.Port, server.RequestHandler); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
