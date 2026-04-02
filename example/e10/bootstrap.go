package e10

import (
	"github.com/vladopajic/go-actor/actor"
)

// This example shows how to stop all combined actors as soon as
// any one of them stops.
func Run() {
	finishedC := make(chan any)

	a := actor.Combine(
		actor.New(&solverWorker{searchedNumber: 42}),
		actor.New(&solverWorker{searchedNumber: 42}),
		actor.New(&solverWorker{searchedNumber: 10000}), // this solver will never find solution
	).WithOptions(
		actor.OptStopTogether(), // stops all actors as soon as any one of them stops
		actor.OptOnStopCombined(func() { close(finishedC) }),
	).Build()

	a.Start()
	//	defer a.Stop(), not needed because combined actor is stopped together

	<-finishedC
}
