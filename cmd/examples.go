package main

import (
	"github.com/vladopajic/go-actor-examples/example/e01"
	"github.com/vladopajic/go-actor-examples/example/e02"
	"github.com/vladopajic/go-actor-examples/example/e03"
	"github.com/vladopajic/go-actor-examples/example/e04"
	"github.com/vladopajic/go-actor-examples/example/e05"
	"github.com/vladopajic/go-actor-examples/example/e06"
	"github.com/vladopajic/go-actor-examples/example/e07"
	"github.com/vladopajic/go-actor-examples/example/e08"
	"github.com/vladopajic/go-actor-examples/example/e09"
	"github.com/vladopajic/go-actor-examples/example/e10"
	"github.com/vladopajic/go-actor-examples/example/e11"
)

func examples() map[int]func() {
	return map[int]func(){
		1:  e01.Run,
		2:  e02.Run,
		3:  e03.Run,
		4:  e04.Run,
		5:  e05.Run,
		6:  e06.Run,
		7:  e07.Run,
		8:  e08.Run,
		9:  e09.Run,
		10: e10.Run,
		11: e11.Run,
	}
}
