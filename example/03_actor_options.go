package example

import (
	"fmt"
	"time"

	"github.com/vladopajic/go-actor/actor"
)

func init() {
	register(3, "actor_options", Run03)
}

// Run03 demonstrates how to create actors with options.
//
// Expected behavior:
//   - Prints a countdown, then prints the launch message.
//
// Watch for:
//   - DoWork intentionally does not watch c.Done here. Run04 shows why
//     workers should normally support cancellation.
func Run03() {
	launchReadySigC := make(chan struct{})

	a := newOptionsCountdownActor(launchReadySigC)

	a.Start()
	// Note: It's good practice to stop actors, but in this case
	// worker will end after countdown is finished - Stop() will have no effect.
	defer a.Stop()

	<-launchReadySigC
	launchRocket()
}

func launchRocket() {
	fmt.Println("Launching rocket!")
}

// newOptionsCountdownActor creates new actor for launch pad countdowns.
func newOptionsCountdownActor(launchReadySigC chan struct{}) actor.Actor {
	w := &optionsCountdownWorker{
		launchReadySigC: launchReadySigC,
		secondsCount:    3,
	}

	return actor.New(w,
		actor.OptOnStart(w.onStart),
		actor.OptOnStop(w.onStop),
	)
}

type optionsCountdownWorker struct {
	launchReadySigC chan struct{}
	secondsCount    int
}

func (w *optionsCountdownWorker) DoWork(_ actor.Context) actor.WorkerStatus {
	// Note: it's bad practice to implement workers that are not
	// responding on c.Done() signal. See Run04.
	for i := w.secondsCount; i > 0; i-- {
		fmt.Printf("%d\n", i)
		time.Sleep(time.Second)
	}

	close(w.launchReadySigC)

	return actor.WorkerEnd
}

func (w *optionsCountdownWorker) onStart(_ actor.Context) {
	fmt.Println("countdown started")
}

func (w *optionsCountdownWorker) onStop() {
	fmt.Println("countdown ended")
}
