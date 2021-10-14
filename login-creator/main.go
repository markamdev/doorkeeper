package main

import (
	"login-creator/lifecycle"
	"login-creator/logger"
	"login-creator/rest"
	"os"
)

const (
	listenPort = "8091"
)

func main() {
	logger.LogDebug("login-creator service")

	logger.LogDebug("Creating server listening at port:", listenPort)
	srv, err := rest.CreateServer(listenPort)

	if err != nil {
		logger.LogError("Failed to create server:", err.Error())
		os.Exit(1)
	}

	srv.Start()
	logger.LogDebug("... server launched")

	lifecycle.WaitForSigint()

	logger.LogDebug("Stopping server")
	srv.Stop()

	logger.LogDebug("Closing application")
}
