package e01

import (
	"github.com/vladopajic/go-actor/actor"

	"github.com/vladopajic/go-actor-examples/lib"
)

// This program will demonstrate how to create actors for producer-consumer use case.
// Producer will create incremented number on every 1 second interval and
// consumer will print whatever number it receives.
func Run() {
	mailbox := actor.NewMailbox[int]()

	// Produce and consume workers are created with same mailbox
	// so that produce worker can send messages directly to consume worker
	pw := &producerWorker{outMbx: mailbox}
	cw1 := &consumerWorker{inMbx: mailbox, id: 1}

	// Note: Example creates two consumers for the sake of demonstration
	// since having one or more consumers will produce the same result.
	// Message on stdout will be written by first consumer that reads from mailbox.
	cw2 := &consumerWorker{inMbx: mailbox, id: 2}

	// Create actors using these workers and combine them to singe actor
	a := actor.Combine(
		mailbox,
		actor.New(pw),
		actor.New(cw1),
		actor.New(cw2),
	).Build()

	// Finally all actors are started and stopped at once
	a.Start()
	defer a.Stop()

	<-lib.WaitForTermination()
}
