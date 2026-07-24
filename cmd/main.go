package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/vladopajic/go-actor-examples/example"
)

func main() {
	var exampleName string
	//nolint:lll // usage description text in one line
	flag.StringVar(&exampleName, "example", "", fmt.Sprintf("Example number or name to be started. Expected value in: 1 ... %d, or %s", len(example.Names()), strings.Join(example.Names(), ", ")))
	flag.Parse()

	if exampleName == "" {
		fmt.Printf("specify example you wish to run with -example flag\n")
		return
	}

	runExample(exampleName)
}

func runExample(name string) {
	fmt.Printf("======================= running example %s\n\n", name)

	if !example.Run(name) {
		fmt.Printf("example %s not found\n", name)
	}
}
