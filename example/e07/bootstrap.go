package e07

import (
	"fmt"

	"github.com/vladopajic/go-actor/actor"
)

const endWorkAtStep = 3

// This example is improvement of example 6.
// Here consumer will be ended when there is nothing more to be consumed.
// This is achieved by stopping mailbox when producer is sopped.
//
// See change in consumer.go.
func Run() {
	finishedC := make(chan any)

	mbx := actor.NewMailbox[int]()
	mbx.Start()

	a := actor.Combine(
		actor.New(&producerWorker{outMbx: mbx}, actor.OptOnStop(mbx.Stop)),
		actor.New(&consumerWorker{inMbx: mbx}),
	).WithOptions(
		actor.OptOnStartCombined(func(_ actor.Context) {
			fmt.Println("example started")
		}),
		actor.OptOnStopCombined(func() {
			fmt.Println("example finished")
			close(finishedC)
		}),
	).Build()

	a.Start()
	defer a.Stop()

	<-finishedC
}
