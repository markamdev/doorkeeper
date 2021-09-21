package main

import (
	"github.com/markamdev/doorkeeper/lifecycle"
	"github.com/markamdev/doorkeeper/logger"
)

func main() {
	logger.LogDebug("authenticator service")

	lifecycle.WaitForSigint()
	logger.LogDebug("... closing application")
}
