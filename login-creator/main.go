package main

import (
	"login-creator/lifecycle"
	"login-creator/logger"
)

func main() {
	logger.LogDebug("login-creator service")

	lifecycle.WaitForSigint()
	logger.LogDebug("... closing application")
}
