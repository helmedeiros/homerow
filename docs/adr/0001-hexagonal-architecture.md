# ADR 0001: Hexagonal architecture

## Status

Accepted

## Context

The tutor has rules that are easy to state and easy to get subtly wrong:
how a typing mistake moves the cursor, how speed is measured, how stars are
earned, which characters a drill may use. It also has a terminal UI and a
progress file, both of which are slow and awkward to test.

## Decision

Keep the rules in a core that knows nothing about terminals or files, and reach
the outside world only through ports.

```
cmd/homerow                  composition root: builds adapters, injects them
internal/domain/...          pure rules: keyboard, drill, scoring, exercise,
                             curriculum, game
internal/app                 use cases and the ports they need
                             (Catalog, ProgressStore)
internal/adapters/content    driven: the lessons of each module
internal/adapters/filestore  driven: progress saved as JSON
internal/adapters/tui        driving: Bubble Tea terminal UI
```

Dependencies point inward: adapters import `app` and `domain`; `app` imports
`domain`; `domain` imports only the standard library. Time and randomness are
passed in, so every rule is deterministic under test.

## Consequences

- Domain and use cases are tested fast and without fakes of the terminal.
- The UI can be replaced (or driven by a test) without touching the rules.
- A little more ceremony: ports and wiring live in separate packages.
