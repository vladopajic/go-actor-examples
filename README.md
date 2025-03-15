# go-actor examples

[![test](https://github.com/vladopajic/go-actor-examples/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/vladopajic/go-actor-examples/actions/workflows/test.yml)
[![lint](https://github.com/vladopajic/go-actor-examples/actions/workflows/lint.yml/badge.svg?branch=main)](https://github.com/vladopajic/go-actor-examples/actions/workflows/lint.yml)


This repository hosts various examples for [`go-actor`](https://github.com/vladopajic/go-actor).

# How to use examples

To make the most out of this, read each example package found in the [/example](/example/) directory. Every example has a bootstrap.go file that sets up the logic and explains the idea behind the example. It is also advised to read other source files in the example directory (package).

Each example can be executed with the following command:

```bash
# Use the following syntax to run a specific example:  
make run {example number}

make run 1
```


| Example   | Description |
|-----------|------------|
|  1 | Demonstrates how to create actors for a producer-consumer use case. |
|  2 | Demonstrates how to fan-out a mailbox (an example with multiple consumers).  |
|  3 | Demonstrates how to create actors with options. |
|  4 | An improved Countdown actor, building on example 3.  |
|  5 | Demonstrates when a producer-consumer case should end based on a specific condition. |
|  6 | Builds on example 5, fixing an introduced issue (and puzzle). |
|  7 | A small improvement to example 6. |
|  8 | Explores a scenario where the producer is much faster than the consumer; introduces a new puzzle.  |
|  9 | Provides the solution to the puzzle introduced in example 8.  |
| 10 | Demonstrates how to stop a combined actor when the first actor ends. |
| 11 | Create a custom actor (HTTPService) and seamlessly compose it with other actors. |


## Contribution

All contributions are useful, whether it is a simple typo, a more complex change, or just pointing out an issue. We welcome any contribution so feel free to open PR or issue. 
