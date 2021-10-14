package main

import (
	"token-validator/lifecycle"
	"token-validator/logger"
)

func main() {
	logger.LogDebug("authenticator service")

	lifecycle.WaitForSigint()
	logger.LogDebug("... closing application")
}
