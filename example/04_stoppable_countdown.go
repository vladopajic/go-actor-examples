package example

import (
	"fmt"
	"time"

	"github.com/vladopajic/go-actor/actor"
)

func init() {
	register(4, "stoppable_countdown", Run04)
}

var (
	_ actor.StartableWorker = (*stoppableCountdownWorker)(nil)
	_ actor.StoppableWorker = (*stoppableCountdownWorker)(nil)
)

// Run04 shows improved Countdown actor from Run03.
//
// Expected behavior:
//   - Countdown starts, the actor is stopped before launch, and launchReadySigC
//     is never closed.
//
// Watch for:
//   - The final receive demonstrates that launch never becomes ready after
//     cancellation. The program exits because no goroutine can make progress.
func Run04() {
	launchReadySigC := make(chan struct{})

	a := newStoppableCountdownActor(launchReadySigC)

	a.Start()

	// Stop before countdown has ended.
	time.Sleep(time.Second * 2)
	a.Stop()

	// This program will wait for launchReadySigC but it will never
	// happen because countdown was stopped. Program will exit anyway
	// because all goroutines are asleep.
	//
	// This is intentional for the example. In application code, prefer waiting
	// on a real completion signal or returning after Stop.
	<-launchReadySigC
}

// newStoppableCountdownActor creates new actor for launch pad countdowns.
func newStoppableCountdownActor(launchReadySigC chan struct{}) actor.Actor {
	w := &stoppableCountdownWorker{
		launchReadySigC: launchReadySigC,
		secondsCount:    3,
	}

	// Note: in Run03, onStop and onStart options were provided here:
	//
	// return actor.New(w,
	// 	actor.OptOnStart(w.onStart),
	// 	actor.OptOnStop(w.onStop),
	// )
	//
	// Instead of providing options we can implement OnStart and OnStop
	// functions in worker.

	return actor.New(w)
}

type stoppableCountdownWorker struct {
	launchReadySigC chan struct{}
	secondsCount    int
}

func (w *stoppableCountdownWorker) DoWork(c actor.Context) actor.WorkerStatus {
	select {
	case <-c.Done():
		return actor.WorkerEnd

	case <-time.After(time.Second):
		fmt.Printf("%d\n", w.secondsCount)

		w.secondsCount--

		if w.secondsCount == 0 {
			close(w.launchReadySigC)
			return actor.WorkerEnd
		}

		return actor.WorkerContinue
	}
}

func (w *stoppableCountdownWorker) OnStart(_ actor.Context) {
	fmt.Println("countdown started")
}

func (w *stoppableCountdownWorker) OnStop() {
	fmt.Println("countdown ended")
}
