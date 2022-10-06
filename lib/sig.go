package lib

import (
	"os"
	"os/signal"
	"syscall"
)

// WaitForTermination returns new signal channel that will be notified when
// process receives terminate signal.
func WaitForTermination() <-chan struct{} {
	sig := make(chan os.Signal, 1)
	done := make(chan struct{})

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sig
		close(done)
	}()

	return done
}
