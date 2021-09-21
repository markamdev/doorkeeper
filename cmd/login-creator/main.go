package main

import (
	"github.com/markamdev/doorkeeper/lifecycle"
	"github.com/markamdev/doorkeeper/logger"
)

func main() {
	logger.LogDebug("login-creator service")

	lifecycle.WaitForSigint()
	logger.LogDebug("... closing application")
}
