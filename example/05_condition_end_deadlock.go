package example

import (
	"fmt"
	"time"

	"github.com/vladopajic/go-actor/actor"
)

func init() {
	register(5, "condition_end_deadlock", Run05)
}

// Run05 shows how to end producer and consumer actors when a condition is reached.
// Run01 and Run02 demonstrated producer and consumer working indefinitely, but
// sometimes we want to end program when a condition is reached - in this example
// it is after 3 working steps.
//
// This example will signal consumer and producer actors to end. And after they both
// end we will terminate program.
//
// Expected behavior:
//   - This example intentionally does not finish cleanly.
//   - "example finished" is not printed.
//
// Watch for:
//   - The mailbox remains part of the combined actor after producer and consumer
//     end. Run06 shows one way to structure the lifecycle correctly.
func Run05() {
	finishedC := make(chan any)

	const endAtStep = 3

	mbx := actor.NewMailbox[int]()

	a := actor.Combine(
		actor.New(&conditionEndDeadlockProducer{outMbx: mbx, endAtStep: endAtStep}),
		actor.New(&conditionEndDeadlockConsumer{inMbx: mbx, endAtStep: endAtStep}),
		mbx,
	).WithOptions(
		// Add function to be executed before all actors are started.
		actor.OptOnStartCombined(func(_ actor.Context) {
			fmt.Println("example started")
		}),

		// Add function to be executed after all actors are stopped.
		actor.OptOnStopCombined(func() {
			fmt.Println("example finished")
			close(finishedC)
			// Note this will never be called, this was intentional puzzle.
			// Try to think reason why. Solution for this puzzle is in Run06.
		}),
	).Build()

	a.Start()
	defer a.Stop()

	<-finishedC
}

type conditionEndDeadlockProducer struct {
	outMbx    actor.MailboxSender[int]
	num       int
	endAtStep int
}

func (w *conditionEndDeadlockProducer) DoWork(c actor.Context) actor.WorkerStatus {
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

type conditionEndDeadlockConsumer struct {
	inMbx     actor.MailboxReceiver[int]
	endAtStep int
}

func (w *conditionEndDeadlockConsumer) DoWork(c actor.Context) actor.WorkerStatus {
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
