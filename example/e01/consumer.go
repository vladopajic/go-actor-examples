package e01

import (
	"fmt"

	"github.com/vladopajic/go-actor/actor"
)

// consumerWorker will consume numbers received on inC channel
type consumerWorker struct {
	inMbx actor.MailboxReceiver[int]
	id    int
}

func (w *consumerWorker) DoWork(c actor.Context) actor.WorkerStatus {
	select {
	case num := <-w.inMbx.ReceiveC():
		fmt.Printf("consumed %d \t(worker %d)\n", num, w.id)

		return actor.WorkerContinue

	case <-c.Done():
		return actor.WorkerEnd
	}
}
