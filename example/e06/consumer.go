package e06

import (
	"fmt"

	"github.com/vladopajic/go-actor/actor"
)

type consumerWorker struct {
	inMbx actor.MailboxReceiver[int]
}

func (w *consumerWorker) DoWork(c actor.Context) actor.WorkerStatus {
	select {
	case <-c.Done():
		return actor.WorkerEnd

	case num := <-w.inMbx.ReceiveC():
		fmt.Printf("consumed %d\n", num)

		if num == endWorkAtStep {
			return actor.WorkerEnd
		}

		return actor.WorkerContinue
	}
}
