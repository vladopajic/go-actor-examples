package e05

import (
	"fmt"

	"github.com/vladopajic/go-actor/actor"
)

const endWorkAtStep = 3

// Example 1 and 2 have demonstrated case when producer and consumer are working indefinitely.
// However, sometimes we want to end program when some condition is reached - in this example
// it is after 3 working steps.
//
// This example will signal consumer and producer actors to end. And after they both
// end we will terminate program.
func Run() {
	finishedC := make(chan any)

	mbx := actor.NewMailbox[int]()

	a := actor.Combine(
		actor.New(&producerWorker{outMbx: mbx}),
		actor.New(&consumerWorker{inMbx: mbx}),
		mbx,
	).WithOptions(
		// add function to be executed before all actors are started
		actor.OptOnStartCombined(func(_ actor.Context) {
			fmt.Println("example started")
		}),

		// add function to be executed after all actors are stopped
		actor.OptOnStopCombined(func() {
			fmt.Println("example finished")
			close(finishedC)
			// note this will never be called, this was intentional puzzle, try to think reason why.
			// solution for this puzzle is in example 6.
		}),
	).Build()

	a.Start()
	defer a.Stop()

	<-finishedC
}
