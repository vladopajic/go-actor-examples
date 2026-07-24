package example

import (
	"fmt"
	"time"

	"github.com/vladopajic/go-actor/actor"
)

func init() {
	register(6, "condition_end_fix", Run06)
}

// Run06 fixes the problem introduced in Run05.
//
// Problem in Run05 is happening because combined actor was created using 3 actors:
// consumer, producer and mailbox. Because consumer and producers were stopped by
// themselves there was only mailbox actor that was left and its goroutine was
// blocked because it was waiting on messages, but no messages could be sent
// because there were no other goroutines alive that could possibly send to this
// mailbox - hence deadlock message in stdout.
//
// To fix this, combined actor is created only using consumer and producer.
// Therefore `OnStop` will be called because at one point both actors will be stopped.
//
// Expected behavior:
//   - Producer and consumer stop after three messages.
//   - "example finished" is printed.
//
// Watch for:
//   - The outer combined actor still owns the mailbox, but finishedC is closed
//     by the inner producer-consumer group.
func Run06() {
	finishedC := make(chan any)

	const endAtStep = 3

	mbx := actor.NewMailbox[int]()

	cp := actor.Combine(
		actor.New(&conditionEndFixProducer{outMbx: mbx, endAtStep: endAtStep}),
		actor.New(&conditionEndFixConsumer{inMbx: mbx, endAtStep: endAtStep}),
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

type conditionEndFixProducer struct {
	outMbx    actor.MailboxSender[int]
	num       int
	endAtStep int
}

func (w *conditionEndFixProducer) DoWork(c actor.Context) actor.WorkerStatus {
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

type conditionEndFixConsumer struct {
	inMbx     actor.MailboxReceiver[int]
	endAtStep int
}

func (w *conditionEndFixConsumer) DoWork(c actor.Context) actor.WorkerStatus {
	select {
	case <-c.Done():
		return actor.WorkerEnd

	case num := <-w.inMbx.ReceiveC():
		fmt.Printf("consumed %d\n", num)

		if num == w.endAtStep {
			return actor.WorkerEnd
		}

		return actor.WorkerContinue
	}
}
