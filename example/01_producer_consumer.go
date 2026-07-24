package example

import (
	"fmt"
	"time"

	"github.com/vladopajic/go-actor/actor"
)

func init() {
	register(1, "producer_consumer", Run01)
}

// Run01 demonstrates how to create actors for a producer-consumer use case.
// Producer will create incremented number on every 1 second interval and
// consumer will print whatever number it receives.
//
// Expected behavior:
//   - Runs until interrupted with Ctrl+C.
//   - Every number is consumed by one of the consumers.
//
// Watch for:
//   - This is work sharing, not broadcasting. With one shared mailbox, each
//     message is received by only one consumer.
func Run01() {
	mailbox := actor.NewMailbox[int]()

	// Producer and consumer workers are created with same mailbox
	// so that producer worker can send messages directly to consumer worker.
	pw := &producerConsumerProducer{outMbx: mailbox}
	cw1 := &producerConsumerConsumer{inMbx: mailbox, id: 1}

	// Note: Example creates two consumers for the sake of demonstration
	// since having one or more consumers will produce the same result.
	// Message on stdout will be written by first consumer that reads from mailbox.
	cw2 := &producerConsumerConsumer{inMbx: mailbox, id: 2}

	// Create actors using these workers and combine them to single actor.
	a := actor.Combine(
		mailbox,
		actor.New(pw),
		actor.New(cw1),
		actor.New(cw2),
	).Build()

	// Finally all actors are started and stopped at once.
	a.Start()
	defer a.Stop()

	<-WaitForTermination()
}

// producerConsumerProducer produces an incremented number on a 1 second interval.
type producerConsumerProducer struct {
	outMbx actor.MailboxSender[int]
	num    int
}

func (w *producerConsumerProducer) DoWork(c actor.Context) actor.WorkerStatus {
	select {
	case <-c.Done():
		return actor.WorkerEnd

	case <-time.After(time.Second):
		w.num++
		if err := w.outMbx.Send(c, w.num); err != nil {
			// The result of sending to a mailbox has to be processed.
			// If the mailbox is stopped, this worker has no more useful work
			// because its only output is that mailbox.
			return actor.WorkerEnd
		}

		return actor.WorkerContinue
	}
}

// producerConsumerConsumer consumes numbers received from the mailbox.
type producerConsumerConsumer struct {
	inMbx actor.MailboxReceiver[int]
	id    int
}

func (w *producerConsumerConsumer) DoWork(c actor.Context) actor.WorkerStatus {
	select {
	case <-c.Done():
		return actor.WorkerEnd

	case num, isOpen := <-w.inMbx.ReceiveC():
		if !isOpen {
			// Mailbox is closed, so no more messages can arrive and worker can end.
			return actor.WorkerEnd
		}

		fmt.Printf("consumed %d \t(worker %d)\n", num, w.id)

		return actor.WorkerContinue
	}
}
