package example

import (
	"fmt"
	"time"

	"github.com/vladopajic/go-actor/actor"
)

func init() {
	register(7, "stop_mailbox_on_producer_stop", Run07)
}

// Run07 improves Run06. Here consumer will be ended when there is nothing more
// to be consumed. This is achieved by stopping mailbox when producer is stopped.
func Run07() {
	finishedC := make(chan any)

	const endAtStep = 3

	mbx := actor.NewMailbox[int]()
	mbx.Start()

	a := actor.Combine(
		actor.New(&stopMailboxProducer{outMbx: mbx, endAtStep: endAtStep}, actor.OptOnStop(mbx.Stop)),
		actor.New(&stopMailboxConsumer{inMbx: mbx}),
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

type stopMailboxProducer struct {
	outMbx    actor.MailboxSender[int]
	num       int
	endAtStep int
}

func (w *stopMailboxProducer) DoWork(c actor.Context) actor.WorkerStatus {
	select {
	case <-c.Done():
		return actor.WorkerEnd

	case <-time.After(time.Second):
		w.num++
		w.outMbx.Send(c, w.num) //nolint:errcheck // This example assumes the mailbox will never stop before this worker.

		if w.num == w.endAtStep {
			return actor.WorkerEnd
		}

		return actor.WorkerContinue
	}
}

type stopMailboxConsumer struct {
	inMbx actor.MailboxReceiver[int]
}

func (w *stopMailboxConsumer) DoWork(c actor.Context) actor.WorkerStatus {
	select {
	case <-c.Done():
		return actor.WorkerEnd

	case num, isOpen := <-w.inMbx.ReceiveC():
		if !isOpen {
			// Mailbox is closed, so no more messages can arrive and worker can end.
			return actor.WorkerEnd
		}

		fmt.Printf("consumed %d\n", num)

		return actor.WorkerContinue
	}
}
