package e04

import (
	"fmt"
	"time"
)

// This program shows improved example #03
func Run() {
	launchReadySigC := make(chan struct{})

	a := NewCountdownActor(launchReadySigC)

	a.Start()

	// Stop before countdown has ended
	time.Sleep(time.Second * 2)
	a.Stop()

	// This will timeout program execution because countdown was stopped
	// (there are no live goroutines which can write to this channel)
	select {
	case <-launchReadySigC:
	case <-time.After(time.Second * 5):
		fmt.Println("timeout")
	}
}
