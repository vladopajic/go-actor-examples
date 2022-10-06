package main

import (
	"github.com/vladopajic/go-actor-examples/example/e01"
	"github.com/vladopajic/go-actor-examples/example/e02"
	"github.com/vladopajic/go-actor-examples/example/e03"
	"github.com/vladopajic/go-actor-examples/example/e04"
)

func examples() map[int]func() {
	return map[int]func(){
		1: e01.Run,
		2: e02.Run,
		3: e03.Run,
		4: e04.Run,
	}
}
