package lifecycle

import (
	"os"
	"os/signal"
)

func WaitForSigint() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	<-sigchan
}
