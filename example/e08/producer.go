package e08

import (
	"github.com/vladopajic/go-actor/actor"
)

type producerWorker struct {
	outMbx actor.MailboxSender[int]
	num    int
}

func (w *producerWorker) DoWork(c actor.Context) actor.WorkerStatus {
	select {
	case <-c.Done():
		return actor.WorkerEnd
	default:
		w.num++
		w.outMbx.Send(c, w.num)

		if w.num == endWorkAtStep {
			return actor.WorkerEnd
		}

		return actor.WorkerContinue
	}
}
