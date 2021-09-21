package main

import (
	"github.com/markamdev/doorkeeper/logger"
	"github.com/namsral/flag"
)

func main() {
	logger.LogDebug("login-creator service")
	flag.Parse()
}
