package e08

import (
	"fmt"
	"time"

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
			return actor.WorkerEnd
		}

		if num == 3 {
			// consumer is vary slow to process message 3
			time.Sleep(time.Second)
		}

		fmt.Printf("consumed %d\n", num)

		return actor.WorkerContinue
	}
}
