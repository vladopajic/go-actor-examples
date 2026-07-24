package example

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/vladopajic/go-actor/actor"
)

func init() {
	register(10, "stop_combined_together", Run10)
}

// Run10 shows how to stop all combined actors as soon as any one of them stops.
//
// Expected behavior:
//   - Dots are printed while solvers search.
//   - The whole group stops after the first solver finds the target number.
//
// Watch for:
//   - The worker that searches for 10000 is intentionally impossible because
//     generated values are below 200. It exists to show group cancellation.
func Run10() {
	finishedC := make(chan any)

	a := actor.Combine(
		actor.New(&solverWorker{searchedNumber: 42}),
		actor.New(&solverWorker{searchedNumber: 42}),
		actor.New(&solverWorker{searchedNumber: 10000}), // this solver will never find solution
	).WithOptions(
		actor.OptStopTogether(), // stops all actors as soon as any one of them stops
		actor.OptOnStopCombined(func() { close(finishedC) }),
	).Build()

	a.Start()
	// defer a.Stop(), not needed because combined actor is stopped together.

	<-finishedC
}

type solverWorker struct {
	searchedNumber int32
}

func (w *solverWorker) DoWork(c actor.Context) actor.WorkerStatus {
	select {
	case <-c.Done():
		return actor.WorkerEnd
	default:
	}

	num := w.searchForSolution()
	if num == w.searchedNumber {
		fmt.Printf("\nsolution found!!!\n")
		return actor.WorkerEnd
	}

	fmt.Print(".")

	return actor.WorkerContinue
}

func (w *solverWorker) searchForSolution() int32 {
	time.Sleep(time.Millisecond * 20)
	return rand.Int31n(200) //nolint:gosec // relax
}
