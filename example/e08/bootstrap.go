package e08

import (
	"fmt"

	"github.com/vladopajic/go-actor/actor"
)

const endWorkAtStep = 100

// But what if producer finishes much faster then consumer?
// Note there is just 64 messages in stdout, but we wanted to end at step 100.
//
// Solution for this puzzle in next example. This puzzle is bit harder as it requires
// deep understanding of actor.Mailbox is implemented.
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
