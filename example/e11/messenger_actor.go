package e11

import (
	"math/rand"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/vladopajic/go-actor/actor"
)

type Messenger interface {
	Message() string
}

type MessageActor struct {
	actor.Actor
	Messenger
}

func NewMessageActor() *MessageActor {
	w := &messageWorker{
		reqC: make(chan chan string),
	}

	return &MessageActor{
		Actor:     actor.New(w),
		Messenger: w,
	}
}

type messageWorker struct {
	reqC       chan chan string
	currentMsg string
}

func (w *messageWorker) Message() string {
	req := make(chan string, 1)
	w.reqC <- req

	return <-req
}

func (w *messageWorker) OnStart(actor.Context) {
	w.handleGenerateMessage()
}

func (w *messageWorker) DoWork(ctx actor.Context) actor.WorkerStatus {
	select {
	case <-ctx.Done():
		return actor.WorkerEnd
	case <-time.After(time.Second * 2):
		w.handleGenerateMessage()
		return actor.WorkerContinue
	case reqC := <-w.reqC:
		w.handleResponse(reqC)
		return actor.WorkerContinue
	}
}

func (w *messageWorker) handleResponse(reqC chan string) {
	log.Info().Msg("handling message request")

	reqC <- w.currentMsg
}

func (w *messageWorker) handleGenerateMessage() {
	log.Info().Msg("generating new message")

	w.currentMsg = msgs[rand.Intn(len(msgs))]
}

var msgs = []string{
	"Build resilient, concurrent systems with go-actor - a lightweight actor framework for Go",
	"The go-actor library lets you model your Go applications with message-passing actors - no more race conditions!",
	"go-actor + Go’s goroutines = highly scalable, actor-driven systems with minimal overhead!",
	"Want better fault isolation? Each go-actor runs independently, handling failures gracefully!",
	"With go-actor, you get fast, non-blocking message passing - perfect for high-performance apps",
	"Manage complex concurrent workflows with go-actor, a Go library designed for event - driven architectures.",
	"go-actor makes writing concurrent Go programs easier, safer, and more intuitive!",
	"Focus on business logic, not concurrency bugs, let go-actor handle the tough parts for you!",
	"Scale Go applications effortlessly with go-actor, because actors and CSP are a perfect match!",
}
