package e02

import (
	"time"

	"github.com/vladopajic/go-actor/actor"
)

// producerWorker will produce incremented number on 1 second interval
type producerWorker struct {
	outMbx actor.MailboxSender[int]
	num    int
}

func (w *producerWorker) DoWork(c actor.Context) actor.WorkerStatus {
	select {
	case <-time.After(time.Second):
		w.num++
		w.outMbx.Send(c, w.num)

		return actor.WorkerContinue

	case <-c.Done():
		return actor.WorkerEnd
	}
}
