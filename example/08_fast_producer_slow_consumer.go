package example

import (
	"fmt"
	"time"

	"github.com/vladopajic/go-actor/actor"
)

func init() {
	register(8, "fast_producer_slow_consumer", Run08)
}

// Run08 explores what happens if producer finishes much faster than consumer.
// Note there is just 64 messages in stdout, but we wanted to end at step 100.
//
// Solution for this puzzle is in Run09. This puzzle is bit harder as it requires
// deep understanding of how actor.Mailbox is implemented.
func Run08() {
	finishedC := make(chan any)

	const endAtStep = 100

	mbx := actor.NewMailbox[int]()
	mbx.Start()

	a := actor.Combine(
		actor.New(&fastProducerSlowConsumerProducer{outMbx: mbx, endAtStep: endAtStep}, actor.OptOnStop(mbx.Stop)),
		actor.New(&fastProducerSlowConsumerConsumer{inMbx: mbx}),
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

type fastProducerSlowConsumerProducer struct {
	outMbx    actor.MailboxSender[int]
	num       int
	endAtStep int
}

func (w *fastProducerSlowConsumerProducer) DoWork(c actor.Context) actor.WorkerStatus {
	select {
	case <-c.Done():
		return actor.WorkerEnd
	default:
		w.num++
		w.outMbx.Send(c, w.num) //nolint:errcheck // This example assumes the mailbox will never stop before this worker.

		if w.num == w.endAtStep {
			return actor.WorkerEnd
		}

		return actor.WorkerContinue
	}
}

type fastProducerSlowConsumerConsumer struct {
	inMbx actor.MailboxReceiver[int]
}

func (w *fastProducerSlowConsumerConsumer) DoWork(c actor.Context) actor.WorkerStatus {
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
