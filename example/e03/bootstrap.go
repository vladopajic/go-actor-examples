package e03

import (
	"fmt"
)

// This program will demonstrate how to create actors with options
func Run() {
	launchReadySigC := make(chan struct{})

	a := NewCountdownActor(launchReadySigC)

	a.Start()
	// Note: It's good practice to stop actors, but in this cases
	// worker will end after countdown is finished
	defer a.Stop()

	<-launchReadySigC
	launchRocket()
}

func launchRocket() {
	fmt.Println("Launching rocket!")
}
