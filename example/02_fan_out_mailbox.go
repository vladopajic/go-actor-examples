package example

import (
	"fmt"
	"time"

	"github.com/vladopajic/go-actor/actor"

	"github.com/vladopajic/go-actor-examples/lib"
)

func init() {
	register(2, "fan_out_mailbox", Run02)
}

// Run02 demonstrates how to fan-out a mailbox. It is very similar to Run01,
// except this time we intentionally want to have single producer that sends
// messages to many consumers.
func Run02() {
	mbx := actor.NewMailbox[int]()
	mbxx := actor.NewMailboxes[int](3)

	// Everything received by mbx will be forwarded to all mailboxes in mbxx.
	actor.FanOut(mbx.ReceiveC(), mbxx)

	pw := &fanOutProducer{outMbx: mbx}
	cw1 := &fanOutConsumer{inMbx: mbxx[0], id: 1}
	cw2 := &fanOutConsumer{inMbx: mbxx[1], id: 2}
	cw3 := &fanOutConsumer{inMbx: mbxx[2], id: 3}

	a := actor.Combine(
		mbx,
		actor.FromMailboxes(mbxx),
		actor.New(pw),
		actor.New(cw1),
		actor.New(cw2),
		actor.New(cw3),
	).Build()

	a.Start()
	defer a.Stop()

	<-lib.WaitForTermination()
}

// fanOutProducer produces an incremented number on a 1 second interval.
type fanOutProducer struct {
	outMbx actor.MailboxSender[int]
	num    int
}

func (w *fanOutProducer) DoWork(c actor.Context) actor.WorkerStatus {
	select {
	case <-c.Done():
		return actor.WorkerEnd

	case <-time.After(time.Second):
		w.num++
		w.outMbx.Send(c, w.num)

		return actor.WorkerContinue
	}
}

// fanOutConsumer consumes numbers received from its mailbox.
type fanOutConsumer struct {
	inMbx actor.MailboxReceiver[int]
	id    int
}

func (w *fanOutConsumer) DoWork(c actor.Context) actor.WorkerStatus {
	select {
	case <-c.Done():
		return actor.WorkerEnd

	case num := <-w.inMbx.ReceiveC():
		fmt.Printf("consumed %d \t(worker %d)\n", num, w.id)

		return actor.WorkerContinue
	}
}
