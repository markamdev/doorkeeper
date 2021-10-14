package main

import (
	"authenticator/lifecycle"
	"authenticator/logger"
)

func main() {
	logger.LogDebug("authenticator service")

	lifecycle.WaitForSigint()
	logger.LogDebug("... closing application")
}
