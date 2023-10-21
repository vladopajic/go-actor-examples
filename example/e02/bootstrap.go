package e02

import (
	"github.com/vladopajic/go-actor/actor"

	"github.com/vladopajic/go-actor-examples/lib"
)

// This program will demonstrate how to fan-out Mailbox. Example is vary similar
// to previous except that this time we intentionally want to have single producer
// that sends messages to many consumers.
func Run() {
	mbx := actor.NewMailbox[int]()
	mbxx := actor.NewMailboxes[int](3)

	// everything received by mbx will be forwarded to all mailboxes in mbxx
	actor.FanOut(mbx.ReceiveC(), mbxx)

	pw := &producerWorker{outMbx: mbx}
	cw1 := &consumerWorker{inMbx: mbxx[0], id: 1}
	cw2 := &consumerWorker{inMbx: mbxx[1], id: 2}
	cw3 := &consumerWorker{inMbx: mbxx[2], id: 3}

	a := actor.Combine(
		mbx,
		actor.FromMailboxes(mbxx),
		actor.New(pw),
		actor.New(cw1),
		actor.New(cw2),
		actor.New(cw3),
	)

	a.Start()
	defer a.Stop()

	<-lib.WaitForTermination()
}
