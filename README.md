# homerow

Learn touch typing in your terminal, one row at a time.

`homerow` walks you through a structured course: new keys are introduced one
pair at a time, you drill them, review them, race the clock with them, and play
a short game before moving on. Every exercise is scored on accuracy and speed
and earns up to five stars.

> Status: in development. Module 1 (home row, lessons 1–23) is the first goal.
> See [docs/PLAN.md](docs/PLAN.md).

## Development

```sh
make check   # gofmt, go vet, golangci-lint, go test -race
make build   # ./bin/homerow
```

The design is hexagonal — see
[docs/adr/0001-hexagonal-architecture.md](docs/adr/0001-hexagonal-architecture.md).
