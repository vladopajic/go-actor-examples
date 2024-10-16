package e10

import (
	"github.com/vladopajic/go-actor/actor"
)

// This example shows how to stop all combined actors when any of actors
// has been stopped.
func Run() {
	finishedC := make(chan any)

	a := actor.Combine(
		actor.New(&solverWorker{searchedNumber: 42}),
		actor.New(&solverWorker{searchedNumber: 42}),
		actor.New(&solverWorker{searchedNumber: 10000}), // this solver will never find solution
	).WithOptions(
		actor.OptStopTogether(), // this option will stop all actors when any of actor is sopped
		actor.OptOnStopCombined(func() { close(finishedC) }),
	).Build()

	a.Start()
	//	defer a.Stop(), not needed because combined actor is stopped together

	<-finishedC
}
