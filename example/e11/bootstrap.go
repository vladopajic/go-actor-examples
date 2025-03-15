package e11

import (
	"fmt"
	"net/http"

	"github.com/vladopajic/go-actor/actor"

	"github.com/vladopajic/go-actor-examples/lib"
)

// This example demonstrates how to create a custom actor (HTTPService)
// and seamlessly compose it with other actors.
//
// After running example you can run `curl http://localhost:9988` to see message!
func Run() {
	messageActor := NewMessageActor()
	handler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, messageActor.Messenger.Message())
	})
	httpService := NewHTTPService("localhost:9988", handler)

	a := actor.Combine(messageActor, httpService).Build()

	a.Start()
	defer a.Stop()

	<-lib.WaitForTermination()
}
