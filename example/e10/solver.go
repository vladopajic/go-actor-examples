package e10

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/vladopajic/go-actor/actor"
)

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
