package e04

import (
	"fmt"
	"time"

	"github.com/vladopajic/go-actor/actor"
)

var (
	_ actor.StartableWorker = (*countdownWorker)(nil)
	_ actor.StoppableWorker = (*countdownWorker)(nil)
)

// NewCountdownActor creates new actor for launch pad countdowns.
func NewCountdownActor(launchReadySigC chan struct{}) actor.Actor {
	w := &countdownWorker{
		launchReadySigC: launchReadySigC,
		secondsCount:    3,
	}

	// Note: in example 3, onStop and onStart options were provided here:
	//
	// return actor.New(w,
	// 	actor.OptOnStart(w.onStart),
	// 	actor.OptOnStop(w.onStop),
	// )
	//
	// Instead of providing options we can implement OnStart and OnStop
	// function in worker.

	return actor.New(w)
}

type countdownWorker struct {
	launchReadySigC chan struct{}
	secondsCount    int
}

func (w *countdownWorker) DoWork(c actor.Context) actor.WorkerStatus {
	select {
	case <-time.After(time.Second):
		fmt.Printf("%d\n", w.secondsCount)

		w.secondsCount--

		if w.secondsCount == 0 {
			w.launchReadySigC <- struct{}{}
			return actor.WorkerEnd
		}

		return actor.WorkerContinue

	case <-c.Done():
		return actor.WorkerEnd
	}
}

func (w *countdownWorker) OnStart(c actor.Context) {
	fmt.Println("countdown started")
}

func (w *countdownWorker) OnStop() {
	fmt.Println("countdown ended")
}
