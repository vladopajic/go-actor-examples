package e09

import (
	"fmt"

	"github.com/vladopajic/go-actor/actor"
)

const endWorkAtStep = 100

// Since mailbox is stopped after producer finishes, there will still be messages in queue,
// but mailbox actor will not process them because it was sopped.
// To fix this mailbox actor should be active until there are messages in queue.
func Run() {
	finishedC := make(chan any)

	mbx := actor.NewMailbox[int](
		// This option will end mailbox after it is stopped and queue is fully emptied - that is
		// all messages from mailbox are fully received
		actor.OptStopAfterReceivingAll(),
	)
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
