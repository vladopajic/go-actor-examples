package main

import (
	"flag"
	"fmt"
)

func main() {
	var exampleNo int
	//nolint:lll // usage description text in one line
	flag.IntVar(&exampleNo, "example", 0, fmt.Sprintf("Example number to be started. Expacted value in range { 1 ... %d }", len(examples())))
	flag.Parse()

	if exampleNo <= 0 {
		fmt.Printf("specify example you wish to run with -example flag\n")
		return
	}

	runExample(exampleNo)
}

func runExample(no int) {
	runFn, exists := examples()[no]
	if exists {
		fmt.Printf("running example %d\n", no)
		runFn()
	} else {
		fmt.Printf("example %d not found\n", no)
	}
}
