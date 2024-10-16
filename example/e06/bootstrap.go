package e06

import (
	"fmt"

	"github.com/vladopajic/go-actor/actor"
)

const endWorkAtStep = 3

// Problem in example 5 is happening because combined actor was created using 3 actors: consumer, producer and mailbox.
// Because consumer and producers were stopped by themselves there was only mailbox actor that was left and
// it's goroutine was blocked because it was waiting on messages, but no messages could be sent because there was no
// other gorutines alive that could possibly send to this mailbox - hance deadlock message in stdout.
//
// To fix this, combined actor is created only using consumer and producer. Therefore `OnStop` will be called because at
// one point both actors will be stopped.
func Run() {
	finishedC := make(chan any)

	mbx := actor.NewMailbox[int]()

	cp := actor.Combine(
		actor.New(&producerWorker{outMbx: mbx}),
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

	a := actor.Combine(cp, mbx).Build()

	a.Start()
	defer a.Stop()

	<-finishedC
}
