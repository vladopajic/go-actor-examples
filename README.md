# go-actor examples

[![test](https://github.com/vladopajic/go-actor-examples/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/vladopajic/go-actor-examples/actions/workflows/test.yml)
[![lint](https://github.com/vladopajic/go-actor-examples/actions/workflows/lint.yml/badge.svg?branch=main)](https://github.com/vladopajic/go-actor-examples/actions/workflows/lint.yml)

This repository hosts various examples for [`go-actor`](https://github.com/vladopajic/go-actor).

## How to use examples

These examples are meant to be read in order. They work like a tutorial,
guiding users from basic actor building blocks to mailbox ownership, lifecycle
coordination, backpressure, and service composition.

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

Use the same loop for each example:

1. **Read** the next file in the [/example](/example/) directory. Each example
   is contained in a single `NN_name.go` file with the idea, setup, and
   supporting code in one place.
2. **Run** the example and watch the stdout output. The output usually reveals
   the actor lifecycle, mailbox behavior, or shutdown behavior being taught.
3. **Compare** the output with the comments at the top of the example. If the
   result is surprising, pause and reason about which actor is still running,
   which mailbox is open, or which stop signal was received.
4. **Experiment** by changing small values such as timer duration, consumer
   count, stop condition, or mailbox options.
5. **Continue** to the next example only after the current behavior is clear.

Each example can be executed with the following command:

```bash
make run 1                     # run by number
make run producer_consumer     # run by name
make run 01_producer_consumer  # run by full file name
```

Some examples are intentionally surprising. They may stop early, wait forever,
or expose a lifecycle problem. Treat those examples as puzzles: predict why the
behavior happens before reading the following fix example.

## Contribution

All contributions are useful, whether it is a simple typo, a more complex change, or just pointing out an issue. We welcome any contribution so feel free to open a PR or issue.
