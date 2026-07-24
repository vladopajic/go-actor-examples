# go-actor examples

[![test](https://github.com/vladopajic/go-actor-examples/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/vladopajic/go-actor-examples/actions/workflows/test.yml)
[![lint](https://github.com/vladopajic/go-actor-examples/actions/workflows/lint.yml/badge.svg?branch=main)](https://github.com/vladopajic/go-actor-examples/actions/workflows/lint.yml)

This repository hosts various examples for [`go-actor`](https://github.com/vladopajic/go-actor).

## How to use examples

These examples are meant to be read in order. They start with basic actor
building blocks, then move into mailbox ownership, lifecycle coordination,
backpressure, and service composition.

Before reading the examples, it helps to know the core terms:

- `Actor`: a component that can be started and stopped.
- `Worker`: the logic run by an actor.
- `DoWork`: the worker method called repeatedly while work should continue.
- `WorkerContinue`: tells the actor to call `DoWork` again.
- `WorkerEnd`: tells the actor the worker is finished.
- `Mailbox`: a message queue used to pass values between actors.
- `Start` and `Stop`: lifecycle methods that begin and end actor execution.
- `Context.Done`: the cancellation signal workers should watch so they can stop promptly.

## Workflow 
To make the most out of these examples
- **Read example**: each example file found in the [/example](/example/) directory. 
  Every example is contained in a single `NN_name.go` file that sets up the logic, explains the idea behind the example,and includes the supporting code.
- **Run example**
  - Observe output to stdout
- **Play with code**
  - Some examples intorduces puzzles, try to crack them before proceeding next

Each example can be executed with the following command:

```bash
make run 1                     # run by number
make run producer_consumer     # run by name
make run 01_producer_consumer  # run by full file name
```


## Contribution

All contributions are useful, whether it is a simple typo, a more complex change, or just pointing out an issue. We welcome any contribution so feel free to open a PR or issue.
