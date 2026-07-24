# go-actor examples

[![test](https://github.com/vladopajic/go-actor-examples/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/vladopajic/go-actor-examples/actions/workflows/test.yml)
[![lint](https://github.com/vladopajic/go-actor-examples/actions/workflows/lint.yml/badge.svg?branch=main)](https://github.com/vladopajic/go-actor-examples/actions/workflows/lint.yml)

This repository hosts various examples for [`go-actor`](https://github.com/vladopajic/go-actor).

## How to use examples

To make the most out of these examples, read each example file found in the [/example](/example/) directory. Every example is contained in a single `NN_name.go` file that sets up the logic, explains the idea behind the example, and includes the supporting code.

Each example can be executed with the following command:

```bash
make run 1                     # run by number
make run producer_consumer     # run by name
make run 01_producer_consumer  # run by full file name
```


| Example | Description |
|---------|------------|
| `01_producer_consumer` | Demonstrates how to create actors for a producer-consumer use case. |
| `02_fan_out_mailbox` | Demonstrates how to fan-out a mailbox (an example with multiple consumers). |
| `03_actor_options` | Demonstrates how to create actors with options. |
| `04_stoppable_countdown` | An improved Countdown actor, building on example 3. |
| `05_condition_end_deadlock` | Demonstrates when a producer-consumer case should end based on a specific condition. |
| `06_condition_end_fix` | Builds on example 5, fixing an introduced issue (and puzzle). |
| `07_stop_mailbox_on_producer_stop` | A small improvement to example 6. |
| `08_fast_producer_slow_consumer` | Explores a scenario where the producer is much faster than the consumer; introduces a new puzzle. |
| `09_drain_mailbox_on_stop` | Provides the solution to the puzzle introduced in example 8. |
| `10_stop_combined_together` | Demonstrates how to stop a combined actor when the first actor ends. |
| `11_http_service` | Demonstrates how to create a custom actor (HTTPService) and seamlessly compose it with other actors. |


## Contribution

All contributions are useful, whether it is a simple typo, a more complex change, or just pointing out an issue. We welcome any contribution so feel free to open a PR or issue.
