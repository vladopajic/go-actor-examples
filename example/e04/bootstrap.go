package e04

import (
	"time"
)

// This example shows improved Countdown actor from example #03.
func Run() {
	launchReadySigC := make(chan struct{})

	a := NewCountdownActor(launchReadySigC)

	a.Start()

	// Stop before countdown has ended
	time.Sleep(time.Second * 2)
	a.Stop()

	// This program will wait for launchReadySigC but it will never
	// happen because countdown was stopped. Program will exit anyway
	// because all goroutines are asleep.
	<-launchReadySigC
}
