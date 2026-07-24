package example

import (
	"fmt"
	"time"

	"github.com/vladopajic/go-actor/actor"
)

func init() {
	register(9, "drain_mailbox_on_stop", Run09)
}

// Run09 solves the puzzle introduced in Run08.
//
// Since mailbox is stopped after producer finishes, there will still be messages
// in queue, but mailbox actor will not process them because it was stopped.
// To fix this mailbox actor should be active until there are messages in queue.
func Run09() {
	finishedC := make(chan any)

	const endAtStep = 100

	mbx := actor.NewMailbox[int](
		// This option will end mailbox after it is stopped and queue is fully
		// emptied - that is all messages from mailbox are fully received.
		actor.OptStopAfterReceivingAll(),
	)
	mbx.Start()

	a := actor.Combine(
		actor.New(&drainMailboxOnStopProducer{outMbx: mbx, endAtStep: endAtStep}, actor.OptOnStop(mbx.Stop)),
		actor.New(&drainMailboxOnStopConsumer{inMbx: mbx}),
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

type drainMailboxOnStopProducer struct {
	outMbx    actor.MailboxSender[int]
	num       int
	endAtStep int
}

func (w *drainMailboxOnStopProducer) DoWork(c actor.Context) actor.WorkerStatus {
	select {
	case <-c.Done():
		return actor.WorkerEnd
	default:
	}

	w.num++
	w.outMbx.Send(c, w.num) //nolint:errcheck // This example assumes the mailbox will never stop before this worker.

	if w.num == w.endAtStep {
		return actor.WorkerEnd
	}

	return actor.WorkerContinue
}

type drainMailboxOnStopConsumer struct {
	inMbx actor.MailboxReceiver[int]
}

func (w *drainMailboxOnStopConsumer) DoWork(c actor.Context) actor.WorkerStatus {
	select {
	case <-c.Done():
		return actor.WorkerEnd

	case num, isOpen := <-w.inMbx.ReceiveC():
		if !isOpen {
			return actor.WorkerEnd
		}

		if num == 3 {
			// Consumer is very slow to process message 3.
			time.Sleep(time.Second)
		}

		fmt.Printf("consumed %d\n", num)

		return actor.WorkerContinue
	}
}
