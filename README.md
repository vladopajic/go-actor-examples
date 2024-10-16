# go-actor examples

[![test](https://github.com/vladopajic/go-actor-examples/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/vladopajic/go-actor-examples/actions/workflows/test.yml)
[![lint](https://github.com/vladopajic/go-actor-examples/actions/workflows/lint.yml/badge.svg?branch=main)](https://github.com/vladopajic/go-actor-examples/actions/workflows/lint.yml)


This repository hosts various examples for [`go-actor`](https://github.com/vladopajic/go-actor).

# How to use examples

To make the most out of this read each example package found in [`/example`](/example/) directory. Every example has `bootstrap.go` that sets up logic and explains idea behind this example. It is also advised to read other source files in example directory (package).

Every example cloud be executed with following command:
```
// make run {example number}

make run 1
```


| Example   | Description |
|-----------|------------|
|  1 | demonstrate how to create actors for producer-consumer use case |
|  2 | demonstrate how to fan-out Mailbox (example with multiple consumers)  |
|  3 | demonstrate how to create actors with options  |
|  4 | improved Countdown actor from example 3  |
|  5 | demonstrates when producer-consumer case should be ended when some condition is meet |
|  6 | example 5 introduced an issue (and a puzzle), that is fixed in this example |
|  7 | small improvement of example 6  |
|  8 | example when producer is much faster then consumer; new puzzle is introduced  |
|  9 | example with solution to puzzle from example 9  |
| 10 | demonstrate how to stop combined actor when first actor ends |


## Contribution

All contributions are useful, whether it is a simple typo, a more complex change, or just pointing out an issue. We welcome any contribution so feel free to open PR or issue. 
