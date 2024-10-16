package e07

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

	case num, isOpen := <-w.inMbx.ReceiveC():
		if !isOpen {
			// When receiver channel is closed we can safely end
			// because no more messages could ever be sent to this channel.
			return actor.WorkerEnd
		}

		fmt.Printf("consumed %d\n", num)

		return actor.WorkerContinue
	}
}
